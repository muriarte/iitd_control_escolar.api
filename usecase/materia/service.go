package materia

import (
	"strings"

	"iitd_control_escolar.api/entity"
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

//CreateMateria Create an materia
func (s *Service) CreateMateria(nombre, observaciones, activo string) (int, error) {
	e, err := entity.NewMateria(nombre, observaciones, activo)
	if err != nil {
		return e.ID, err
	}
	return s.repo.Create(e)
}

//GetMateria Get an materia
func (s *Service) GetMateria(id int) (*entity.Materia, error) {
	return s.repo.Get(id)
}

//SearchMaterias Search materias
func (s *Service) SearchMaterias(query string) ([]*entity.Materia, error) {
	return s.repo.Search(strings.ToLower(query))
}

//ListMaterias List materias
func (s *Service) ListMaterias() ([]*entity.Materia, error) {
	return s.repo.List()
}

//DeleteMateria Delete an materia
func (s *Service) DeleteMateria(id int) error {
	u, err := s.GetMateria(id)
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

//UpdateMateria Update an materia
func (s *Service) UpdateMateria(e *entity.Materia) error {
	err := e.Validate()
	if err != nil {
		return entity.ErrInvalidEntity
	}
	//e.UpdatedAt = time.Now()
	return s.repo.Update(e)
}
