package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"strconv"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type Datacpu struct {
	Fecha      string `json:"fecha"`
	Porcentaje int    `json:"porcentaje"`
}

type Datacpu2 struct {
	Fecha      string `json:"fecha"`
	Porcentaje int    `json:"porcentaje"`
}

type child struct {
	Pid    int    `json:"pid"`
	Nombre string `json:"name"`
	Estado int    `json:"state"`
	Padre  int    `json:"pidPadre"`
}

type Process struct {
	Pid     int     `json:"pid"`
	Nombre  string  `json:"name"`
	Usuario int     `json:"usuario"`
	Estado  int     `json:"state"`
	Ram     int     `json:"ram"`
	Padre   int     `json:"pidPadre"`
	Chil    []child `json:"child"`
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
type Dataram struct {
	Fecha      string `json:"fecha"`
	Porcentaje int    `json:"porcentaje"`
}

type Ip struct {
	Ip string `json:"ip"`
}
type Respuestacpu struct {
	Mensaje string `json:"mensaje"`
}

type Respuestaram struct {
	Mensaje string `json:"mensaje"`
}

func Index(x http.ResponseWriter, w *http.Request) {
	fmt.Fprintf(x, "sserver")

}

var process *exec.Cmd

func StartProcess(w http.ResponseWriter, r *http.Request) {
	// Crear un nuevo proceso con un comando de espera
	cmd := exec.Command("sleep", "infinity")
	err := cmd.Start()
	if err != nil {
		fmt.Print(err)
		http.Error(w, "Error al iniciar el proceso", http.StatusInternalServerError)
		return
	}

	// Obtener el comando con PID
	process = cmd

	fmt.Fprintf(w, "Proceso iniciado con PID: %d y estado en espera", process.Process.Pid)
}

func KillProcess(w http.ResponseWriter, r *http.Request) {
	pidStr := r.URL.Query().Get("pid")
	if pidStr == "" {
		http.Error(w, "Se requiere el parámetro 'pid'", http.StatusBadRequest)
		return
	}

	pid, err := strconv.Atoi(pidStr)
	if err != nil {
		http.Error(w, "El parámetro 'pid' debe ser un número entero", http.StatusBadRequest)
		return
	}

	// Enviar señal SIGCONT al proceso con el PID proporcionado
	cmd := exec.Command("kill", "-9", strconv.Itoa(pid))
	err = cmd.Run()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error al intentar terminar el proceso con PID %d", pid), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Proceso con PID %d ha terminado", pid)
}

func CPUModuleHandler(w http.ResponseWriter, r *http.Request) {
	cmdCpu := exec.Command("sh", "-c", "cat /proc/cpu_so1_1s2024")
	outCpu, err := cmdCpu.CombinedOutput()
	if err != nil {
		fmt.Println("eerror", err)
	}
	var cpu_info Cpu
	err = json.Unmarshal([]byte(outCpu), &cpu_info)
	if err != nil {
		fmt.Println(err)
	}
	//Mandar respuesta
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode((cpu_info))
}

func RAMModuleHandler(w http.ResponseWriter, r *http.Request) {
	cmdRam := exec.Command("sh", "-c", "cat /proc/ram_so1_1s2024")
	outRam, err := cmdRam.CombinedOutput()
	if err != nil {
		fmt.Println("error", err)
	}
	var ram_info Ram
	err = json.Unmarshal([]byte(outRam), &ram_info)
	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode((ram_info))

}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	}).Methods("GET", "POST")
	router.HandleFunc("/api/cpu", CPUModuleHandler).Methods("GET", "POST")
	router.HandleFunc("/api/ram", RAMModuleHandler).Methods("GET", "POST")

	router.HandleFunc("/api/start", StartProcess)
	router.HandleFunc("/api/kill", KillProcess)

	go func() {
		log.Fatal(http.ListenAndServe(":5200", handlers.CORS()(router)))
	}()
	interval := 5

	ticker := time.NewTicker(time.Second * time.Duration(interval))
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			now := time.Now()
			second1 := now.Second()
			re := strconv.Itoa(now.Year()) + "-" + fmt.Sprintf("%02d", now.Month()) + "-" + fmt.Sprintf("%02d", now.Day()) + "-" + fmt.Sprintf("%02d", now.Hour()) + "-" + fmt.Sprintf("%02d", now.Minute()) + "-" + fmt.Sprintf("%02d", second1)
			fmt.Println("re")
			fmt.Println(re)
			cmdCpu := exec.Command("sh", "-c", "cat /proc/cpu_so1_jun2024")
			outCpu, err := cmdCpu.CombinedOutput()
			if err != nil {
				fmt.Println("error", err)
				return
			}

			//---CPU
			fmt.Println("-------------------- CPU --------------------")
			var cpuInfo Cpu
			err = json.Unmarshal(outCpu, &cpuInfo)
			if err != nil {
				fmt.Println(err)
				return
			}

			fmt.Println("porcentaje:", cpuInfo.Porcentaje)
			fmt.Println("cpu informacion de processos: ", cpuInfo.Procesos)
			//controller.insertdata(Datacpu{"CPU%":  Porcentaje: cpuInfo.Porcentaje})

			//Mandar respuesta (ejemplo simple de uso)
			for _, proceso := range cpuInfo.Procesos {
				// fmt.Printf("Process ID: %d, Name: %s, User: %d, State: %d, RAM: %d\n",
				// 	proceso.Pid, proceso.Nombre, proceso.Usuario, proceso.Estado, proceso.Ram)
				for _, child := range proceso.Chil {
					fmt.Printf("  Child Process ID: %d, Name: %s, State: %d, Parent ID: %d\n",
						child.Pid, child.Nombre, child.Estado, child.Padre)
				}
			}
			go sendToAPI("/cpu", cpuInfo.Porcentaje)
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
			go sendToAPI("/ram", ram_info.Porcentaje)
			fmt.Println(ram_info.Porcentaje)
			fmt.Println(ram_info.En_uso)

		}
	}
}

func sendToAPI(route string, data interface{}) {

	url := fmt.Sprintf("http://localhost:5200/api%s", route)

	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Println("Error al convertir datos a JSON:", err)
		return
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Println("Error al enviar datos a la API:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Println("La API respondió con un código de estado no válido:", resp.StatusCode)
		return
	}

	log.Printf("Datos enviados a la ruta %s\n", route)
}
