package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"time"
)

// NewConnection is a struct that wraps a Querier and the actual sql.DB connection
func NewConnection(cfg DBConfig) (*Connection, error) {
	connectionString := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host,
		cfg.Port,
		cfg.User,
		cfg.Password,
		cfg.DBName,
	)
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	//Set defaults
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	conn := Connection{
		Qs: Querier(New(db)),
		DB: db,
	}
	return &conn, nil
}
