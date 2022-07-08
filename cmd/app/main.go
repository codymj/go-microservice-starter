package main

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"go-microservice-starter/cmd/app/config"
	"go-microservice-starter/cmd/app/routes"
	"go-microservice-starter/internal/database"
	"net/http"
	"time"
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
	readTOProp := config.Registry.GetString("SERVER_READ_TIMEOUT")
	readTimeout, err := time.ParseDuration(readTOProp)
	if err != nil {
		log.Fatal().Msg("please set a valid server read timeout duration")
		return err
	}
	writeTOProp := config.Registry.GetString("SERVER_WRITE_TIMEOUT")
	writeTimeout, err := time.ParseDuration(writeTOProp)
	if err != nil {
		log.Fatal().Msg("please set a valid server write timeout duration")
		return err
	}
	idleTOProp := config.Registry.GetString("SERVER_IDLE_TIMEOUT")
	idleTimeout, err := time.ParseDuration(idleTOProp)
	if err != nil {
		log.Fatal().Msg("please set a valid server idle timeout duration")
		return err
	}

	// init db
	dbSettings := config.GetDBSettings()
	db, err := database.NewConnection(dbSettings)
	if err != nil {
		log.Fatal().Msg("error setting up database connection")
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
		log.Fatal().Msg("error setting up routes")
		return err
	}

	// start application
	log.Info().Msg(fmt.Sprintf("service running on port %s", port))
	srv := &http.Server{
		Handler:      router.Router,
		Addr:         ":" + port,
		WriteTimeout: writeTimeout,
		ReadTimeout:  readTimeout,
		IdleTimeout:  idleTimeout,
	}
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal().Msg("error setting up application")
		return err
	}

	return nil
}
