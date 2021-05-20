package student

import (
	"fmt"
	"iitd_control_escolar.api/entity"
	"strings"
)

//inmem in memory repo
type inmem struct {
	m map[entity.ID]*entity.Student
}

//newInmem create new repository
func newInmem() *inmem {
	var m = map[entity.ID]*entity.Student{}
	return &inmem{
		m: m,
	}
}

//Create an student
func (r *inmem) Create(e *entity.Student) (entity.ID, error) {
	r.m[e.ID] = e
	return e.ID, nil
}

//Get an student
func (r *inmem) Get(id entity.ID) (*entity.Student, error) {
	if r.m[id] == nil {
		return nil, entity.ErrNotFound
	}
	return r.m[id], nil
}

//Update an student
func (r *inmem) Update(e *entity.Student) error {
	_, err := r.Get(e.ID)
	if err != nil {
		return err
	}
	r.m[e.ID] = e
	return nil
}

//Search students
func (r *inmem) Search(query string) ([]*entity.Student, error) {
	var d []*entity.Student
	for _, j := range r.m {
		if strings.Contains(strings.ToLower(j.Nombres), query) {
			d = append(d, j)
		}
	}
	if len(d) == 0 {
		return nil, entity.ErrNotFound
	}

	return d, nil
}

//List students
func (r *inmem) List() ([]*entity.Student, error) {
	var d []*entity.Student
	for _, j := range r.m {
		d = append(d, j)
	}
	return d, nil
}

//Delete an student
func (r *inmem) Delete(id entity.ID) error {
	if r.m[id] == nil {
		return fmt.Errorf("not found")
	}
	r.m[id] = nil
	return nil
}
