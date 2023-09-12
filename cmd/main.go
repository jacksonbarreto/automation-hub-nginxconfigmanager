package main

import (
	"automation-hub-nginxconfigmanager/internal/app/config"
	"automation-hub-nginxconfigmanager/internal/app/router"
	"log"
	"net/http"
)

func main() {
	config.Init()

	mux := router.SetupRoutes()

	server := &http.Server{
		Addr:    config.Config.ServerPort,
		Handler: mux,
	}

	log.Println("Starting server on :8080")
	log.Fatal(server.ListenAndServe())
}
