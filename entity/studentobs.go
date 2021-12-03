package entity

import "time"

//Student data
type StudentObs struct {
	ID          int
	StudentId   int
	Fecha       time.Time
	Observacion string
}

//NewStudentObs create a new student
func NewStudentObs(studentId int, Fecha time.Time, observacion string) (*StudentObs, error) {
	u := &StudentObs{
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
func (u *StudentObs) Validate() error {
	if u.StudentId <= 0 || u.Fecha.Year() <= 2000 || u.Observacion == "" {
		return ErrInvalidEntity
	}

	return nil
}
