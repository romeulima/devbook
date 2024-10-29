package controller

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/romeulima/devbook/internal/database"
	"github.com/romeulima/devbook/internal/models"
	"github.com/romeulima/devbook/internal/repository"
	"github.com/romeulima/devbook/pkg/jsonr"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	userRequest := new(models.UserRequest)

	defer r.Body.Close()

	if err := json.NewDecoder(r.Body).Decode(userRequest); err != nil {
		jsonr.WriteJSON(w, http.StatusBadRequest, models.Error{
			Message: "Invalid request Payload",
		})
		return
	}

	if err := userRequest.Prepare("cadastro", userRequest); err != nil {
		jsonr.WriteJSON(w, http.StatusBadRequest, models.Error{
			Message: err.Error(),
		})
		return
	}

	userEntity := models.NewUser(userRequest)

	dbpool, err := database.Connect()

	if err != nil {
		jsonr.WriteJSON(w, http.StatusInternalServerError, models.Error{Message: err.Error()})
		return
	}

	userRepository := repository.NewRepository(dbpool)

	if err = userRepository.CreateUser(r.Context(), userEntity); err != nil {
		status := http.StatusInternalServerError
		errorMessage := err.Error()

		if strings.Contains(err.Error(), "23505") {
			status = http.StatusBadRequest
			errorMessage = "Email or nick is already registered"
		}

		jsonr.WriteJSON(w, status, models.Error{Message: errorMessage})
		return
	}

	jsonr.WriteJSON(w, http.StatusCreated, userEntity)

}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	userParam := strings.ToLower(r.URL.Query().Get("user"))

	dbpool, err := database.Connect()
	if err != nil {
		jsonr.WriteJSON(w, http.StatusInternalServerError, models.Error{Message: err.Error()})
		return
	}

	userRepository := repository.NewRepository(dbpool)

	users, err := userRepository.GetUsers(r.Context(), userParam)

	if err != nil {
		jsonr.WriteJSON(w, http.StatusInternalServerError, models.Error{Message: err.Error()})
		return
	}

	jsonr.WriteJSON(w, http.StatusOK, users)
}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userId, err := strconv.Atoi(params["id"])

	if err != nil {
		jsonr.WriteJSON(w, http.StatusBadRequest, models.Error{Message: err.Error()})
		return
	}

	dbpool, err := database.Connect()
	if err != nil {
		jsonr.WriteJSON(w, http.StatusInternalServerError, models.Error{Message: err.Error()})
		return
	}

	userRepository := repository.NewRepository(dbpool)

	user, err := userRepository.GetUserById(r.Context(), userId)

	if err != nil {
		jsonr.WriteJSON(w, http.StatusNotFound, models.Error{Message: "user not found"})
		return
	}

	jsonr.WriteJSON(w, http.StatusOK, user)

}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userId, err := strconv.Atoi(params["id"])

	if err != nil {
		jsonr.WriteJSON(w, http.StatusBadRequest, models.Error{Message: err.Error()})
		return
	}

	userRequest := new(models.UserRequest)

	defer r.Body.Close()

	if err = json.NewDecoder(r.Body).Decode(userRequest); err != nil {
		jsonr.WriteJSON(w, http.StatusBadRequest, models.Error{
			Message: "Invalid request Payload",
		})
		return
	}

	if err = userRequest.ValidadeFields("update", userRequest); err != nil {
		jsonr.WriteJSON(w, http.StatusBadRequest, models.Error{
			Message: err.Error(),
		})
		return
	}

	user := models.NewUser(userRequest)

	dbpool, err := database.Connect()
	if err != nil {
		jsonr.WriteJSON(w, http.StatusInternalServerError, models.Error{Message: err.Error()})
		return
	}

	userRepository := repository.NewRepository(dbpool)

	_, err = userRepository.GetUserById(r.Context(), userId)

	if err != nil {
		jsonr.WriteJSON(w, http.StatusNotFound, models.Error{Message: "user not found"})
		return
	}

	if err = userRepository.UpdateUser(r.Context(), userId, user); err != nil {
		jsonr.WriteJSON(w, http.StatusInternalServerError, models.Error{Message: err.Error()})
		return
	}

	jsonr.WriteJSON(w, http.StatusNoContent, nil)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userId, err := strconv.Atoi(params["id"])

	if err != nil {
		jsonr.WriteJSON(w, http.StatusBadRequest, models.Error{Message: err.Error()})
		return
	}

	dbpool, err := database.Connect()

	if err != nil {
		jsonr.WriteJSON(w, http.StatusInternalServerError, models.Error{Message: err.Error()})
		return
	}

	userRepository := repository.NewRepository(dbpool)

	_, err = userRepository.GetUserById(r.Context(), userId)

	if err != nil {
		jsonr.WriteJSON(w, http.StatusNotFound, models.Error{Message: "user not found"})
		return
	}

	if err = userRepository.DeleteUser(r.Context(), userId); err != nil {
		jsonr.WriteJSON(w, http.StatusInternalServerError, models.Error{Message: err.Error()})
		return
	}

	jsonr.WriteJSON(w, http.StatusNoContent, nil)
}
