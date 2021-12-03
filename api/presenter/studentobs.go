package presenter

import (
	jd "iitd_control_escolar.api/pkg/jsondate"
)

//Student data
type StudentObs struct {
	ID          int         `json:"id"`
	StudentId   int         `json:"studentId"`
	Fecha       jd.JsonDate `json:"fecha"`
	Observacion string      `json:"observacion"`
}
