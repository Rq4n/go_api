package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Connect() {
	connStr := "host=localhost port=5432 user=api password=api dbname=projetoAPI sslmode=disable"

	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Erro ao abrir conexão com DB:", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("Erro ao conectar ao DB:", err)
	}

	log.Println("Conexão com o Postgres estabelecida com sucesso!")
}
