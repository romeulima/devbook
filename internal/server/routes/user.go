package routes

import (
	"net/http"

	"github.com/romeulima/devbook/internal/controller"
)

var usersRoutes = []Route{

	{
		URI:      "/users",
		Method:   http.MethodPost,
		Function: controller.CreateUser,
		NeedAuth: false,
	},
	{
		URI:      "/users",
		Method:   http.MethodGet,
		Function: controller.GetAllUsers,
		NeedAuth: false,
	},
	{
		URI:      "/users/{id}",
		Method:   http.MethodGet,
		Function: controller.GetUserById,
		NeedAuth: false,
	},
	{
		URI:      "/users/{id}",
		Method:   http.MethodPut,
		Function: controller.UpdateUser,
		NeedAuth: false,
	},
	{
		URI:      "/users/{id}",
		Method:   http.MethodDelete,
		Function: controller.DeleteUser,
		NeedAuth: false,
	},
}
