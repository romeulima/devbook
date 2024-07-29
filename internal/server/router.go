package server

import (
	"github.com/gorilla/mux"
	"github.com/romeulima/devbook/internal/server/routes"
)

func CreateRouter() *mux.Router {
	r := mux.NewRouter()
	return routes.InsertRoutes(r)
}
