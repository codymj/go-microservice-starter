package main

import (
	"fmt"
	"go-microservice-starter/cmd/app/config"
	"go-microservice-starter/cmd/app/routes"
	"log"
	"net/http"
)

func main() {
	err := start()
	if err != nil {
		panic(fmt.Errorf("fatal error starting service: %w \n", err))
	}
}

func start() error {
	log.Println("initializing application...")

	// set up service
	config.Set()
	port := config.Registry.GetString("SERVER_PORT")

	// set up routes
	router := routes.NewRouter()
	router.Setup()

	// start application
	log.Printf("service running on port %s\n", port)
	err := http.ListenAndServe(port, router.Router)
	if err != nil {
		log.Println("error setting up application")
		return err
	}

	return nil
}
