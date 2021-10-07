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
	"iitd_control_escolar.api/usecase/maestro"
)

func listMaestros(service maestro.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var data []*entity.Maestro
		var err error
		name := r.URL.Query().Get("name")
		switch {
		case name == "":
			data, err = service.ListMaestros()
		default:
			data, err = service.SearchMaestros(name)
		}
		if err != nil { //&& err != entity.ErrNotFound {
			sendErrorResponse(w, http.StatusInternalServerError, err_db_reading, err)
			return
		}

		var toJ []*presenter.Maestro = []*presenter.Maestro{}
		for _, d := range data {
			toJ = append(toJ, &presenter.Maestro{
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

func createMaestro(service maestro.UseCase) http.Handler {
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
			log.Println(err.Error())
			sendErrorResponse(w, http.StatusInternalServerError, err_decoding_body, err)
			return
		}
		var id int
		var s *entity.Maestro
		if input.ID == 0 {
			id, err = service.CreateMaestro(input.Nombres, input.Apellidos, input.Nacimiento.ToTime(), input.Sexo, input.Calle,
				input.NumeroExt, input.NumeroInt, input.Colonia, input.Municipio, input.Estado, input.Pais, input.CP,
				input.TelCelular, input.TelCasa, input.Email, input.FechaInicio.ToTime(), input.Observaciones, input.Activo)
			if err != nil {
				sendErrorResponse(w, http.StatusInternalServerError, err_db_inserting, err)
				return
			}
		} else {
			id = input.ID
			s, err = entity.NewMaestro(input.Nombres, input.Apellidos, input.Nacimiento.ToTime(), input.Sexo, input.Calle,
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
			err = service.UpdateMaestro(s)
			if err != nil {
				sendErrorResponse(w, http.StatusInternalServerError, err_db_updating, err)
				return
			}
		}
		toJ := &presenter.Maestro{
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

func getMaestro(service maestro.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			sendErrorResponse(w, http.StatusBadRequest, err_slug_value, err)
			return
		}
		data, err := service.GetMaestro(id)
		if err != nil { //&& err != entity.ErrNotFound {
			sendErrorResponse(w, http.StatusInternalServerError, err_db_reading, err)
			return
		}

		if data == nil {
			sendOkResponse(w, http.StatusNotFound, nil)
			return
		}
		toJ := &presenter.Maestro{
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

func deleteMaestro(service maestro.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			sendErrorResponse(w, http.StatusBadRequest, err_slug_value, err)
			return
		}
		err = service.DeleteMaestro(id)
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

//MakeMaestroHandlers make url handlers
func MakeMaestroHandlersNegroni(r *mux.Router, n negroni.Negroni, service maestro.UseCase) {
	r.Handle("/v1/maestros", n.With(
		negroni.Wrap(listMaestros(service)),
	)).Methods("GET", "OPTIONS").Name("listMaestros")

	r.Handle("/v1/maestros", n.With(
		negroni.Wrap(createMaestro(service)),
	)).Methods("POST", "OPTIONS").Name("createMaestro")

	r.Handle("/v1/maestros/{id}", n.With(
		negroni.Wrap(getMaestro(service)),
	)).Methods("GET", "OPTIONS").Name("getMaestro")

	r.Handle("/v1/maestros/{id}", n.With(
		negroni.Wrap(deleteMaestro(service)),
	)).Methods("DELETE").Name("deleteMaestro")
}

func MakeMaestroHandlers(r *mux.Router, service maestro.UseCase) {
	r.Handle("/v1/maestros", listMaestros(service)).Methods("GET", "OPTIONS").Name("listMaestros")

	r.Handle("/v1/maestros", createMaestro(service)).Methods("POST", "OPTIONS").Name("createMaestro")

	r.Handle("/v1/maestros/{id}", getMaestro(service)).Methods("GET", "OPTIONS").Name("getMaestro")

	r.Handle("/v1/maestros/{id}", deleteMaestro(service)).Methods("DELETE").Name("deleteMaestro")
}
