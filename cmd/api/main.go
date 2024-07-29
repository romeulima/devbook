package main

import (
	"log"
	"net/http"

	"github.com/romeulima/devbook/internal/config"
	"github.com/romeulima/devbook/internal/server"
)

func main() {
	config.LoadVariables()
	router := server.CreateRouter()

	log.Fatal(http.ListenAndServe(":5000", router))
}
