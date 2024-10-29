package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	middleware "github.com/romeulima/devbook/internal/middlewares"
)

type Route struct {
	URI      string
	Method   string
	Function http.HandlerFunc
	NeedAuth bool
}

func InsertRoutes(r *mux.Router) *mux.Router {
	routes := usersRoutes
	routes = append(routes, loginRoute)

	for _, route := range routes {
		if route.NeedAuth {
			r.HandleFunc(route.URI, middleware.VerifyRequest(route.Function))
		} else {
			r.HandleFunc(route.URI, route.Function).Methods(route.Method)
		}
	}

	return r
}
