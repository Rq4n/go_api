package database

import (
	"database/sql"
)

var DB *sql.DB

func Connect() error {
	connStr := "host:localhost port=5432 user=api password=api dbname=projetoAPI sslmode=disable"
	var err error
	DB, err := sql.Open("postgres", connStr)
	return err
}
