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

	// "main/Database"
	"os/exec"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

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
	app.Get("/cpuyram", getPorcentajeRamyCpu)
	app.Get("/cpu", getCPUInfo)
	app.Get("cpu/iniProc/crear", StartProcess)
	app.Post("/cpu/killProc", KillProcess)
	//startDeletionRoutine()

	//app.Get("/ram", )
	//getMem()
	//Routes.Setup(app)

	// Iniciar el servidor
	if err := app.Listen(":3000"); err != nil { //aqui no se porque da error al poner una ip
		fmt.Println("Error en el servidor")
	}
}

// Obtener información de la RAM y mostrarla en el Frontend
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

func getPorcentajeRamyCpu(c *fiber.Ctx) error {
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
	getMem()
	getCpuPercentage("cpu%")
	return c.JSON(estadisticas)
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
	total = total / (2048)
	enUso = enUso / (2048)
	libre = libre / (2048)
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
	Controller.ResetCollection("cpu")
	for _, process := range cpuInfo.Processes {
		err := getCPU(&process)
		if err != nil {
			fmt.Println("Error:", err)
		}
	}
	getCpuPercentage("cpu%")

	return c.JSON(cpuInfo)
}

func getCPU(cpuInfo *Model.Process) error {
	//Controller.ResetCollection("cpu")
	PID := cpuInfo.PID
	Name := cpuInfo.Name
	State := cpuInfo.State
	var PidPadre int
	if len(cpuInfo.Child) > 0 {
		PidPadre = cpuInfo.Child[0].PID
	} else {
		PidPadre = 0
	}
	Controller.InserProcess("cpu", PID, Name, State, PidPadre)

	// Insertar solo los procesos hijo sin duplicar el proceso padre
	// for _, hijo := range cpuInfo.Child {
	// 	Controller.InserProcess("cpu", hijo.PID, hijo.Name, hijo.State, PID)
	// }

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

// Función para insertar el porcentaje de CPU en la base de datos
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

// función para crear el proceso sleep infinity
func StartProcess(c *fiber.Ctx) error {
	// Crear un nuevo proceso con un comando de espera
	cmd := exec.Command("sleep", "infinity")
	err := cmd.Start()
	if err != nil {
		fmt.Print(err)
		return c.Status(fiber.StatusInternalServerError).SendString("Error al iniciar el proceso")
	}

	// Obtener el comando con PID
	process = cmd
	fmt.Println("Proceso iniciado con PID:", process.Process.Pid)
	return c.SendString(fmt.Sprintf("Proceso iniciado con PID: %d y estado en espera", process.Process.Pid))
}

// función para eliminar un proceso mediante el pid.
// func KillProcess(c *fiber.Ctx) error {
// 	var body Model.RequestBody
// 	if err := c.BodyParser(&body); err != nil {
// 		return c.Status(fiber.StatusBadRequest).SendString("Error al parsear el cuerpo de la solicitud")
// 	}

// 	pid := body.PID
// 	if pid == 0 {
// 		return c.Status(fiber.StatusBadRequest).SendString("El parámetro 'pid' es requerido y debe ser un número entero válido")
// 	}

// 	// Enviar señal SIGKILL (9) al proceso con el PID proporcionado
// 	cmd := exec.Command("kill", "-9", strconv.Itoa(pid))
// 	err := cmd.Run()
// 	if err != nil {
// 		return c.Status(fiber.StatusInternalServerError).SendString(fmt.Sprintf("Error al intentar terminar el proceso con PID %d", pid))
// 	}

//		return c.SendString(fmt.Sprintf("Proceso con PID %d ha terminado", pid))
//	}
func KillProcess(c *fiber.Ctx) error {
	pidStr := c.Query("pid")
	if pidStr == "" {
		return c.Status(fiber.StatusBadRequest).SendString("Se requiere el parámetro 'pid'")
	}

	pid, err := strconv.Atoi(pidStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("El parámetro 'pid' debe ser un número entero")
	}

	// Enviar señal SIGKILL (9) al proceso con el PID proporcionado
	cmd := exec.Command("kill", "-9", strconv.Itoa(pid))
	err = cmd.Run()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(fmt.Sprintf("Error al intentar terminar el proceso con PID %d", pid))
	}

	return c.SendString(fmt.Sprintf("Proceso con PID %d ha terminado", pid))
}

// func KillProcess(c *fiber.Ctx) error {
// 	pidStr := c.Query("pid")
// 	if pidStr == "" {
// 		return c.Status(fiber.StatusBadRequest).SendString("Se requiere el parámetro 'pid'")
// 	}

// 	pid, err := strconv.Atoi(pidStr)
// 	if err != nil {
// 		return c.Status(fiber.StatusBadRequest).SendString("El parámetro 'pid' debe ser un número entero")
// 	}

// 	// Enviar señal SIGKILL (9) al proceso con el PID proporcionado
// 	cmd := exec.Command("kill", "-9", strconv.Itoa(pid))
// 	err = cmd.Run()
// 	if err != nil {
// 		return c.Status(fiber.StatusInternalServerError).SendString(fmt.Sprintf("Error al intentar terminar el proceso con PID %d", pid))
// 	}

// 	return c.SendString(fmt.Sprintf("Proceso con PID %d ha terminado", pid))
// }
