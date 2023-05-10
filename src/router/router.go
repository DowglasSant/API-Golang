package router

import (
	"api/src/router/routes"

	"github.com/gorilla/mux"
)

func RouterGenerate() *mux.Router {
	router := mux.NewRouter()
	return routes.Config(router)
}
