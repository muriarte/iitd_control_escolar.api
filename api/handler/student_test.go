package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"iitd_control_escolar.api/api/presenter"
	"iitd_control_escolar.api/entity"
	"iitd_control_escolar.api/usecase/student/mock"
)

func Test_listStudents(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	m := mock.NewMockUseCase(controller)
	r := mux.NewRouter()
	MakeStudentHandlers(r, m)
	path, err := r.GetRoute("listStudents").GetPathTemplate()
	assert.Nil(t, err)
	assert.Equal(t, "/v1/students", path)
	u := &entity.Student{
		ID: entity.NewID(),
	}
	m.EXPECT().
		ListStudents().
		Return([]*entity.Student{u}, nil)
	ts := httptest.NewServer(listStudents(m))
	defer ts.Close()
	res, err := http.Get(ts.URL)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
}

func Test_listStudents_NotFound(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	m := mock.NewMockUseCase(controller)
	ts := httptest.NewServer(listStudents(m))
	defer ts.Close()
	m.EXPECT().
		SearchStudents("dio").
		Return(nil, entity.ErrNotFound)
	res, err := http.Get(ts.URL + "?name=dio")
	assert.Nil(t, err)
	assert.Equal(t, http.StatusNotFound, res.StatusCode)
}

func Test_listStudents_Search(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	m := mock.NewMockUseCase(controller)
	u := &entity.Student{
		ID: entity.NewID(),
	}
	m.EXPECT().
		SearchStudents("ozzy").
		Return([]*entity.Student{u}, nil)
	ts := httptest.NewServer(listStudents(m))
	defer ts.Close()
	res, err := http.Get(ts.URL + "?name=ozzy")
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
}

func Test_createStudent(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	m := mock.NewMockUseCase(controller)
	r := mux.NewRouter()
	MakeStudentHandlers(r, m)
	path, err := r.GetRoute("createStudent").GetPathTemplate()
	assert.Nil(t, err)
	assert.Equal(t, "/v1/students", path)

	m.EXPECT().
		CreateStudent(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(),
			gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(),
			gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
		Return(entity.NewID(), nil)
	h := createStudent(m)

	ts := httptest.NewServer(h)
	defer ts.Close()
	payload := `{
"activo": "S",
"email": "pepe@test.com",
"nombres":"Pepe",
"apellidos":"Potamo"
}`
	resp, _ := http.Post(ts.URL+"/v1/student", "application/json", strings.NewReader(payload))
	assert.Equal(t, http.StatusCreated, resp.StatusCode)

	var u *presenter.Student
	_ = json.NewDecoder(resp.Body).Decode(&u)
	assert.Equal(t, "Pepe Potamo", fmt.Sprintf("%s %s", u.Nombres, u.Apellidos))
}

func Test_getStudent(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	m := mock.NewMockUseCase(controller)
	r := mux.NewRouter()
	MakeStudentHandlers(r, m)
	path, err := r.GetRoute("getStudent").GetPathTemplate()
	assert.Nil(t, err)
	assert.Equal(t, "/v1/students/{id}", path)
	u := &entity.Student{
		ID: entity.NewID(),
	}
	m.EXPECT().
		GetStudent(u.ID).
		Return(u, nil)
	handler := getStudent(m)
	r.Handle("/v1/student/{id}", handler)
	ts := httptest.NewServer(r)
	defer ts.Close()
	res, err := http.Get(ts.URL + "/v1/student/" + entity.IDToString(u.ID))
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	var d *presenter.Student
	_ = json.NewDecoder(res.Body).Decode(&d)
	assert.NotNil(t, d)
	assert.Equal(t, u.ID, d.ID)
}

func Test_deleteStudent(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	m := mock.NewMockUseCase(controller)
	r := mux.NewRouter()
	MakeStudentHandlers(r, m)
	path, err := r.GetRoute("deleteStudent").GetPathTemplate()
	assert.Nil(t, err)
	assert.Equal(t, "/v1/students/{id}", path)
	u := &entity.Student{
		ID: entity.NewID(),
	}
	m.EXPECT().DeleteStudent(u.ID).Return(nil)
	handler := deleteStudent(m)
	req, _ := http.NewRequest("DELETE", "/v1/student/"+entity.IDToString(u.ID), nil)
	r.Handle("/v1/student/{id}", handler).Methods("DELETE", "OPTIONS")
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
}
