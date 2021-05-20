package student

import (
	"iitd_control_escolar.api/entity"
	"strings"
	"time"
)

//Service  interface
type Service struct {
	repo Repository
}

//NewService create new use case
func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

//CreateStudent Create an student
func (s *Service) CreateStudent(nombres, apellidos string, nacimiento time.Time, sexo, calle, numeroExt, numeroInt, colonia,
	municipio, estado, pais, cp, telCelular, telCasa, email string, fechaInicio time.Time,
	observaciones, activo string) (entity.ID, error) {
	e, err := entity.NewStudent(nombres, apellidos, nacimiento, sexo, calle, numeroExt, numeroInt, colonia,
		municipio, estado, pais, cp, telCelular, telCasa, email, fechaInicio, observaciones, activo)
	if err != nil {
		return e.ID, err
	}
	return s.repo.Create(e)
}

//GetStudent Get an student
func (s *Service) GetStudent(id entity.ID) (*entity.Student, error) {
	return s.repo.Get(id)
}

//SearchStudents Search students
func (s *Service) SearchStudents(query string) ([]*entity.Student, error) {
	return s.repo.Search(strings.ToLower(query))
}

//ListStudents List students
func (s *Service) ListStudents() ([]*entity.Student, error) {
	return s.repo.List()
}

//DeleteStudent Delete an student
func (s *Service) DeleteStudent(id entity.ID) error {
	u, err := s.GetStudent(id)
	if u == nil {
		return entity.ErrNotFound
	}
	if err != nil {
		return err
	}
	//if len(u.Books) > 0 {
	//	return entity.ErrCannotBeDeleted
	//}
	return s.repo.Delete(id)
}

//UpdateStudent Update an student
func (s *Service) UpdateStudent(e *entity.Student) error {
	err := e.Validate()
	if err != nil {
		return entity.ErrInvalidEntity
	}
	e.UpdatedAt = time.Now()
	return s.repo.Update(e)
}
