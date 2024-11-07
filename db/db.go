package db

import (
	"database/sql"
	"device-service/configs"
	"fmt"

	_ "github.com/lib/pq"
)

func NewDB(config configs.Config) (*sql.DB, error) {
	connStr := fmt.Sprintf("postgres://%s:%s@localhost/%s?sslmode=disable", config.Username, config.Password, config.DatabaseName)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
