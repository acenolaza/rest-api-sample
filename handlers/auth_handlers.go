package handlers

import (
	"github.com/acenolaza/rest-api-sample/api"
	"github.com/acenolaza/rest-api-sample/services"
	"github.com/acenolaza/rest-api-sample/services/models"
	"net/http"
)

// { "uuid":"", "username":"testuser", "password":"passw0rd"
func CreateToken(w http.ResponseWriter, r *http.Request) {
	req := api.Request{Request: r}
	res := api.Response{ResponseWriter: w}

	requestUser := new(models.User)
	if err := req.GetJSONBody(requestUser); err != nil {
		res.RespondWithError(http.StatusBadRequest, err.Error())
		return
	}

	responseStatus, token := services.CreateToken(requestUser)

	res.RespondWithJSON(responseStatus, token)
}

func RefreshToken(w http.ResponseWriter, r *http.Request) {
	res := api.Response{ResponseWriter: w}

	res.RespondWithError(http.StatusNotImplemented, "Endpoint is not yet implemented")
}

func RemoveToken(w http.ResponseWriter, r *http.Request) {
	res := api.Response{ResponseWriter: w}

	res.RespondWithError(http.StatusNotImplemented, "Endpoint is not yet implemented")
}
