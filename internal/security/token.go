package security

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/romeulima/devbook/internal/config"
)

var (
	t *jwt.Token
)

func GenerateToken(userId int) (string, error) {
	t = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss": "devbook-api",
		"sub": userId,
		"exp": time.Now().Add(time.Hour * 6).Unix(),
	})
	config.SigningKey = []byte("3fgrgrg")
	return t.SignedString(config.SigningKey)
}

func ValidadeToken(r *http.Request) error {
	tokenStr := extrairToken(r)

	token, err := jwt.Parse(tokenStr, returnVerificationKey)

	if err != nil {
		fmt.Println("oi")
		return err
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}

	return errors.New("token invalido")
}

func extrairToken(r *http.Request) string {
	header := r.Header.Get("Authorization")

	if len(strings.Split(header, " ")) != 2 {
		return ""
	}

	return strings.Split(header, " ")[1]
}

func returnVerificationKey(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("metodo de assinatura inesperado %v", token.Header["alg"])
	}

	return config.SecretJwt, nil
}
