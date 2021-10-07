package observacion

import (
	"time"

	"iitd_control_escolar.api/entity"
)

//Reader interface
type Reader interface {
	Get(id int) (*entity.Observacion, error)
	Search(studentId int) ([]*entity.Observacion, error)
	List() ([]*entity.Observacion, error)
}

//Writer observaciones writer
type Writer interface {
	Create(e *entity.Observacion) (int, error)
	Update(e *entity.Observacion) error
	Delete(id int) error
}

//Repository interface
type Repository interface {
	Reader
	Writer
}

//UseCase interface
type UseCase interface {
	GetObservacion(id int) (*entity.Observacion, error)
	SearchObservaciones(studentId int) ([]*entity.Observacion, error)
	ListObservaciones() ([]*entity.Observacion, error)
	CreateObservacion(studentId int, fecha time.Time, observacion string) (int, error)
	UpdateObservacion(e *entity.Observacion) error
	DeleteObservacion(id int) error
}
