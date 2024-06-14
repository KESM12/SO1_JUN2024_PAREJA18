package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"main/Controller"
	"main/Database"
	"main/Model"

	// "log"
	// "main/Database"
	"strings"

	//"main/Database"
	"os/exec"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

// type Hijos struct {
// 	PID      int    `json:"pid"`
// 	Name     string `json:"name"`
// 	State    int    `json:"state"`
// 	PIDPadre int    `json:"pidPadre,omitempty"` // Se omite si está vacío.
// }

// type Process struct {
// 	PID   int     `json:"pid"`
// 	Name  string  `json:"name"`
// 	State int     `json:"state"`
// 	Child []Hijos `json:"child"`
// }

// type Cpu struct {
// 	Porcentaje int       `json:"cpu_porcentaje"`
// 	Processes  []Process `json:"processes"`
// }

// type Ram struct {
// 	Total      int `json:"totalRam"`
// 	En_uso     int `json:"memoriaEnUso"`
// 	Libre      int `json:"libre"`
// 	Porcentaje int `json:"porcentaje"`
// }

var process *exec.Cmd

func main() {
	app := fiber.New()

	// Conectar a la base de datos
	if err := Database.Connect(); err != nil {
		log.Fatal("Error connecting to the database:", err)
	} else {
		log.Println("Successfully connected to the database.")
	}
	// Habilitar CORS
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,DELETE",
	}))

	// Definir rutas
	app.Get("/cpuyram", getPOrcentajeRamyCpu)
	//app.Get("/cpu", getCPUInfo)
	//app.Get("/ram", )
	//getMem()
	//Routes.Setup(app)

	// Iniciar el servidor
	if err := app.Listen(":3000"); err != nil { //aqui no se porque da error al poner una ip
		fmt.Println("Error en el servidor")
	}
}

// Obtener información de la RAM y mostrarla en el Frontend
// Funcion para obtener datos de la RAM
func getRAMdata() (int, error) {
	cmd := exec.Command("sh", "-c", "cat /proc/ram_so1_jun2024")
	stdout, err := cmd.CombinedOutput()

	if err != nil {
		return 0, err
	}

	// Convertir la salida a formato JSON
	var data Model.Ram
	err = json.Unmarshal(stdout, &data)
	if err != nil {
		return 0, err
	}

	return data.Porcentaje, nil
}

func getPOrcentajeRamyCpu(c *fiber.Ctx) error {
	// Obtener datos de la RAM
	freeRAMPercentage, err := getRAMdata()
	if err != nil {
		return c.Status(500).SendString("Error al obtener datos de la RAM")
	}

	// Obtener datos de la CPU
	usedCPUPercentage, err := getCpuPercentage1()
	if err != nil {
		return c.Status(500).SendString("Error al obtener datos de la CPU")
	}

	usedRAMPercentage := 100 - freeRAMPercentage
	//freeCPUPercentage := 100 - usedCPUPercentage

	estadisticas := map[string]int{
		"ram_percentage": usedRAMPercentage,
		"cpu_percentage": usedCPUPercentage,
	}

	return c.JSON(estadisticas)
}

// func getRAMInfo(c *fiber.Ctx) error {
// 	cmdRam := exec.Command("sh", "-c", "cat /proc/ram_so1_jun2024")
// 	outRam, err := cmdRam.CombinedOutput()
// 	if err != nil {
// 		return c.Status(500).SendString("Error al obtener información de la RAM")
// 	}

// 	var ramInfo Model.Ram
// 	err = json.Unmarshal(outRam, &ramInfo)
// 	if err != nil {
// 		return c.Status(500).SendString("Error al parsear información de la RAM")
// 	}
// 	getMem()
// 	return c.JSON(ramInfo)
// }

func eliminarProceso(c *fiber.Ctx) error {
	pid := c.Params("pid")
	cmd := exec.Command("sh", "-c", "kill "+pid)
	err := cmd.Run()
	if err != nil {
		return c.Status(500).SendString("Error al eliminar el proceso")
	}

	return c.SendString("Proceso eliminado exitosamente")
}

// Obtener información de la RAM
func getRAMInfo1() (*Model.Ram, error) {
	cmdRam := exec.Command("sh", "-c", "cat /proc/ram_so1_jun2024")
	outRam, err := cmdRam.CombinedOutput()
	if err != nil {
		return nil, err
	}

	var ramInfo Model.Ram
	err = json.Unmarshal(outRam, &ramInfo)
	if err != nil {
		return nil, err
	}

	return &ramInfo, nil
}

