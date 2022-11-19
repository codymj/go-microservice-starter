package database

import (
	"database/sql"
	"time"
)

// Config is a struct that contains database information
type Config struct {
	User            string
	Password        string
	Host            string
	Port            int
	DBName          string
	ConnMaxLifetime time.Duration
	MaxOpenConns    int
	MaxIdleConns    int
}

// Connection for database
type Connection struct {
	DB *sql.DB
}
