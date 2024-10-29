package routes

import (
	"net/http"

	"github.com/romeulima/devbook/internal/controller"
)

var loginRoute = Route{
	URI:      "/login",
	Method:   http.MethodPost,
	Function: controller.Login,
	NeedAuth: false,
}
