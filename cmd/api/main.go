package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/romeulima/devbook/internal/config"
	"github.com/romeulima/devbook/internal/server"
)

func main() {
	config.LoadEnvironments()

	router := server.CreateRouter()

	fmt.Printf("Server is running on port %v\n", config.ApiPort)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", config.ApiPort), router))
}
