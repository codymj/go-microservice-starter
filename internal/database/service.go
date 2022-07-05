package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"time"
)

// NewConnection is a struct that wraps a Querier and the actual sql.DB connection
func NewConnection(cfg DBConfig) (*Connection, error) {
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
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	conn := Connection{
		DB: db,
	}
	return &conn, nil
}
