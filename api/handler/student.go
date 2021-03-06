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
	"iitd_control_escolar.api/usecase/student"
)

func listStudents(service student.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var data []*entity.Student
		var err error
		name := r.URL.Query().Get("name")
		switch {
		case name == "":
			data, err = service.ListStudents()
		default:
			data, err = service.SearchStudents(name)
		}
		if err != nil { //&& err != entity.ErrNotFound {
			sendErrorResponse(w, http.StatusInternalServerError, err_db_reading, err)
			return
		}

		if data == nil {
			sendOkResponse(w, http.StatusOK, []*presenter.Student{})
			return
		}
		var toJ []*presenter.Student
		for _, d := range data {
			toJ = append(toJ, &presenter.Student{
				ID:            d.ID,
				Nombres:       d.Nombres,
				Apellidos:     d.Apellidos,
				Nacimiento:    jd.JsonDate(d.Nacimiento),
				Sexo:          d.Sexo,
				Calle:         d.Calle,
				NumeroExt:     d.NumeroExt,
				NumeroInt:     d.NumeroInt,
				Colonia:       d.Colonia,
				Municipio:     d.Municipio,
				Estado:        d.Estado,
				Pais:          d.Pais,
				CP:            d.CP,
				TelCelular:    d.TelCelular,
				TelCasa:       d.TelCasa,
				Email:         d.Email,
				FechaInicio:   jd.JsonDate(d.FechaInicio),
				Observaciones: d.Observaciones,
				Activo:        d.Activo,
			})
		}
		sendOkResponse(w, http.StatusOK, toJ)
	})
}

func createStudent(service student.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var input struct {
			ID            int         `json:"id"`
			Nombres       string      `json:"nombres"`
			Apellidos     string      `json:"apellidos"`
			Nacimiento    jd.JsonDate `json:"nacimiento"`
			Sexo          string      `json:"sexo"`
			Calle         string      `json:"calle"`
			NumeroExt     string      `json:"numeroExt"`
			NumeroInt     string      `json:"numeroInt"`
			Colonia       string      `json:"colonia"`
			Municipio     string      `json:"municipio"`
			Estado        string      `json:"estado"`
			Pais          string      `json:"pais"`
			CP            string      `json:"cp"`
			TelCelular    string      `json:"telCelular"`
			TelCasa       string      `json:"telCasa"`
			Email         string      `json:"email"`
			FechaInicio   jd.JsonDate `json:"fechaInicio"`
			Observaciones string      `json:"observaciones"`
			Activo        string      `json:"activo"`
		}
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			sendErrorResponse(w, http.StatusInternalServerError, err_decoding_body, err)
			return
		}
		var id int
		var s *entity.Student
		if input.ID == 0 {
			id, err = service.CreateStudent(input.Nombres, input.Apellidos, input.Nacimiento.ToTime(), input.Sexo, input.Calle,
				input.NumeroExt, input.NumeroInt, input.Colonia, input.Municipio, input.Estado, input.Pais, input.CP,
				input.TelCelular, input.TelCasa, input.Email, input.FechaInicio.ToTime(), input.Observaciones, input.Activo)
			if err != nil {
				sendErrorResponse(w, http.StatusInternalServerError, err_db_inserting, err)
				return
			}
		} else {
			id = input.ID
			s, err = entity.NewStudent(input.Nombres, input.Apellidos, input.Nacimiento.ToTime(), input.Sexo, input.Calle,
				input.NumeroExt, input.NumeroInt, input.Colonia, input.Municipio, input.Estado, input.Pais, input.CP,
				input.TelCelular, input.TelCasa, input.Email, input.FechaInicio.ToTime(), input.Observaciones, input.Activo)
			if err != nil {
				sendErrorResponse(w, http.StatusInternalServerError, err_building_object, err)
				return
			}
			if s == nil {
				sendErrorResponse(w, http.StatusInternalServerError, err_unexpected_nil_object, nil)
				return
			}
			s.ID = input.ID
			err = service.UpdateStudent(s)
			if err != nil {
				sendErrorResponse(w, http.StatusInternalServerError, err_db_updating, err)
				return
			}
		}
		toJ := &presenter.Student{
			ID:            id,
			Nombres:       input.Nombres,
			Apellidos:     input.Apellidos,
			Nacimiento:    input.Nacimiento,
			Sexo:          input.Sexo,
			Calle:         input.Calle,
			NumeroExt:     input.NumeroExt,
			NumeroInt:     input.NumeroInt,
			Colonia:       input.Colonia,
			Municipio:     input.Municipio,
			Estado:        input.Estado,
			Pais:          input.Pais,
			CP:            input.CP,
			TelCelular:    input.TelCelular,
			TelCasa:       input.TelCasa,
			Email:         input.Email,
			FechaInicio:   input.FechaInicio,
			Observaciones: input.Observaciones,
			Activo:        input.Activo,
		}

		sendOkResponse(w, http.StatusCreated, toJ)
	})
}

