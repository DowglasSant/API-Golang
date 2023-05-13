package routes

import (
	"api/src/controllers"
	"net/http"
)

var usersRoutes = []Route{
	{
		URI:                    "/users",
		Method:                 http.MethodPost,
		Function:               controllers.AddUser,
		RequiresAuthentication: false,
	},
	{
		URI:                    "/users",
		Method:                 http.MethodGet,
		Function:               controllers.ShowAllUsers,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/users/{userId}",
		Method:                 http.MethodGet,
		Function:               controllers.ShowUser,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/users/{userId}",
		Method:                 http.MethodPut,
		Function:               controllers.UpdateUser,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/users/{userId}",
		Method:                 http.MethodDelete,
		Function:               controllers.DeleteUser,
		RequiresAuthentication: true,
	},
}
