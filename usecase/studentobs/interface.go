package studentobs

import (
	"time"

	"iitd_control_escolar.api/entity"
)

//Reader interface
type Reader interface {
	Get(id int) (*entity.StudentObs, error)
	Search(studentId int) ([]*entity.StudentObs, error)
	List() ([]*entity.StudentObs, error)
}

//Writer observaciones writer
type Writer interface {
	Create(e *entity.StudentObs) (int, error)
	Update(e *entity.StudentObs) error
	Delete(id int) error
}

//Repository interface
type Repository interface {
	Reader
	Writer
}

//UseCase interface
type UseCase interface {
	GetStudentObs(id int) (*entity.StudentObs, error)
	SearchStudentObs(studentId int) ([]*entity.StudentObs, error)
	ListStudentObs() ([]*entity.StudentObs, error)
	CreateStudentObs(studentId int, fecha time.Time, observacion string) (int, error)
	UpdateStudentObs(e *entity.StudentObs) error
	DeleteStudentObs(id int) error
}
