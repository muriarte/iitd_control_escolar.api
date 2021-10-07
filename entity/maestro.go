package entity

import (
	"time"
)

//Student data
type Maestro struct {
	ID            int
	Nombres       string
	Apellidos     string
	Nacimiento    time.Time
	Sexo          string
	Calle         string
	NumeroExt     string
	NumeroInt     string
	Colonia       string
	Municipio     string
	Estado        string
	Pais          string
	CP            string
	TelCelular    string
	TelCasa       string
	Email         string
	FechaInicio   time.Time
	Observaciones string
	Activo        string
}

//NewMaestro create a new student
func NewMaestro(nombres, apellidos string, nacimiento time.Time, sexo, calle, numeroExt, numeroInt, colonia,
	municipio, estado, pais, cp, telCelular, telCasa, email string, fechaInicio time.Time,
	observaciones, activo string) (*Maestro, error) {
	u := &Maestro{
		ID:            0,
		Nombres:       nombres,
		Apellidos:     apellidos,
		Nacimiento:    nacimiento,
		Sexo:          sexo,
		Calle:         calle,
		NumeroExt:     numeroExt,
		NumeroInt:     numeroInt,
		Colonia:       colonia,
		Municipio:     municipio,
		Estado:        estado,
		Pais:          pais,
		CP:            cp,
		TelCelular:    telCelular,
		TelCasa:       telCasa,
		Email:         email,
		FechaInicio:   fechaInicio,
		Observaciones: observaciones,
		Activo:        activo,
	}
	err := u.Validate()
	if err != nil {
		return nil, ErrInvalidEntity
	}
	return u, nil
}

//Validate validate data
func (u *Maestro) Validate() error {
	if u.Nombres == "" {
		return ErrInvalidEntity
	}

	return nil
}
