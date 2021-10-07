package entity

import "time"

//Student data
type Observacion struct {
	ID          int
	StudentId   int
	Fecha       time.Time
	Observacion string
}

//NewObservacion create a new student
func NewObservacion(studentId int, Fecha time.Time, observacion string) (*Observacion, error) {
	u := &Observacion{
		ID:          0,
		StudentId:   studentId,
		Fecha:       Fecha,
		Observacion: observacion,
	}
	err := u.Validate()
	if err != nil {
		return nil, ErrInvalidEntity
	}
	return u, nil
}

//Validate validate data
func (u *Observacion) Validate() error {
	if u.StudentId > 0 {
		return ErrInvalidEntity
	}

	return nil
}
