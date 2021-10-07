package repository

import (
	"fmt"
	"strings"

	"iitd_control_escolar.api/entity"
)

//inmem in memory repo
type inmem struct {
	m map[int]*entity.Student
}

//NewStudentInmem create new repository
func NewStudentInmem() *inmem {
	var m = map[int]*entity.Student{}
	return &inmem{
		m: m,
	}
}

//Create an student
func (r *inmem) Create(e *entity.Student) (int, error) {
	// Gets a new ID
	var max = 0
	for k := range r.m {
		if k > max {
			max = k
		}
	}
	max += 1

	e.ID = max
	r.m[e.ID] = e
	return e.ID, nil
}

//Get an student
func (r *inmem) Get(id int) (*entity.Student, error) {
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
func (r *inmem) Delete(id int) error {
	if r.m[id] == nil {
		return fmt.Errorf("not found")
	}

	//if r.m[id].Books.length >0 {
	//	return "cannot Be Deleted"
	//}

	r.m[id] = nil
	return nil
}
