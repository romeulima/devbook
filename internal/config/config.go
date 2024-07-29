package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	connectionString = ""
	port             = 0
)

func LoadVariables() {
	var err error
	if err = godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	port, err = strconv.Atoi(os.Getenv("API_PORT"))

	if err != nil {
		port = 9000
	}

}
