package maestro

import (
	"strings"
	"time"

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

//CreateMaestro Create an maestro
func (s *Service) CreateMaestro(nombres, apellidos string, nacimiento time.Time, sexo, calle, numeroExt, numeroInt, colonia,
	municipio, estado, pais, cp, telCelular, telCasa, email string, fechaInicio time.Time,
	observaciones, activo string) (int, error) {
	e, err := entity.NewMaestro(nombres, apellidos, nacimiento, sexo, calle, numeroExt, numeroInt, colonia,
		municipio, estado, pais, cp, telCelular, telCasa, email, fechaInicio, observaciones, activo)
	if err != nil {
		return e.ID, err
	}
	return s.repo.Create(e)
}

//GetMaestro Get an maestro
func (s *Service) GetMaestro(id int) (*entity.Maestro, error) {
	return s.repo.Get(id)
}

//SearchMaestros Search maestros
func (s *Service) SearchMaestros(query string) ([]*entity.Maestro, error) {
	return s.repo.Search(strings.ToLower(query))
}

//ListMaestros List maestros
func (s *Service) ListMaestros() ([]*entity.Maestro, error) {
	return s.repo.List()
}

//DeleteMaestro Delete an maestro
func (s *Service) DeleteMaestro(id int) error {
	u, err := s.GetMaestro(id)
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

//UpdateMaestro Update an maestro
func (s *Service) UpdateMaestro(e *entity.Maestro) error {
	err := e.Validate()
	if err != nil {
		return entity.ErrInvalidEntity
	}
	//e.UpdatedAt = time.Now()
	return s.repo.Update(e)
}
