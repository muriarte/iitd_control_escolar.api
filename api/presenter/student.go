package presenter

import (
	"iitd_control_escolar.api/entity"
	jd "iitd_control_escolar.api/pkg/jsondate"
)

//Student data
type Student struct {
	ID            entity.ID   `json:"id"`
	Nombres       string      `json:"nombres"`
	Apellidos     string      `json:"apellidos"`
	Nacimiento    jd.JsonDate `json:"nacimiento"`
	Sexo          string      `json:"sexo"`
	Calle         string      `json:"calle"`
	NumeroExt     string      `json:"numeroExt"`
	NumeroInt     string      `json:"numeroInt"`
	Colonia       string      `json:"colonia"`
	Municipio     string      `json:"municipio"`
	Estado        string      `json:"estado"`
	Pais          string      `json:"pais"`
	CP            string      `json:"cp"`
	TelCelular    string      `json:"telCelular"`
	TelCasa       string      `json:"telCasa"`
	Email         string      `json:"email"`
	FechaInicio   jd.JsonDate `json:"fechaInicio"`
	Observaciones string      `json:"observaciones"`
	Activo        string      `json:"activo"`
}
