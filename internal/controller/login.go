package controller

import (
	"encoding/json"
	"net/http"

	"github.com/badoux/checkmail"
	"github.com/romeulima/devbook/internal/database"
	"github.com/romeulima/devbook/internal/models"
	"github.com/romeulima/devbook/internal/repository"
	"github.com/romeulima/devbook/internal/security"
	"github.com/romeulima/devbook/pkg/jsonr"
)

func Login(w http.ResponseWriter, r *http.Request) {
	userRequest := new(models.UserRequest)

	defer r.Body.Close()

	if err := json.NewDecoder(r.Body).Decode(userRequest); err != nil {
		jsonr.WriteJSON(w, http.StatusBadRequest, models.Error{
			Message: "Invalid request Payload",
		})
		return
	}

	if userRequest.Email == "" || userRequest.Password == "" {
		jsonr.WriteJSON(w, http.StatusBadRequest, models.Error{
			Message: "Missing required fields",
		})
		return
	}

	if err := checkmail.ValidateFormat(userRequest.Email); err != nil {
		jsonr.WriteJSON(w, http.StatusBadRequest, models.Error{
			Message: "this email is invalid",
		})
		return
	}

	dbpool, err := database.Connect()

	if err != nil {
		jsonr.WriteJSON(w, http.StatusInternalServerError, models.Error{
			Message: err.Error(),
		})
		return
	}

	userRepository := repository.NewRepository(dbpool)
	userEntity, err := userRepository.GetUserByEmail(r.Context(), userRequest.Email)

	if err != nil {
		jsonr.WriteJSON(w, http.StatusNotFound, models.Error{
			Message: "user not found",
		})
		return
	}

	err = security.ComparePasswords(userEntity.Password, userRequest.Password)

	if err != nil {
		jsonr.WriteJSON(w, http.StatusUnauthorized, models.Error{
			Message: err.Error(),
		})
		return
	}

	token, err := security.GenerateToken(userEntity.ID)

	if err != nil {
		jsonr.WriteJSON(w, http.StatusInternalServerError, models.Error{Message: err.Error()})
		return
	}

	jsonr.WriteJSON(w, http.StatusCreated, token)

}
