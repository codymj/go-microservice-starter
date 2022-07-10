package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

// NewConnection is a struct that wraps a Querier and the actual sql.DB connection
func NewConnection(cfg *Config) (*Connection, error) {
	connStr := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d",
		cfg.Host,
		cfg.User,
		cfg.Password,
		cfg.DBName,
		cfg.Port,
	)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	// set defaults
	db.SetConnMaxLifetime(cfg.ConnMaxLifetime)
	db.SetMaxOpenConns(cfg.MaxOpenConns)
	db.SetMaxIdleConns(cfg.MaxIdleConns)

	conn := Connection{
		DB: db,
	}

	return &conn, nil
}
