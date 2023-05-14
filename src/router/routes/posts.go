package routes

import (
	"api/src/controllers"
	"net/http"
)

var postsRoutes = []Route{
	{
		URI:                    "/posts",
		Method:                 http.MethodPost,
		Function:               controllers.CreatePost,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/posts",
		Method:                 http.MethodGet,
		Function:               controllers.ShowPosts,
		RequiresAuthentication: true,
	},

	{
		URI:                    "/posts/{postId}",
		Method:                 http.MethodGet,
		Function:               controllers.ShowPost,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/posts/{postId}",
		Method:                 http.MethodPut,
		Function:               controllers.UpdatePost,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/posts/{postId}",
		Method:                 http.MethodDelete,
		Function:               controllers.DeletePost,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/users/{userId}/posts",
		Method:                 http.MethodGet,
		Function:               controllers.FindPostsByUser,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/posts/{postId}/like",
		Method:                 http.MethodPost,
		Function:               controllers.LikePost,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/posts/{postId}/unlike",
		Method:                 http.MethodPost,
		Function:               controllers.UnlikePost,
		RequiresAuthentication: true,
	},
}