func getStudent(service student.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			sendErrorResponse(w, http.StatusBadRequest, err_slug_value, err)
			return
		}
		data, err := service.GetStudent(id)
		if err != nil { //&& err != entity.ErrNotFound {
			sendErrorResponse(w, http.StatusInternalServerError, err_db_reading, err)
			return
		}

		if data == nil {
			sendOkResponse(w, http.StatusNotFound, nil)
			return
		}
		toJ := &presenter.Student{
			ID:            data.ID,
			Nombres:       data.Nombres,
			Apellidos:     data.Apellidos,
			Nacimiento:    jd.JsonDate(data.Nacimiento),
			Sexo:          data.Sexo,
			Calle:         data.Calle,
			NumeroExt:     data.NumeroExt,
			NumeroInt:     data.NumeroInt,
			Colonia:       data.Colonia,
			Municipio:     data.Municipio,
			Estado:        data.Estado,
			Pais:          data.Pais,
			CP:            data.CP,
			TelCelular:    data.TelCelular,
			TelCasa:       data.TelCasa,
			Email:         data.Email,
			FechaInicio:   jd.JsonDate(data.FechaInicio),
			Observaciones: data.Observaciones,
			Activo:        data.Activo,
		}
		sendOkResponse(w, http.StatusOK, toJ)
	})
}

func deleteStudent(service student.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			sendErrorResponse(w, http.StatusBadRequest, err_slug_value, err)
			return
		}
		err = service.DeleteStudent(id)
		if err != nil {
			sendErrorResponse(w, http.StatusInternalServerError, "error reading record from database", err)
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

//MakeStudentHandlers make url handlers
func MakeStudentHandlersNegroni(r *mux.Router, n negroni.Negroni, service student.UseCase) {
	r.Handle("/v1/students", n.With(
		negroni.Wrap(listStudents(service)),
	)).Methods("GET", "OPTIONS").Name("listStudents")

	r.Handle("/v1/students", n.With(
		negroni.Wrap(createStudent(service)),
	)).Methods("POST", "OPTIONS").Name("createStudent")

	r.Handle("/v1/students/{id}", n.With(
		negroni.Wrap(getStudent(service)),
	)).Methods("GET", "OPTIONS").Name("getStudent")

	r.Handle("/v1/students/{id}", n.With(
		negroni.Wrap(deleteStudent(service)),
	)).Methods("DELETE").Name("deleteStudent")
}

func MakeStudentHandlers(r *mux.Router, service student.UseCase) {
	r.Handle("/v1/students", listStudents(service)).Methods("GET", "OPTIONS").Name("listStudents")

	r.Handle("/v1/students", createStudent(service)).Methods("POST", "OPTIONS").Name("createStudent")

	r.Handle("/v1/students/{id}", getStudent(service)).Methods("GET", "OPTIONS").Name("getStudent")

	r.Handle("/v1/students/{id}", deleteStudent(service)).Methods("DELETE").Name("deleteStudent")
}
