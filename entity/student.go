package entity

import (
	"time"
)

//Student data
type Student struct {
	ID            ID
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
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

//NewStudent create a new student
func NewStudent(nombres, apellidos string, nacimiento time.Time, sexo, calle, numeroExt, numeroInt, colonia,
	municipio, estado, pais, cp, telCelular, telCasa, email string, fechaInicio time.Time,
	observaciones, activo string) (*Student, error) {
	u := &Student{
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
		CreatedAt:     time.Now(),
	}
	err := u.Validate()
	if err != nil {
		return nil, ErrInvalidEntity
	}
	return u, nil
}

//Validate validate data
func (u *Student) Validate() error {
	if u.Nombres == "" || u.Apellidos == "" {
		return ErrInvalidEntity
	}

	return nil
}
