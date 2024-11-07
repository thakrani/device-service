package db

import (
	"database/sql"
	"device-service/configs"
	"fmt"

	_ "github.com/lib/pq"
)

type DB struct {
	conn *sql.DB
}

type IDB interface {
	Open(connectionString string) (*sql.DB, error)
	Ping() error
	QueryRow(query string, args ...interface{}) *sql.Row
	Query(query string, args ...interface{}) (*sql.Rows, error)
	Exec(query string, args ...interface{}) (sql.Result, error)
	Close() error
}

func (d *DB) Open(connectionString string) (*sql.DB, error) {
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}
	d.conn = db
	return db, nil
}

func (d *DB) Ping() error {
	return d.conn.Ping()
}

func (d *DB) QueryRow(query string, args ...interface{}) *sql.Row {
	return d.conn.QueryRow(query, args...)
}

func (d *DB) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return d.conn.Query(query, args...)
}

func (d *DB) Exec(query string, args ...interface{}) (sql.Result, error) {
	return d.conn.Exec(query, args...)
}

func (d *DB) Close() error {
	return d.conn.Close()
}

func NewDB(config configs.Config) (DB, error) {
	connStr := fmt.Sprintf("postgres://%s:%s@localhost/%s?sslmode=disable", config.Username, config.Password, config.DatabaseName)
	dbInstance := DB{}
	db, err := dbInstance.Open(connStr)
	if err != nil {
		return DB{}, err
	}

	if err = db.Ping(); err != nil {
		return DB{}, err
	}

	return dbInstance, nil
}
