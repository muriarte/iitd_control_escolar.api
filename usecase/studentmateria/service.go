package studentmateria

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

//CreateStudentMateria Create an studentmateria
func (s *Service) CreateStudentMateria(studentId, materiaId int, inicio, fin time.Time, observaciones string) (int, string, error) {
	e, err := entity.NewStudentMateria(studentId, materiaId, inicio, fin, observaciones)
	if err != nil {
		return 0, "", err
	}
	err = e.Validate()
	if err != nil {
		return 0, "", entity.ErrInvalidEntity
	}
	return s.repo.Create(e)
}

//GetStudentMateria Get an studentmateria
func (s *Service) GetStudentMateria(id int) (*entity.StudentMateria, error) {
	return s.repo.Get(id)
}

//SearchStudentMaterias Search studentmaterias
func (s *Service) SearchStudentMaterias(studentId int, materiaId int) ([]*entity.StudentMateria, error) {
	return s.repo.Search(studentId, materiaId)
}

//ListStudentMaterias List studentmaterias
func (s *Service) ListStudentMaterias() ([]*entity.StudentMateria, error) {
	return s.repo.List()
}

//DeleteStudentMateria Delete an studentmateria
func (s *Service) DeleteStudentMateria(id int) error {
	u, err := s.GetStudentMateria(id)
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

//UpdateStudentMateria Update an studentmateria
func (s *Service) UpdateStudentMateria(e *entity.StudentMateria) (string, error) {
	err := e.Validate()
	if err != nil {
		return "", entity.ErrInvalidEntity
	}
	//e.UpdatedAt = time.Now()
	return s.repo.Update(e)
}
