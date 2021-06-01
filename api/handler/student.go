package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"iitd_control_escolar.api/api/presenter"
	"iitd_control_escolar.api/entity"
	"iitd_control_escolar.api/usecase/student"
	"log"
	"net/http"
	"time"
)

func listStudents(service student.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error reading students"
		var data []*entity.Student
		var err error
		name := r.URL.Query().Get("name")
		switch {
		case name == "":
			data, err = service.ListStudents()
		default:
			data, err = service.SearchStudents(name)
		}
		w.Header().Set("Content-Type", "application/json")
		if err != nil && err != entity.ErrNotFound {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(errorMessage))
			return
		}

		if data == nil {
			w.WriteHeader(http.StatusNotFound)
			_, _ = w.Write([]byte(errorMessage))
			return
		}
		var toJ []*presenter.Student
		for _, d := range data {
			toJ = append(toJ, &presenter.Student{
				ID:            d.ID,
				Nombres:       d.Nombres,
				Apellidos:     d.Apellidos,
				Nacimiento:    d.Nacimiento,
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
				FechaInicio:   d.FechaInicio,
				Observaciones: d.Observaciones,
				Activo:        d.Activo,
			})
		}
		if err := json.NewEncoder(w).Encode(toJ); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(errorMessage))
		}
	})
}

func createStudent(service student.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error adding student"
		var input struct {
			ID            entity.ID `json:"id"`
			Nombres       string    `json:"nombres"`
			Apellidos     string    `json:"apellidos"`
			Nacimiento    time.Time `json:"nacimiento"`
			Sexo          string    `json:"sexo"`
			Calle         string    `json:"calle"`
			NumeroExt     string    `json:"numeroExt"`
			NumeroInt     string    `json:"numeroInt"`
			Colonia       string    `json:"colonia"`
			Municipio     string    `json:"municipio"`
			Estado        string    `json:"estado"`
			Pais          string    `json:"pais"`
			CP            string    `json:"cp"`
			TelCelular    string    `json:"telCelular"`
			TelCasa       string    `json:"telCasa"`
			Email         string    `json:"email"`
			FechaInicio   time.Time `json:"fechaInicio"`
			Observaciones string    `json:"observaciones"`
			Activo        string    `json:"activo"`
		}
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(errorMessage))
			return
		}
		var id entity.ID
		var s *entity.Student
		if input.ID == 0 {
			id, err = service.CreateStudent(input.Nombres, input.Apellidos, input.Nacimiento, input.Sexo, input.Calle,
				input.NumeroExt, input.NumeroInt, input.Colonia, input.Municipio, input.Estado, input.Pais, input.CP,
				input.TelCelular, input.TelCasa, input.Email, input.FechaInicio, input.Observaciones, input.Activo)
		} else {
			id = input.ID
			s, err = entity.NewStudent(input.Nombres, input.Apellidos, input.Nacimiento, input.Sexo, input.Calle,
				input.NumeroExt, input.NumeroInt, input.Colonia, input.Municipio, input.Estado, input.Pais, input.CP,
				input.TelCelular, input.TelCasa, input.Email, input.FechaInicio, input.Observaciones, input.Activo)
			if err != nil {
				log.Println(err.Error())
				w.WriteHeader(http.StatusInternalServerError)
				_, _ = w.Write([]byte(errorMessage))
				return
			}
			if s == nil {
				log.Println("")
				w.WriteHeader(http.StatusInternalServerError)
				_, _ = w.Write([]byte("error updating student"))
				return
			}
			s.ID = input.ID
			err = service.UpdateStudent(s)
		}
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(errorMessage))
			return
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

		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(toJ); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(errorMessage))
			return
		}
	})
}

func getStudent(service student.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error reading student"
		vars := mux.Vars(r)
		id, err := entity.StringToID(vars["id"])
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(errorMessage))
			return
		}
		data, err := service.GetStudent(id)
		w.Header().Set("Content-Type", "application/json")
		if err != nil && err != entity.ErrNotFound {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(errorMessage))
			return
		}

		if data == nil {
			w.WriteHeader(http.StatusNotFound)
			_, _ = w.Write([]byte(errorMessage))
			return
		}
		toJ := &presenter.Student{
			ID:            data.ID,
			Nombres:       data.Nombres,
			Apellidos:     data.Apellidos,
			Nacimiento:    data.Nacimiento,
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
			FechaInicio:   data.FechaInicio,
			Observaciones: data.Observaciones,
			Activo:        data.Activo,
		}
		if err := json.NewEncoder(w).Encode(toJ); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(errorMessage))
		}
	})
}

func deleteStudent(service student.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error removing student"
		vars := mux.Vars(r)
		id, err := entity.StringToID(vars["id"])
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(errorMessage))
			return
		}
		err = service.DeleteStudent(id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(errorMessage))
			return
		}

		ret := struct {
			Status bool `json:"status"`
		}{
			Status: true,
		}
		if err := json.NewEncoder(w).Encode(ret); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(errorMessage))
		}
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
