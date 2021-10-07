package presenter

import "time"

//Student data
type StudentMateria struct {
	ID            int       `json:"id"`
	StudentId     int       `json:"studentId"`
	MateriaId     int       `json:"materiaId"`
	Inicio        time.Time `json:"inicio"`
	Fin           time.Time `json:"fin"`
	Observaciones string    `json:"observaciones"`
}
