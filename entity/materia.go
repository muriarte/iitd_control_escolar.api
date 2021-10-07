package entity

//Student data
type Materia struct {
	ID            int
	Nombre        string
	Observaciones string
	Activo        string
}

//NewMateria create a new student
func NewMateria(nombre, observaciones string, activo string) (*Materia, error) {
	u := &Materia{
		ID:            0,
		Nombre:        nombre,
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
func (u *Materia) Validate() error {
	if u.Nombre == "" {
		return ErrInvalidEntity
	}

	return nil
}
