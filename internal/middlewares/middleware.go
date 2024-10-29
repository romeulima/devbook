package middleware

import (
	"net/http"

	"github.com/romeulima/devbook/internal/models"
	"github.com/romeulima/devbook/internal/security"
	"github.com/romeulima/devbook/pkg/jsonr"
)

func VerifyRequest(function http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := security.ValidadeToken(r); err != nil {
			jsonr.WriteJSON(w, http.StatusUnauthorized, models.Error{Message: err.Error()})
			return
		}
		function(w, r)
	}
}
