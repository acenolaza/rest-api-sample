package api

import (
	"encoding/json"
	"github.com/acenolaza/rest-api-sample/api/parameters"
	"net/http"
)

type Response struct {
	ResponseWriter http.ResponseWriter
}

func (r *Response) RespondWithError(code int, message string) {
	r.RespondWithJSON(code, parameters.JsonError{Code: code, Text: message})
}

func (r *Response) RespondWithJSON(code int, payload interface{}) {
	r.ResponseWriter.Header().Set("Content-Type", "application/json; charset=UTF-8")
	r.ResponseWriter.WriteHeader(code)
	if err := json.NewEncoder(r.ResponseWriter).Encode(payload); err != nil {
		r.RespondWithError(http.StatusInternalServerError, err.Error())
	}
}
