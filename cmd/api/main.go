package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/romeulima/devbook/internal/config"
	"github.com/romeulima/devbook/internal/server"
)

func main() {
	if os.Getenv("LOAD_ENV_FILE") == "true" {
		err := godotenv.Load()
		if err != nil {
			panic("Error loading .env file")
		}
	}
	config.LoadEnvironments()

	router := server.CreateRouter()

	fmt.Printf("Server is running on port %s\n", ":8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
