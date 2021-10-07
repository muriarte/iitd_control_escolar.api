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
	"iitd_control_escolar.api/usecase/materia"
)

func listMaterias(service materia.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var data []*entity.Materia
		var err error
		name := r.URL.Query().Get("name")
		switch {
		case name == "":
			data, err = service.ListMaterias()
		default:
			data, err = service.SearchMaterias(name)
		}
		if err != nil { //&& err != entity.ErrNotFound {
			sendErrorResponse(w, http.StatusInternalServerError, "error reading records from database", err)
			return
		}

		if data == nil {
			sendOkResponse(w, http.StatusNotFound, []*presenter.Materia{})
			return
		}
		var toJ []*presenter.Materia
		for _, d := range data {
			toJ = append(toJ, &presenter.Materia{
				ID:            d.ID,
				Nombre:        d.Nombre,
				Observaciones: d.Observaciones,
				Activo:        d.Activo,
			})
		}
		sendOkResponse(w, http.StatusOK, toJ)
	})
}

func createMateria(service materia.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var input struct {
			ID            int    `json:"id"`
			Nombre        string `json:"nombre"`
			Observaciones string `json:"observaciones"`
			Activo        string `json:"activo"`
		}
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			log.Println(err.Error())
			sendErrorResponse(w, http.StatusBadRequest, err_decoding_body, err)
			return
		}
		var id int
		var s *entity.Materia
		if input.ID == 0 {
			id, err = service.CreateMateria(input.Nombre, input.Observaciones, input.Activo)
			if err != nil {
				log.Println(err.Error())
				sendErrorResponse(w, http.StatusInternalServerError, err_db_inserting, err)
				return
			}
		} else {
			id = input.ID
			s, err = entity.NewMateria(input.Nombre, input.Observaciones, input.Activo)
			if err != nil {
				log.Println(err.Error())
				sendErrorResponse(w, http.StatusInternalServerError, err_building_object, err)
				return
			}
			if s == nil {
				log.Println("")
				sendErrorResponse(w, http.StatusInternalServerError, err_unexpected_nil_object, nil)
				return
			}
			s.ID = input.ID
			err = service.UpdateMateria(s)
			if err != nil {
				log.Println(err.Error())
				sendErrorResponse(w, http.StatusInternalServerError, err_db_updating, err)
				return
			}
		}
		toJ := &presenter.Materia{
			ID:            id,
			Nombre:        input.Nombre,
			Observaciones: input.Observaciones,
			Activo:        input.Activo,
		}

		sendOkResponse(w, http.StatusOK, toJ)
	})
}

func getMateria(service materia.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			sendErrorResponse(w, http.StatusBadRequest, err_slug_value, err)
			return
		}
		data, err := service.GetMateria(id)
		if err != nil { //&& err != entity.ErrNotFound {
			sendErrorResponse(w, http.StatusInternalServerError, err_db_reading, err)
			return
		}

		if data == nil {
			sendOkResponse(w, http.StatusNotFound, nil)
			return
		}
		toJ := &presenter.Materia{
			ID:            data.ID,
			Nombre:        data.Nombre,
			Observaciones: data.Observaciones,
			Activo:        data.Activo,
		}
		sendOkResponse(w, http.StatusOK, toJ)
	})
}

func deleteMateria(service materia.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			sendErrorResponse(w, http.StatusBadRequest, err_slug_value, err)
			return
		}
		err = service.DeleteMateria(id)
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

//MakeMateriaHandlers make url handlers
func MakeMateriaHandlersNegroni(r *mux.Router, n negroni.Negroni, service materia.UseCase) {
	r.Handle("/v1/materias", n.With(
		negroni.Wrap(listMaterias(service)),
	)).Methods("GET", "OPTIONS").Name("listMaterias")

	r.Handle("/v1/materias", n.With(
		negroni.Wrap(createMateria(service)),
	)).Methods("POST", "OPTIONS").Name("createMateria")

	r.Handle("/v1/materias/{id}", n.With(
		negroni.Wrap(getMateria(service)),
	)).Methods("GET", "OPTIONS").Name("getMateria")

	r.Handle("/v1/materias/{id}", n.With(
		negroni.Wrap(deleteMateria(service)),
	)).Methods("DELETE").Name("deleteMateria")
}

func MakeMateriaHandlers(r *mux.Router, service materia.UseCase) {
	r.Handle("/v1/materias", listMaterias(service)).Methods("GET", "OPTIONS").Name("listMaterias")

	r.Handle("/v1/materias", createMateria(service)).Methods("POST", "OPTIONS").Name("createMateria")

	r.Handle("/v1/materias/{id}", getMateria(service)).Methods("GET", "OPTIONS").Name("getMateria")

	r.Handle("/v1/materias/{id}", deleteMateria(service)).Methods("DELETE").Name("deleteMateria")
}
