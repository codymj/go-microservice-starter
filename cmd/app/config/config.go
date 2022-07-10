package config

import (
	"fmt"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"go-microservice-starter/internal/database"
	"go-microservice-starter/internal/password"
	"go-microservice-starter/internal/users"
	"go-microservice-starter/internal/users/users_dao"
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

// GetDatabaseConfig returns the DBConfig (database configuration parameters)
func GetDatabaseConfig() *database.Config {
	connMaxLifetimeProp := Registry.GetString("DB_CONN_MAX_LIFETIME")
	connMaxLifetime, err := time.ParseDuration(connMaxLifetimeProp)
	if err != nil {
		log.Warn().Msg("setting database connection max lifetime to 3 minutes")
		connMaxLifetime, _ = time.ParseDuration("3m")
	}

	return &database.Config{
		User:            Registry.GetString("DB_USER"),
		Password:        Registry.GetString("DB_PASSWORD"),
		DBName:          Registry.GetString("DB_DBNAME"),
		Host:            Registry.GetString("DB_HOST"),
		Port:            Registry.GetInt("DB_PORT"),
		ConnMaxLifetime: connMaxLifetime,
		MaxOpenConns:    Registry.GetInt("DB_MAX_OPEN_CONNS"),
		MaxIdleConns:    Registry.GetInt("DB_MAX_IDLE_CONNS"),
	}
}

// GetPasswordConfig returns the config struct for configuring password hashes
func GetPasswordConfig() *password.Config {
	return &password.Config{
		Time:    1,
		Memory:  64 * 1024,
		Threads: 4,
		KeyLen:  32,
	}
}

// NewUserRepository creates an instance of the users_dao
func NewUserRepository(db *database.Connection, ps password.Service) users_dao.Repository {
	return users_dao.New(db, ps)
}

// NewValidateService creates an instance of the json validate service
func NewValidateService() validate.Service {
	return validate.New()
}

// NewPasswordService creates an instance of the password hash service
func NewPasswordService(cfg *password.Config) password.Service {
	return password.New(cfg)
}

// NewUserService creates an instance of the user_service
func NewUserService(ur users_dao.Repository) users.Service {
	return users.New(ur)
}
