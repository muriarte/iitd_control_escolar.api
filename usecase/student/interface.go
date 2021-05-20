package student

import (
	"iitd_control_escolar.api/entity"
	"time"
)

//Reader interface
type Reader interface {
	Get(id entity.ID) (*entity.Student, error)
	Search(query string) ([]*entity.Student, error)
	List() ([]*entity.Student, error)
}

//Writer student writer
type Writer interface {
	Create(e *entity.Student) (entity.ID, error)
	Update(e *entity.Student) error
	Delete(id entity.ID) error
}

//Repository interface
type Repository interface {
	Reader
	Writer
}

//UseCase interface
type UseCase interface {
	GetStudent(id entity.ID) (*entity.Student, error)
	SearchStudents(query string) ([]*entity.Student, error)
	ListStudents() ([]*entity.Student, error)
	CreateStudent(nombres, apellidos string, nacimiento time.Time, sexo, calle, numeroExt, numeroInt, colonia,
		municipio, estado, pais, cp, telCelular, telCasa, email string, fechaInicio time.Time,
		observaciones, activo string) (entity.ID, error)
	UpdateStudent(e *entity.Student) error
	DeleteStudent(id entity.ID) error
}