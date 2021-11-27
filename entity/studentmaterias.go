package entity

import "time"

//Student data
type StudentMateria struct {
	ID            int
	StudentId     int
	MateriaId     int
	MateriaNombre string
	Inicio        time.Time
	Fin           time.Time
	Observaciones string
}

//NewStudentMateria create a new student-materia
func NewStudentMateria(studentId int, materiaId int, inicio, fin time.Time, observaciones string) (*StudentMateria, error) {
	u := &StudentMateria{
		ID:            0,
		StudentId:     studentId,
		MateriaId:     materiaId,
		MateriaNombre: "",
		Inicio:        inicio,
		Fin:           fin,
		Observaciones: observaciones,
	}
	err := u.Validate()
	if err != nil {
		return nil, ErrInvalidEntity
	}
	return u, nil
}

//Validate validate data
func (u *StudentMateria) Validate() error {
	if u.StudentId == 0 || u.MateriaId == 0 {
		return ErrInvalidEntity
	}

	return nil
}
