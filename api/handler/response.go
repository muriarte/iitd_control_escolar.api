package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

type Response struct {
	Data       interface{} `json:"data,omitempty"`
	HttpStatus int         `json:"httpStatus"`
	ErrorInfo  ErrDetail   `json:"errorInfo"`
}

type ErrDetail struct {
	ErrDesc string `json:"errDesc"`
}

// newResponse construye una estructura de tipo Response
func newResponse(data interface{}, httpStatus int, errDesc string) Response {
	return Response{
		Data:       data,
		HttpStatus: httpStatus,
		ErrorInfo:  ErrDetail{ErrDesc: errDesc},
	}
}

// sendOkResponse envia una respuesta exitosa al http.ResponseWriter proporcionado
func sendOkResponse(w http.ResponseWriter, httpStatus int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(httpStatus)
	resp := newResponse(data, httpStatus, "")
	if err1 := json.NewEncoder(w).Encode(resp); err1 != nil {
		log.Printf("Problemas intentar enviar respuesta exitosa: " + err1.Error())
	}
}

// sendOkResponse envia una respuesta con error al http.ResponseWriter proporcionado
func sendErrorResponse(w http.ResponseWriter, httpStatus int, errorMsg string, err error) {
	if err != nil {
		if len(strings.TrimSpace(errorMsg)) > 0 {
			errorMsg = errorMsg + ": " + err.Error()
		}
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatus)
	resp := newResponse(nil, http.StatusInternalServerError, errorMsg)
	if err1 := json.NewEncoder(w).Encode(resp); err1 != nil {
		log.Printf("Problemas intentar enviar respuesta exitosa: " + err1.Error())
	}
}
