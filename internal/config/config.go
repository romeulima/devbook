package config

import (
	"crypto/ecdsa"
	"os"
)

var (
	DbHost     = ""
	DbUser     = ""
	DbPassword = ""
	DbName     = ""
	SigningKey []byte
	SecretJwt  *ecdsa.PrivateKey
)

func LoadEnvironments() {
	DbHost = os.Getenv("DB_HOST")
	DbUser = os.Getenv("DB_USER")
	DbPassword = os.Getenv("DB_PASSWORD")
	DbName = os.Getenv("DB_NAME")
}
