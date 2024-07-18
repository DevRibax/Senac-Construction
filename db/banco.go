package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func Acessar() *sql.DB {
	var conexao = "user=postgres dbname=postgres password=123 host=localhost sslmode=disable"

	db, err := sql.Open("postgres", conexao)

	if err != nil {
		panic(err.Error())
	}
	return db
}
