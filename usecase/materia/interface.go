package materia

import (
	"iitd_control_escolar.api/entity"
)

//Reader interface
type Reader interface {
	Get(id int) (*entity.Materia, error)
	Search(query string) ([]*entity.Materia, error)
	List() ([]*entity.Materia, error)
}

//Writer materia writer
type Writer interface {
	Create(e *entity.Materia) (int, error)
	Update(e *entity.Materia) error
	Delete(id int) error
}

//Repository interface
type Repository interface {
	Reader
	Writer
}

//UseCase interface
type UseCase interface {
	GetMateria(id int) (*entity.Materia, error)
	SearchMaterias(query string) ([]*entity.Materia, error)
	ListMaterias() ([]*entity.Materia, error)
	CreateMateria(nombre, observaciones, activo string) (int, error)
	UpdateMateria(e *entity.Materia) error
	DeleteMateria(id int) error
}
