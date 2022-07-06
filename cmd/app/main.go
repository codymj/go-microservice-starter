package main

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"go-microservice-starter/cmd/app/config"
	"go-microservice-starter/cmd/app/routes"
	"go-microservice-starter/internal/database"
	"net/http"
)

func main() {
	err := start()
	if err != nil {
		panic(fmt.Errorf("fatal error starting service: %w \n", err))
	}
}

func start() error {
	// init application
	config.Set()
	port := config.Registry.GetString("PORT")

	// init db
	dbSettings := config.GetDBSettings()
	db, err := database.NewConnection(dbSettings)
	if err != nil {
		return err
	}

	// init repositories
	ur := config.NewUserRepository(db)

	// init services
	vs := config.NewValidateService()
	us := config.NewUserService(ur)

	// init routes
	routeServices := routes.Services{
		ValidatorService: vs,
		UserService:      us,
	}
	router := routes.NewRouter()
	err = router.Setup(routeServices)
	if err != nil {
		log.Error().Msg("error setting up routes")
		return err
	}

	// start application
	log.Info().Msg(fmt.Sprintf("service running on port %s", port))
	err = http.ListenAndServe(":"+port, router.Router)
	if err != nil {
		log.Error().Msg("error setting up application")
		return err
	}

	return nil
}
