package Model

type Ram struct {
	Total      string `json:"totalRam"`
	En_uso     string `json:"memoriaEnUso"`
	Libre      string `json:"libre"`
	Porcentaje string `json:"porcentaje"`
}

// Define la estructura del proceso
type Process struct {
	Pid    int     `json:"pid"`
	Nombre string  `json:"name"`
	Estado int     `json:"state"`
	Padre  int     `json:"pidPadre"`
	Chil   []Child `json:"child"`
}

type Child struct {
	Pid    int    `json:"pid"`
	Nombre string `json:"name"`
	Estado int    `json:"state"`
}

type CPUInfo struct {
	Processes []Process `json:"processes"`
}
