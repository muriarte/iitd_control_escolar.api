package studentobs

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

//CreateObservaciones Create an StudentObs
func (s *Service) CreateStudentObs(studentId int, fecha time.Time, observacion string) (int, error) {
	e, err := entity.NewStudentObs(studentId, fecha, observacion)
	if err != nil {
		return e.ID, err
	}
	err = e.Validate()
	if err != nil {
		return 0, entity.ErrInvalidEntity
	}
	return s.repo.Create(e)
}

//GetObservacion Get an StudentObs
func (s *Service) GetStudentObs(id int) (*entity.StudentObs, error) {
	return s.repo.Get(id)
}

//SearchObservacioness Search StudentObs
func (s *Service) SearchStudentObs(studentId int) ([]*entity.StudentObs, error) {
	return s.repo.Search(studentId)
}

//ListObservacioness List StudentObs
func (s *Service) ListStudentObs() ([]*entity.StudentObs, error) {
	return s.repo.List()
}

//DeleteObservacion Delete an StudentObs
func (s *Service) DeleteStudentObs(id int) error {
	u, err := s.GetStudentObs(id)
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
func (s *Service) UpdateStudentObs(e *entity.StudentObs) error {
	err := e.Validate()
	if err != nil {
		return entity.ErrInvalidEntity
	}
	//e.UpdatedAt = time.Now()
	return s.repo.Update(e)
}
