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
		RequiresAuthentication: false,
	},
	{
		URI:                    "/users/{userId}",
		Method:                 http.MethodGet,
		Function:               controllers.ShowUser,
		RequiresAuthentication: false,
	},
	{
		URI:                    "/users/{userId}",
		Method:                 http.MethodPut,
		Function:               controllers.UpdateUser,
		RequiresAuthentication: false,
	},
	{
		URI:                    "/users/{userId}",
		Method:                 http.MethodDelete,
		Function:               controllers.DeleteUser,
		RequiresAuthentication: false,
	},
}
