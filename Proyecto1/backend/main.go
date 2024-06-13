package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"main/Database"
	"main/Routes"

	// "log"
	// "main/Database"
	"strings"

	//"main/Database"
	"os/exec"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type child struct {
	Pid    int    `json:"pid"`
	Nombre string `json:"name"`
	Estado int    `json:"state"`
	Padre  int    `json:"pidPadre"`
}

type Process struct {
	Pid    int     `json:"pid"`
	Nombre string  `json:"name"`
	Estado int     `json:"state"`
	Padre  int     `json:"pidPadre"`
	Chil   []child `json:"child"`
}

type Cpu struct {
	Porcentaje int       `json:"cpu_porcentaje"`
	Procesos   []Process `json:"processes"`
}

type Ram struct {
	Total      int `json:"totalRam"`
	En_uso     int `json:"memoriaEnUso"`
	Libre      int `json:"libre"`
	Porcentaje int `json:"porcentaje"`
}

var process *exec.Cmd

func main() {
	app := fiber.New()

	if err := Database.Connect(); err != nil {
		log.Fatal("Error en", err)
	}
	// Habilitar CORS
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,DELETE",
	}))

	// Definir rutas
	app.Get("/cpu", getCPUInfo)
	app.Get("/ram", getRAMInfo)
	Routes.Setup(app)

	// Iniciar el servidor
	if err := app.Listen(":3000"); err != nil { //aqui no se porque da error al poner una ip
		fmt.Println("Error en el servidor")
	}
}

func getCPUInfo(c *fiber.Ctx) error {
	cmdCpu := exec.Command("sh", "-c", "cat /proc/cpu_so1_1s2024")
	outCpu, err := cmdCpu.CombinedOutput()
	if err != nil {
		return c.Status(500).SendString("Error al obtener información de la CPU")
	}

	var cpuInfo Cpu
	err = json.Unmarshal(outCpu, &cpuInfo)
	if err != nil {
		return c.Status(500).SendString("Error al parsear información de la CPU")
	}

	cpuFree := exec.Command("mpstat", "1", "1")
	var out bytes.Buffer
	cpuFree.Stdout = &out
	err = cpuFree.Run()
	if err != nil {
		return c.Status(500).SendString("Error al ejecutar mpstat")
	}

	output := out.String()
	lines := strings.Split(output, "\n")
	var idleStr string
	for _, line := range lines {
		if strings.Contains(line, "all") {
			fields := strings.Fields(line)
			if len(fields) >= 11 {
				idleStr = fields[10]
			}
			break
		}
	}

	idle, err := strconv.ParseFloat(idleStr, 64)
	if err != nil {
		return c.Status(500).SendString("Error al parsear el valor de %idle")
	}

	freeCPU := idle
	cpuInfo.Porcentaje = 100 - int(freeCPU)

	return c.JSON(cpuInfo)
}

func getRAMInfo(c *fiber.Ctx) error {
	cmdRam := exec.Command("sh", "-c", "cat /proc/ram_so1_jun2024")
	outRam, err := cmdRam.CombinedOutput()
	if err != nil {
		return c.Status(500).SendString("Error al obtener información de la RAM")
	}

	var ramInfo Ram
	err = json.Unmarshal(outRam, &ramInfo)
	if err != nil {
		return c.Status(500).SendString("Error al parsear información de la RAM")
	}

	return c.JSON(ramInfo)
}

/*
func main() {

	app := fiber.New()

	// if err := Database.Connect(); err != nil {
	// 	log.Fatal(err)
	// 	fmt.Println("Error en la base de datos")
	// }

	if err := app.Listen("backend:3000"); err != nil {
		fmt.Println("Error en el servidor")
	}

	interval := 5
	fmt.Println("Intervalo de tiempo: ", interval)

	ticker := time.NewTicker(time.Second * time.Duration(interval))
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			now := time.Now()
			second1 := now.Second()
			re := strconv.Itoa(now.Year()) + "-" + fmt.Sprintf("%02d", now.Month()) + "-" + fmt.Sprintf("%02d", now.Day()) + "-" + fmt.Sprintf("%02d", now.Hour()) + "-" + fmt.Sprintf("%02d", now.Minute()) + "-" + fmt.Sprintf("%02d", second1)
			// fmt.Println("re")
			fmt.Println(re)

			cmdCpu := exec.Command("sh", "-c", "cat /proc/cpu_so1_1s2024")
			outCpu, err := cmdCpu.CombinedOutput()
			if err != nil {
				fmt.Println("error", err)
				return
			}

			//---CPU

			cpuFree := exec.Command("mpstat", "1", "1")
			var out bytes.Buffer
			cpuFree.Stdout = &out
			err = cpuFree.Run()
			if err != nil {
				fmt.Println(err)
				return
			}

			fmt.Println("-------------------- CPU --------------------")
			var cpuInfo Cpu
			err = json.Unmarshal(outCpu, &cpuInfo)
			if err != nil {
				fmt.Println(err)
				return
			}

			output := out.String()
			// Dividir la salida en líneas
			lines := strings.Split(output, "\n")

			// Buscar la línea que contiene los datos agregados (all)
			var idleStr string
			for _, line := range lines {
				if strings.Contains(line, "all") {
					// Dividir la línea en campos por espacios y extraer el valor de %idle
					fields := strings.Fields(line)
					if len(fields) >= 11 {
						idleStr = fields[10]
					}
					break
				}
			}

			// Convertir el valor de %idle a float
			idle, err := strconv.ParseFloat(idleStr, 64)
			if err != nil {
				fmt.Println("Error parseando el valor de %idle:", err)
				return
			}

			// Calcular el porcentaje libre de CPU
			freeCPU := idle
			fmt.Printf("Porcentaje libre de CPU: %.2f%%\n", freeCPU)
			//fmt.Println("cpu informacion de processos: ", cpuInfo.Procesos)

			//Mandar respuesta (ejemplo simple de uso)
			for _, proceso := range cpuInfo.Procesos {
				fmt.Printf("Process ID: %d, Name: %s, State: %d, Parent ID: %d\n", proceso.Pid, proceso.Nombre, proceso.Estado, proceso.Padre)
				for _, child := range proceso.Chil {
					fmt.Printf("  Child Process ID: %d, Name: %s, State: %d, Parent ID: %d\n",
						child.Pid, child.Nombre, child.Estado, child.Padre)
				}
			}

			fmt.Println(" ==================== DATOS MODULO RAM ==================== ")
			fmt.Println(" ")

			cmdRam := exec.Command("sh", "-c", "cat /proc/ram_so1_jun2024")
			outRam, err := cmdRam.CombinedOutput()
			if err != nil {
				fmt.Println("error", err)
			}
			//---RAM
			fmt.Println("-------------------- RAM --------------------")
			var ram_info Ram
			err = json.Unmarshal([]byte(outRam), &ram_info)
			if err != nil {
				fmt.Println(err)
			}

			fmt.Println("Porcentaje de RAM:", ram_info.Porcentaje)
			fmt.Println("RAM en uso:", ram_info.En_uso)
			fmt.Println("RAM libre:", ram_info.Libre)
		}
	}
}
*/