// Mandar la información a la base de datos
func getMem() (Model.Ram, error) {
	ramInfo, err := getRAMInfo1()
	if err != nil {
		return Model.Ram{}, err
	}
	total := ramInfo.Total
	enUso := ramInfo.En_uso
	libre := ramInfo.Libre
	porcentaje := ramInfo.Porcentaje

	// convertir los valores de bytes a MB
	total = total / (1024 * 2)
	enUso = enUso / (1024 * 2)
	libre = libre / (1024 * 2)
	DbTotal := total
	DbEnUso := enUso
	Dblibre := libre
	DbPorcentaje := porcentaje

	Controller.InsertRam("ram", DbTotal, DbEnUso, Dblibre, DbPorcentaje)

	return Model.Ram{
		Total:      total,
		En_uso:     enUso,
		Porcentaje: porcentaje,
		Libre:      libre,
	}, nil
}

// Obtener información de la CPU y mostrarla en el Frontend
func getCPUInfo(c *fiber.Ctx) error {
	cmdCpu := exec.Command("sh", "-c", "cat /proc/cpu_so1_1s2024")
	outCpu, err := cmdCpu.CombinedOutput()
	if err != nil {
		return c.Status(500).SendString("Error al obtener información de la CPU")
	}

	var cpuInfo Model.Cpu
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

	for _, process := range cpuInfo.Processes {
		err := getCPU(&process)
		if err != nil {
			fmt.Println("Error:", err)
		}
	}
	getCpuPercentage("cpu%")

	return c.JSON(cpuInfo)
}

func getCPUInfo1(out string) (*Model.Cpu, error) {
	var cpuInfo Model.Cpu
	var data map[string]interface{}
	err := json.Unmarshal([]byte(out), &data)
	if err != nil {
		return nil, err
	}

	porcentaje, ok := data["cpu_porcentaje"]
	if !ok {
		return nil, fmt.Errorf("No se encontró el campo 'cpu_porcentaje'")
	}
	cpuInfo.Porcentaje = int(porcentaje.(float64))

	processesData, ok := data["processes"]
	if !ok {
		return nil, fmt.Errorf("No se encontró el campo 'processes'")
	}

	procesos, err := getProcesses(processesData)
	if err != nil {
		return nil, err
	}
	cpuInfo.Processes = procesos
	return &cpuInfo, nil
}

func getProcesses(processesData interface{}) ([]Model.Process, error) {
	var procesos []Model.Process
	processesJSON, err := json.Marshal(processesData)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(processesJSON, &procesos)
	if err != nil {
		log.Fatal(err)
	}
	return procesos, nil
}

func getCPU(cpuInfo *Model.Process) error {
	PID := cpuInfo.PID
	Name := cpuInfo.Name
	State := cpuInfo.State
	if len(cpuInfo.Child) > 0 {
		PidPadre := cpuInfo.Child[0].PID
		Controller.InserProcess("cpu", PID, Name, State, PidPadre)
		for _, hijo := range cpuInfo.Child {
			Controller.InserProcess("cpu", hijo.PID, hijo.Name, hijo.State, hijo.PIDPadre)
		}
	} else {
		Controller.InserProcess("cpu", PID, Name, State, 0)
	}
	return nil
}

// Función separada para obtener y almacenar el porcentaje de CPU
func getCpuPercentage(nameCol string) (int, error) {
	cpuFree := exec.Command("mpstat", "1", "1")
	var out bytes.Buffer
	cpuFree.Stdout = &out
	err := cpuFree.Run()
	if err != nil {
		return 0, fmt.Errorf("error al ejecutar mpstat: %v", err)
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
		return 0, fmt.Errorf("error al parsear el valor de %%idle: %v", err)
	}

	freeCPU := idle
	cpuPercentage := 100 - int(freeCPU)

	// Insertar el porcentaje de CPU en la base de datos
	err = Controller.InsertCpu(nameCol, cpuPercentage)
	if err != nil {
		return 0, fmt.Errorf("error al insertar el porcentaje de CPU en la base de datos: %v", err)
	}

	return cpuPercentage, nil
}

func getCpuPercentage1() (int, error) {
	cpuFree := exec.Command("mpstat", "1", "1")
	var out bytes.Buffer
	cpuFree.Stdout = &out
	err := cpuFree.Run()
	if err != nil {
		return 0, fmt.Errorf("error al ejecutar mpstat: %v", err)
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
		return 0, fmt.Errorf("error al parsear el valor de %%idle: %v", err)
	}

	freeCPU := idle
	cpuPercentage := 100 - int(freeCPU)

	return cpuPercentage, nil
}

// func getCPU(cpuInfo *Model.Process) error {
// 	Controller.InsertData1("cpu")
// 	PID := cpuInfo.PID
// 	Name := cpuInfo.Name
// 	State := cpuInfo.State
// 	PidPadre := cpuInfo.Child[0].PID
// 	Controller.InsertData2("cpu", PID, Name, State, PidPadre)
// 	for _, hijo := range cpuInfo.Child {
// 		Controller.InsertData2("cpu", hijo.PID, hijo.Name, hijo.State, hijo.PIDPadre)
// 	}
// 	return nil
// }
