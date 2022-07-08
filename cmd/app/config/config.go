package config

import (
	"fmt"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"go-microservice-starter/internal/database"
	"go-microservice-starter/internal/users"
	"go-microservice-starter/internal/users/users_repository"
	"go-microservice-starter/internal/validate"
	"time"

	"github.com/spf13/viper"
)

// Registry for configuration values
var (
	Registry *viper.Viper
)

// Set configuration parameters
func Set() {
	viper.AutomaticEnv()

	Registry = viper.GetViper()
	Registry.AddConfigPath(".")
	Registry.AddConfigPath("../..")
	Registry.SetConfigFile("settings.yaml")
	err := Registry.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w \n", err))
	}

	SetLoggerParams()
}

// SetLoggerParams sets requested log level and other parameters
func SetLoggerParams() {
	zerolog.TimeFieldFormat = time.RFC3339

	loggerLevel := Registry.GetString("LOGGER_LEVEL")
	switch loggerLevel {
	case "trace":
		zerolog.SetGlobalLevel(zerolog.TraceLevel)
	case "debug":
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	case "info":
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	case "warn":
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
	case "error":
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	case "fatal":
		zerolog.SetGlobalLevel(zerolog.FatalLevel)
	case "panic":
		zerolog.SetGlobalLevel(zerolog.PanicLevel)
	default:
		log.Warn().Msg("invalid log level in config, defaulting to info")
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}
}

// GetDBSettings returns the DBConfig (database configuration parameters)
func GetDBSettings() database.DBConfig {
	return database.DBConfig{
		User:     Registry.GetString("DB_USER"),
		Password: Registry.GetString("DB_PASSWORD"),
		DBName:   Registry.GetString("DB_DBNAME"),
		Host:     Registry.GetString("DB_HOST"),
		Port:     Registry.GetInt("DB_PORT"),
	}
}

// NewUserRepository creates an instance of the users_repository
func NewUserRepository(db *database.Connection) users_repository.Repository {
	return users_repository.New(db)
}

// NewValidateService creates an instance of the json validate service
func NewValidateService() validate.Service {
	return validate.New()
}

// NewUserService creates an instance of the user_service
func NewUserService(ur users_repository.Repository) users.Service {
	return users.New(ur)
}
