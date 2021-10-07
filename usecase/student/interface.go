package student

import (
	"time"

	"iitd_control_escolar.api/entity"
)

//Reader interface
type Reader interface {
	Get(id int) (*entity.Student, error)
	Search(query string) ([]*entity.Student, error)
	List() ([]*entity.Student, error)
}

//Writer student writer
type Writer interface {
	Create(e *entity.Student) (int, error)
	Update(e *entity.Student) error
	Delete(id int) error
}

//Repository interface
type Repository interface {
	Reader
	Writer
}

//UseCase interface
type UseCase interface {
	GetStudent(id int) (*entity.Student, error)
	SearchStudents(query string) ([]*entity.Student, error)
	ListStudents() ([]*entity.Student, error)
	CreateStudent(nombres, apellidos string, nacimiento time.Time, sexo, calle, numeroExt, numeroInt, colonia,
		municipio, estado, pais, cp, telCelular, telCasa, email string, fechaInicio time.Time,
		observaciones, activo string) (int, error)
	UpdateStudent(e *entity.Student) error
	DeleteStudent(id int) error
}
