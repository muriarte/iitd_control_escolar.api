package presenter

import jd "iitd_control_escolar.api/pkg/jsondate"

//Student data
type StudentMateria struct {
	ID            int         `json:"id"`
	StudentId     int         `json:"studentId"`
	MateriaId     int         `json:"materiaId"`
	MateriaNombre string      `json:"materiaNombre"`
	Inicio        jd.JsonDate `json:"inicio"`
	Fin           jd.JsonDate `json:"fin"`
	Observaciones string      `json:"observaciones"`
}
