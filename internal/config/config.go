package config

import (
	"os"
	"strconv"
)

var (
	DbHost     = ""
	DbUser     = ""
	DbPassword = ""
	DbName     = ""
	ApiPort    = 0
)

func LoadEnvironments() {
	var err error

	DbHost = os.Getenv("DB_HOST")
	DbUser = os.Getenv("DB_USER")
	DbPassword = os.Getenv("DB_PASSWORD")
	DbName = os.Getenv("DB_NAME")
	ApiPort, err = strconv.Atoi(os.Getenv("API_PORT"))

	if err != nil {
		ApiPort = 9000
	}

}
