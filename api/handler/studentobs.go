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
	"iitd_control_escolar.api/usecase/studentobs"
)

func listStudentObs(service studentobs.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var data []*entity.StudentObs
		var err error
		var studentId int
		studentIdStr := r.URL.Query().Get("studentId")
		switch {
		case studentIdStr == "":
			data, err = service.ListStudentObs()
		default:
			if studentIdStr != "" {
				studentId, err = strconv.Atoi(studentIdStr)
				if err != nil {
					sendErrorResponse(w, http.StatusBadRequest, err_query_param_value, err)
					return
				}
			}
			data, err = service.SearchStudentObs(studentId)
		}
		if err != nil && err != entity.ErrNotFound {
			sendErrorResponse(w, http.StatusInternalServerError, err_db_reading, err)
			return
		}

		if data == nil {
			sendOkResponse(w, http.StatusOK, nil)
			return
		}
		var toJ []*presenter.StudentObs
		for _, d := range data {
			toJ = append(toJ, &presenter.StudentObs{
				ID:          d.ID,
				StudentId:   d.StudentId,
				Fecha:       jd.JsonDate(d.Fecha),
				Observacion: d.Observacion,
			})
		}
		sendOkResponse(w, http.StatusOK, toJ)
	})
}

func createStudentObs(service studentobs.UseCase) http.Handler {
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
		var s *entity.StudentObs
		if input.ID == 0 {
			id, err = service.CreateStudentObs(input.StudentId, input.Fecha.ToTime(), input.Observacion)
			if err != nil {
				sendErrorResponse(w, http.StatusInternalServerError, err_db_inserting, err)
				return
			}
		} else {
			id = input.ID
			s, err = entity.NewStudentObs(input.StudentId, input.Fecha.ToTime(), input.Observacion)
			if err != nil {
				sendErrorResponse(w, http.StatusInternalServerError, err_building_object, err)
				return
			}
			if s == nil {
				sendErrorResponse(w, http.StatusInternalServerError, err_unexpected_nil_object, err)
				return
			}
			s.ID = input.ID
			err = service.UpdateStudentObs(s)
			if err != nil {
				sendErrorResponse(w, http.StatusInternalServerError, err_db_updating, err)
				return
			}
		}
		toJ := &presenter.StudentObs{
			ID:          id,
			StudentId:   input.StudentId,
			Fecha:       input.Fecha,
			Observacion: input.Observacion,
		}
		sendOkResponse(w, http.StatusOK, toJ)
	})
}

func getStudentObs(service studentobs.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			sendErrorResponse(w, http.StatusBadRequest, err_slug_value, err)
			return
		}
		data, err := service.GetStudentObs(id)
		if err != nil && err != entity.ErrNotFound {
			sendErrorResponse(w, http.StatusInternalServerError, err_db_reading, err)
			return
		}

		if data == nil {
			sendOkResponse(w, http.StatusNotFound, nil)
			return
		}
		toJ := &presenter.StudentObs{
			ID:          data.ID,
			StudentId:   data.StudentId,
			Fecha:       jd.JsonDate(data.Fecha),
			Observacion: data.Observacion,
		}
		sendOkResponse(w, http.StatusOK, toJ)
	})
}

func deleteStudentObs(service studentobs.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			sendErrorResponse(w, http.StatusBadRequest, err_slug_value, err)
			return
		}
		err = service.DeleteStudentObs(id)
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
func MakeObservacionHandlersNegroni(r *mux.Router, n negroni.Negroni, service studentobs.UseCase) {
	r.Handle("/v1/studentobs", n.With(
		negroni.Wrap(listStudentObs(service)),
	)).Methods("GET", "OPTIONS").Name("listStudentObs")

	r.Handle("/v1/studentobs", n.With(
		negroni.Wrap(createStudentObs(service)),
	)).Methods("POST", "OPTIONS").Name("createStudentObs")

	r.Handle("/v1/studentobs/{id}", n.With(
		negroni.Wrap(getStudentObs(service)),
	)).Methods("GET", "OPTIONS").Name("getStudentObs")

	r.Handle("/v1/studentobs/{id}", n.With(
		negroni.Wrap(deleteStudentObs(service)),
	)).Methods("DELETE").Name("deleteStudentObs")
}

func MakeStudentObsHandlers(r *mux.Router, service studentobs.UseCase) {
	r.Handle("/v1/studentobs", listStudentObs(service)).Methods("GET", "OPTIONS").Name("listStudentObs")

	r.Handle("/v1/studentobs", createStudentObs(service)).Methods("POST", "OPTIONS").Name("createStudentObs")

	r.Handle("/v1/studentobs/{id}", getStudentObs(service)).Methods("GET", "OPTIONS").Name("getStudentObs")

	r.Handle("/v1/studentobs/{id}", deleteStudentObs(service)).Methods("DELETE").Name("deleteStudentObs")
}
