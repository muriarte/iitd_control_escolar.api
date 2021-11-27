package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"iitd_control_escolar.api/api/presenter"
	"iitd_control_escolar.api/entity"
	jd "iitd_control_escolar.api/pkg/jsondate"
	"iitd_control_escolar.api/usecase/studentmateria"
)

func listStudentMaterias(service studentmateria.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var data []*entity.StudentMateria
		var err error
		var studentId int
		var materiaId int
		studentIdStr := r.URL.Query().Get("studentId")
		materiaIdStr := r.URL.Query().Get("materiaId")
		switch {
		case studentIdStr == "" && materiaIdStr == "":
			data, err = service.ListStudentMaterias()
		default:
			if studentIdStr != "" {
				studentId, err = strconv.Atoi(studentIdStr)
				if err != nil {
					sendErrorResponse(w, http.StatusBadRequest, err_query_param_value, err)
					return
				}
			}
			if materiaIdStr != "" {
				materiaId, err = strconv.Atoi(materiaIdStr)
				if err != nil {
					sendErrorResponse(w, http.StatusBadRequest, err_query_param_value, err)
					return
				}
			}
			data, err = service.SearchStudentMaterias(studentId, materiaId)
		}
		if err != nil && err != entity.ErrNotFound {
			sendErrorResponse(w, http.StatusInternalServerError, err_db_reading, err)
			return
		}

		if data == nil {
			sendOkResponse(w, http.StatusOK, nil)
			return
		}
		var toJ []*presenter.StudentMateria
		for _, d := range data {
			toJ = append(toJ, &presenter.StudentMateria{
				ID:            d.ID,
				StudentId:     d.StudentId,
				MateriaId:     d.MateriaId,
				MateriaNombre: d.MateriaNombre,
				Inicio:        d.Inicio,
				Fin:           d.Fin,
				Observaciones: d.Observaciones,
			})
		}
		sendOkResponse(w, http.StatusOK, toJ)
	})
}

func createStudentMateria(service studentmateria.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var input struct {
			ID            int         `json:"id"`
			StudentId     int         `json:"studentId"`
			MateriaId     int         `json:"materiaId"`
			Inicio        jd.JsonDate `json:"inicio"`
			Fin           jd.JsonDate `json:"fin"`
			Observaciones string      `json:"observaciones"`
		}
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			log.Println(err.Error())
			sendErrorResponse(w, http.StatusBadRequest, err_decoding_body, err)
			return
		}
		var id int
		var s *entity.StudentMateria
		var materiaNombre string
		if input.ID == 0 {
			id, materiaNombre, err = service.CreateStudentMateria(input.StudentId, input.MateriaId, input.Inicio.ToTime(), input.Fin.ToTime(), input.Observaciones)
			if err != nil {
				sendErrorResponse(w, http.StatusInternalServerError, err_db_inserting, err)
				return
			}
		} else {
			id = input.ID
			s, err = entity.NewStudentMateria(input.StudentId, input.MateriaId, input.Inicio.ToTime(), input.Fin.ToTime(), input.Observaciones)
			if err != nil {
				sendErrorResponse(w, http.StatusInternalServerError, err_building_object, err)
				return
			}
			if s == nil {
				sendErrorResponse(w, http.StatusInternalServerError, err_unexpected_nil_object, err)
				return
			}
			s.ID = input.ID
			materiaNombre, err = service.UpdateStudentMateria(s)
			if err != nil {
				sendErrorResponse(w, http.StatusInternalServerError, err_db_updating, err)
				return
			}
		}
		toJ := &presenter.StudentMateria{
			ID:            id,
			StudentId:     input.StudentId,
			MateriaId:     input.MateriaId,
			MateriaNombre: materiaNombre,
			Inicio:        input.Inicio.ToTime(),
			Fin:           input.Fin.ToTime(),
			Observaciones: input.Observaciones,
		}
		sendOkResponse(w, http.StatusOK, toJ)
	})
}

func getStudentMateria(service studentmateria.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			sendErrorResponse(w, http.StatusBadRequest, err_slug_value, err)
			return
		}
		data, err := service.GetStudentMateria(id)
		if err != nil && err != entity.ErrNotFound {
			sendErrorResponse(w, http.StatusInternalServerError, err_db_reading, err)
			return
		}

		if data == nil {
			sendOkResponse(w, http.StatusNotFound, nil)
			return
		}
		toJ := &presenter.StudentMateria{
			ID:            data.ID,
			StudentId:     data.StudentId,
			MateriaId:     data.MateriaId,
			MateriaNombre: data.MateriaNombre,
			Inicio:        data.Inicio,
			Fin:           data.Fin,
			Observaciones: data.Observaciones,
		}
		sendOkResponse(w, http.StatusOK, toJ)
	})
}

func deleteStudentMateria(service studentmateria.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			sendErrorResponse(w, http.StatusBadRequest, err_slug_value, err)
			return
		}
		err = service.DeleteStudentMateria(id)
		if err != nil {
			sendErrorResponse(w, http.StatusInternalServerError, err_db_deleting, err)
			return
		}

		ret := struct {
			Status bool `json:"status"`
		}{
			Status: true,
		}
		sendOkResponse(w, http.StatusOK, ret)
	})
}

//MakeStudentMateriaHandlers make url handlers
func MakeStudentMateriaHandlersNegroni(r *mux.Router, n negroni.Negroni, service studentmateria.UseCase) {
	r.Handle("/v1/studentmaterias", n.With(
		negroni.Wrap(listStudentMaterias(service)),
	)).Methods("GET", "OPTIONS").Name("listStudentMaterias")

	r.Handle("/v1/studentmaterias", n.With(
		negroni.Wrap(createStudentMateria(service)),
	)).Methods("POST", "OPTIONS").Name("createStudentMateria")

	r.Handle("/v1/studentmaterias/{id}", n.With(
		negroni.Wrap(getStudentMateria(service)),
	)).Methods("GET", "OPTIONS").Name("getStudentMateria")

	r.Handle("/v1/studentmaterias/{id}", n.With(
		negroni.Wrap(deleteStudentMateria(service)),
	)).Methods("DELETE").Name("deleteStudentMateria")
}

func MakeStudentMateriaHandlers(r *mux.Router, service studentmateria.UseCase) {
	r.Handle("/v1/studentmaterias", listStudentMaterias(service)).Methods("GET", "OPTIONS").Name("listStudentMaterias")

	r.Handle("/v1/studentmaterias", createStudentMateria(service)).Methods("POST", "OPTIONS").Name("createStudentMateria")

	r.Handle("/v1/studentmaterias/{id}", getStudentMateria(service)).Methods("GET", "OPTIONS").Name("getStudentMateria")

	r.Handle("/v1/studentmaterias/{id}", deleteStudentMateria(service)).Methods("DELETE").Name("deleteStudentMateria")
}
