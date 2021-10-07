package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"iitd_control_escolar.api/api/presenter"
	"iitd_control_escolar.api/entity"
	jd "iitd_control_escolar.api/pkg/jsondate"
	"iitd_control_escolar.api/usecase/observacion"
)

func listObservaciones(service observacion.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var data []*entity.Observacion
		var err error
		var studentId int
		studentIdStr := r.URL.Query().Get("studentId")
		switch {
		case studentIdStr == "":
			data, err = service.ListObservaciones()
		default:
			studentId, err = strconv.Atoi(studentIdStr)
			if err != nil {
				sendErrorResponse(w, http.StatusBadRequest, err_query_param_value, err)
				return
			}
			data, err = service.SearchObservaciones(studentId)
		}
		if err != nil && err != entity.ErrNotFound {
			sendErrorResponse(w, http.StatusInternalServerError, err_db_reading, err)
			return
		}

		if data == nil {
			sendOkResponse(w, http.StatusNotFound, []*presenter.Observacion{})
			return
		}
		var toJ []*presenter.Observacion
		for _, d := range data {
			toJ = append(toJ, &presenter.Observacion{
				ID:          d.ID,
				StudentId:   d.StudentId,
				Observacion: d.Observacion,
			})
		}
		sendOkResponse(w, http.StatusOK, toJ)
	})
}

func createObservacion(service observacion.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var input struct {
			ID          int         `json:"id"`
			StudentId   int         `json:"studentId"`
			Fecha       jd.JsonDate `json:"fecha"`
			Observacion string      `json:"observacion"`
		}
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			sendErrorResponse(w, http.StatusBadRequest, err_decoding_body, err)
			return
		}
		var id int
		var s *entity.Observacion
		if input.ID == 0 {
			id, err = service.CreateObservacion(input.StudentId, input.Fecha.ToTime(), input.Observacion)
			if err != nil {
				sendErrorResponse(w, http.StatusInternalServerError, err_db_inserting, err)
				return
			}
		} else {
			id = input.ID
			s, err = entity.NewObservacion(input.StudentId, input.Fecha.ToTime(), input.Observacion)
			if err != nil {
				sendErrorResponse(w, http.StatusInternalServerError, err_building_object, err)
				return
			}
			if s == nil {
				sendOkResponse(w, http.StatusNotFound, nil)
				return
			}
			s.ID = input.ID
			err = service.UpdateObservacion(s)
			if err != nil {
				sendErrorResponse(w, http.StatusInternalServerError, err_db_updating, err)
				return
			}
		}
		toJ := &presenter.Observacion{
			ID:          id,
			StudentId:   input.StudentId,
			Fecha:       input.Fecha,
			Observacion: input.Observacion,
		}
		sendOkResponse(w, http.StatusOK, toJ)
	})
}

func getObservacion(service observacion.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			sendErrorResponse(w, http.StatusBadRequest, err_slug_value, err)
			return
		}
		data, err := service.GetObservacion(id)
		if err != nil && err != entity.ErrNotFound {
			sendErrorResponse(w, http.StatusInternalServerError, err_db_reading, err)
			return
		}

		if data == nil {
			sendOkResponse(w, http.StatusNotFound, nil)
			return
		}
		toJ := &presenter.Observacion{
			ID:          data.ID,
			StudentId:   data.StudentId,
			Fecha:       jd.JsonDate(data.Fecha),
			Observacion: data.Observacion,
		}
		sendOkResponse(w, http.StatusOK, toJ)
	})
}

func deleteObservacion(service observacion.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			sendErrorResponse(w, http.StatusBadRequest, err_slug_value, err)
			return
		}
		err = service.DeleteObservacion(id)
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

//MakeObservacionHandlers make url handlers
func MakeObservacionHandlersNegroni(r *mux.Router, n negroni.Negroni, service observacion.UseCase) {
	r.Handle("/v1/observaciones", n.With(
		negroni.Wrap(listObservaciones(service)),
	)).Methods("GET", "OPTIONS").Name("listObservaciones")

	r.Handle("/v1/observaciones", n.With(
		negroni.Wrap(createObservacion(service)),
	)).Methods("POST", "OPTIONS").Name("createObservacion")

	r.Handle("/v1/observaciones/{id}", n.With(
		negroni.Wrap(getObservacion(service)),
	)).Methods("GET", "OPTIONS").Name("getObservacion")

	r.Handle("/v1/observaciones/{id}", n.With(
		negroni.Wrap(deleteObservacion(service)),
	)).Methods("DELETE").Name("deleteObservacion")
}

func MakeObservacionHandlers(r *mux.Router, service observacion.UseCase) {
	r.Handle("/v1/observaciones", listObservaciones(service)).Methods("GET", "OPTIONS").Name("listObservaciones")

	r.Handle("/v1/observaciones", createObservacion(service)).Methods("POST", "OPTIONS").Name("createObservacion")

	r.Handle("/v1/observaciones/{id}", getObservacion(service)).Methods("GET", "OPTIONS").Name("getObservacion")

	r.Handle("/v1/observaciones/{id}", deleteObservacion(service)).Methods("DELETE").Name("deleteObservacion")
}
