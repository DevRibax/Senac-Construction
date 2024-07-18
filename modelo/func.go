package modelo

import (
	"modulo/db"
)

type Funcionarios struct {
	Id    int
	Nome  string
	Cargo string
	Cpf   string
	Rg    string
	Email string
	Senha int
}

func Insira(nome string, cargo string, cpf string, rg string, email string, senha int) {
	db := db.Acessar()

	inserefuncionario, err := db.Prepare("insert into funcionario (nome,cargo,cpf,rg,email,senha) values ($1, $2, $3, $4, $5, $6)")
	if err != nil {
		panic(err.Error())
	}

	inserefuncionario.Exec(nome, cargo, cpf, rg, email, senha)
	defer db.Close()
}

func Trazerfuncionario() []Funcionarios {
	db := db.Acessar()

	selectfuncionario, err := db.Query("select * from funcionario")
	if err != nil {
		panic(err.Error())
	}

	funci := Funcionarios{}
	funcionario := []Funcionarios{}

	for selectfuncionario.Next() {
		var id int
		var nome string
		var cargo string
		var cpf string
		var rg string
		var email string
		var senha int

		err = selectfuncionario.Scan(&id, &nome, &cargo, &cpf, &rg, &email, &senha)
		if err != nil {
			panic(err.Error())
		}

		funci.Id = id
		funci.Nome = nome
		funci.Cargo = cargo
		funci.Cpf = cpf
		funci.Rg = rg
		funci.Email = email
		funci.Senha = senha

		funcionario = append(funcionario, funci)
	}
	defer db.Close()
	return funcionario
}

func Deletafuncionario(id string) {
	db := db.Acessar()

	deletarfuncionario, err := db.Prepare("delete from funcionario where id=$1")
	if err != nil {
		panic(err.Error())
	}

	deletarfuncionario.Exec(id)
	defer db.Close()

}
