package maestro

import (
	"time"

	"iitd_control_escolar.api/entity"
)

//Reader interface
type Reader interface {
	Get(id int) (*entity.Maestro, error)
	Search(query string) ([]*entity.Maestro, error)
	List() ([]*entity.Maestro, error)
}

//Writer maestro writer
type Writer interface {
	Create(e *entity.Maestro) (int, error)
	Update(e *entity.Maestro) error
	Delete(id int) error
}

//Repository interface
type Repository interface {
	Reader
	Writer
}

//UseCase interface
type UseCase interface {
	GetMaestro(id int) (*entity.Maestro, error)
	SearchMaestros(query string) ([]*entity.Maestro, error)
	ListMaestros() ([]*entity.Maestro, error)
	CreateMaestro(nombres, apellidos string, nacimiento time.Time, sexo, calle, numeroExt, numeroInt, colonia,
		municipio, estado, pais, cp, telCelular, telCasa, email string, fechaInicio time.Time,
		observaciones, activo string) (int, error)
	UpdateMaestro(e *entity.Maestro) error
	DeleteMaestro(id int) error
}
