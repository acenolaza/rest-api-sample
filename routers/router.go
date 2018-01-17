package routers

import (
	"github.com/acenolaza/rest-api-sample/handlers"
	"github.com/acenolaza/rest-api-sample/services"
	"github.com/auth0/go-jwt-middleware"
	"github.com/codegangsta/negroni"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"net/http"
)

type route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
	IsProtected bool
}

func (route *route) GetHandler(mw *jwtmiddleware.JWTMiddleware) http.Handler {
	var handler http.Handler
	handler = route.HandlerFunc
	if route.IsProtected {
		handler = negroni.New(
			negroni.HandlerFunc(mw.HandlerWithNext),
			negroni.Wrap(handler))
	}
	return handler
}

type routes []route

var authenticationRoutes = routes{
	route{
		"TokenAuth",
		"POST",
		"/token-auth",
		handlers.CreateToken,
		false,
	},
	route{
		"RefreshTokenAuth",
		"GET",
		"/refresh-token-auth",
		handlers.RefreshToken,
		true,
	},
	route{
		"RemoveTokenAuth",
		"GET",
		"/remove-token-auth",
		handlers.RemoveToken,
		true,
	}}

var todoRoutes = routes{
	route{
		"CreateTodo",
		"POST",
		"/todos",
		handlers.CreateTodo,
		false,
	},
	route{
		"FetchAllTodo",
		"GET",
		"/todos",
		handlers.FetchAllTodo,
		false,
	},
	route{
		"FetchSingleTodo",
		"GET",
		"/todos/{id}",
		handlers.FetchSingleTodo,
		false,
	},
	route{
		"DeleteTodo",
		"DELETE",
		"/todos/{id}",
		handlers.DeleteTodo,
		false,
	},
}

func InitRouter() *mux.Router {
	mw := jwtmiddleware.New(jwtmiddleware.Options{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return services.GetPublicKey(), nil
		},
		SigningMethod: jwt.SigningMethodRS512,
	})

	// new router
	router := mux.NewRouter().StrictSlash(true)

	// set status ping handler
	router.HandleFunc("/", handlers.Index)

	// add api version to path
	var v1Router = router.PathPrefix("/v1").Subrouter()

	// set all routes handlers
	for _, route := range append(authenticationRoutes, todoRoutes...) {
		v1Router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.GetHandler(mw))
	}

	return router
}
