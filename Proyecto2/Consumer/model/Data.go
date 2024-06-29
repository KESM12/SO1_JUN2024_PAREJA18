package model

import "time"

type Data struct {
	Texto string    `json:"texto"`
	Pais  string    `json:"pais"`
	Fecha time.Time `json:"fecha"`
}
