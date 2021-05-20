package presenter

import (
	"iitd_control_escolar.api/entity"
	"time"
)

//Student data
type Student struct {
	ID            entity.ID `json:"id"`
	Nombres       string    `json:"nombres"`
	Apellidos     string    `json:"apellidos"`
	Nacimiento    time.Time `json:"nacimiento"`
	Sexo          string    `json:"sexo"`
	Calle         string    `json:"calle"`
	NumeroExt     string    `json:"numeroExt"`
	NumeroInt     string    `json:"numeroInt"`
	Colonia       string    `json:"colonia"`
	Municipio     string    `json:"municipio"`
	Estado        string    `json:"estado"`
	Pais          string    `json:"pais"`
	CP            string    `json:"cp"`
	TelCelular    string    `json:"telCelular"`
	TelCasa       string    `json:"telCasa"`
	Email         string    `json:"email"`
	FechaInicio   time.Time `json:"fechaInicio"`
	Observaciones string    `json:"observaciones"`
	Activo        string    `json:"activo"`
}
