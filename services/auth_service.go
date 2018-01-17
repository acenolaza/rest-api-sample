package services

import (
	"crypto/rsa"
	"github.com/acenolaza/rest-api-sample/api/parameters"
	"github.com/acenolaza/rest-api-sample/core/authentication"
	"github.com/acenolaza/rest-api-sample/services/models"
	"net/http"
)

func CreateToken(requestUser *models.User) (int, *parameters.TokenAuthentication) {
	authBackend := authentication.GetJWTAuthenticationBackend()

	if authBackend.Authenticate(requestUser) {
		token, err := authBackend.GenerateToken(requestUser.UUID)
		if err != nil {
			return http.StatusInternalServerError, new(parameters.TokenAuthentication)
		} else {
			return http.StatusOK, &parameters.TokenAuthentication{Token: token}
		}
	}

	return http.StatusUnauthorized, new(parameters.TokenAuthentication)
}

func GetPublicKey() *rsa.PublicKey {
	authBackend := authentication.GetJWTAuthenticationBackend()

	return authBackend.PublicKey
}
