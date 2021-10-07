package observacion

import (
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

//CreateObservaciones Create an observaciones
func (s *Service) CreateObservacion(studentId int, fecha time.Time, observacion string) (int, error) {
	e, err := entity.NewObservacion(studentId, fecha, observacion)
	if err != nil {
		return e.ID, err
	}
	return s.repo.Create(e)
}

//GetObservacion Get an observaciones
func (s *Service) GetObservacion(id int) (*entity.Observacion, error) {
	return s.repo.Get(id)
}

//SearchObservacioness Search observacioness
func (s *Service) SearchObservaciones(studentId int) ([]*entity.Observacion, error) {
	return s.repo.Search(studentId)
}

//ListObservacioness List observacioness
func (s *Service) ListObservaciones() ([]*entity.Observacion, error) {
	return s.repo.List()
}

//DeleteObservacion Delete an observaciones
func (s *Service) DeleteObservacion(id int) error {
	u, err := s.GetObservacion(id)
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

//UpdateObservacion Update an observaciones
func (s *Service) UpdateObservacion(e *entity.Observacion) error {
	err := e.Validate()
	if err != nil {
		return entity.ErrInvalidEntity
	}
	//e.UpdatedAt = time.Now()
	return s.repo.Update(e)
}
