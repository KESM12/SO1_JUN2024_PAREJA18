package Model

// Define la estructura de la memoria RAM
type Ram struct {
	Total      int `json:"totalRam"`
	En_uso     int `json:"memoriaEnUso"`
	Libre      int `json:"libre"`
	Porcentaje int `json:"porcentaje"`
}

// Process representa un proceso individual.
type Hijos struct {
	PID      int    `json:"pid"`
	Name     string `json:"name"`
	State    int    `json:"state"`
	PIDPadre int    `json:"pidPadre,omitempty"` // Se omite si está vacío.
}

// ParentProcess representa un proceso padre que contiene procesos hijos.
type Process struct { //DataProcess
	PID   int     `json:"pid"`
	Name  string  `json:"name"`
	State int     `json:"state"`
	Child []Hijos `json:"child"`
}

// Processes representa la estructura principal que contiene una lista de procesos padres.
type Cpu struct {
	Porcentaje int       `json:"cpu_porcentaje"`
	Processes  []Process `json:"processes"`
}
