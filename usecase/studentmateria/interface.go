package studentmateria

import (
	"time"

	"iitd_control_escolar.api/entity"
)

//Reader interface
type Reader interface {
	Get(id int) (*entity.StudentMateria, error)
	Search(studentId int, materiaId int) ([]*entity.StudentMateria, error)
	List() ([]*entity.StudentMateria, error)
}

//Writer studentmateria writer
type Writer interface {
	Create(e *entity.StudentMateria) (int, error)
	Update(e *entity.StudentMateria) error
	Delete(id int) error
}

//Repository interface
type Repository interface {
	Reader
	Writer
}

//UseCase interface
type UseCase interface {
	GetStudentMateria(id int) (*entity.StudentMateria, error)
	SearchStudentMaterias(studentId int, materiaId int) ([]*entity.StudentMateria, error)
	ListStudentMaterias() ([]*entity.StudentMateria, error)
	CreateStudentMateria(studentId, materiaId int, inicio, fin time.Time, observaciones string) (int, error)
	UpdateStudentMateria(e *entity.StudentMateria) error
	DeleteStudentMateria(id int) error
}
