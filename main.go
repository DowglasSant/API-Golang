package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.Load()

	fmt.Printf("API Running on port %d...", config.Port)

	router := router.RouterGenerate()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), router))
}
