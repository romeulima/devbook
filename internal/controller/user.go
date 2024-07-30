package controller

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/romeulima/devbook/internal/database"
	"github.com/romeulima/devbook/internal/models"
	"github.com/romeulima/devbook/internal/repository"
	"github.com/romeulima/devbook/pkg"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	userRequest := new(models.UserRequest)

	defer r.Body.Close()

	if err := json.NewDecoder(r.Body).Decode(userRequest); err != nil {
		pkg.WriteJSON(w, http.StatusBadRequest, models.Error{
			Message: "Invalid request Payload",
		})
		return
	}

	if err := userRequest.ValidadeFields(userRequest); err != nil {
		pkg.WriteJSON(w, http.StatusBadRequest, models.Error{
			Message: err.Error(),
		})
		return
	}

	userEntity := models.NewUser(userRequest)

	dbpool, err := database.Connect()

	if err != nil {
		pkg.WriteJSON(w, http.StatusInternalServerError, models.Error{Message: err.Error()})
		return
	}

	defer dbpool.Close()

	userRepository := repository.NewRepository(dbpool)

	if err = userRepository.CreateUser(r.Context(), userEntity); err != nil {
		status := http.StatusInternalServerError
		errorMessage := err.Error()

		if strings.Contains(err.Error(), "23505") {
			status = http.StatusBadRequest
			errorMessage = "Email or nick is already registered"
		}

		pkg.WriteJSON(w, status, models.Error{Message: errorMessage})
		return
	}

	pkg.WriteJSON(w, http.StatusCreated, userEntity)

}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("message: Getting all users"))
}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("message: Getting user by id"))
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("message: Updating user informations"))
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("message: Deleting user"))
}
