package database

import (
	"database/sql"
	"time"
)

// DBConfig is a struct that represents database connection information
type DBConfig struct {
	User            string
	Password        string
	Host            string
	Port            int
	DBName          string
	ConnMaxLifetime time.Duration
	MaxOpenConns    int
	MaxIdleConns    int
}

// Connection holds the database connection and the auto generated queries by sqlc
type Connection struct {
	DB *sql.DB
}
