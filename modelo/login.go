package modelo

import (
	"database/sql"
	"log"
	"modulo/db"
)

type Usuario struct {
	Id    int
	Email string
	Senha int
}

func Insert(email string, senha int) {
	db := db.Acessar()
	defer db.Close()

	insereDados, err := db.Prepare("INSERT INTO login(email,senha) VALUES ($1, $2)")
	if err != nil {
		panic(err.Error())
	}

	_, err = insereDados.Exec(email, senha)
	if err != nil {
		panic(err.Error())
	}
}

func AuthenticateUser(email string, senha int) bool {
	db := db.Acessar()
	defer db.Close()

	var user Usuario
	err := db.QueryRow("SELECT id, email, senha FROM login WHERE email=$1 AND senha=$2", email, senha).Scan(&user.Id, &user.Email, &user.Senha)
	if err != nil {
		if err == sql.ErrNoRows {
			return false
		}
		log.Fatal(err)
	}
	return true
}
