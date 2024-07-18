package modelo

import (
	"modulo/db"
)

type Mater struct {
	Id        int
	Categoria string
	Tipo      string
	Descricao string
}

func TrazerTodosOsMateriais() []Mater {
	db := db.Acessar()

	selectDeTodosOsMateriais, err := db.Query("select * from mater")
	if err != nil {
		panic(err.Error())
	}

	material := Mater{}
	materiais := []Mater{}

	for selectDeTodosOsMateriais.Next() {
		var id int
		var categoria, tipo, descricao string
		err = selectDeTodosOsMateriais.Scan(&id, &categoria, &tipo, &descricao)
		if err != nil {
			panic(err.Error())
		}
		material.Id = id
		material.Categoria = categoria
		material.Tipo = tipo
		material.Descricao = descricao
		materiais = append(materiais, material)
	}

	defer db.Close()
	return materiais
}

func InserirMaterial(categoria, tipo, descricao string) {
	db := db.Acessar()

	insert, err := db.Prepare("insert into mater (categoria, tipo, descricao) values ($1, $2, $3)")
	if err != nil {
		panic(err.Error())
	}
	insert.Exec(categoria, tipo, descricao)
	defer db.Close()
}

func DeletarMaterial(id string) {
	db := db.Acessar()

	delete, err := db.Prepare("delete from mater where id = $1")
	if err != nil {
		panic(err.Error())
	}
	delete.Exec(id)
	defer db.Close()
}

func EditarMaterial(id string) Mater {
	db := db.Acessar()

	atualizar, err := db.Query("select * from mater where id = $1", id)
	if err != nil {
		panic(err.Error())
	}

	MaterialAtualizar := Mater{}

	for atualizar.Next() {
		var id int
		var categoria, tipo, descricao string
		err = atualizar.Scan(&id, &categoria, &tipo, &descricao)
		if err != nil {
			panic(err.Error())
		}
		MaterialAtualizar.Id = id
		MaterialAtualizar.Categoria = categoria
		MaterialAtualizar.Tipo = tipo
		MaterialAtualizar.Descricao = descricao
	}

	defer db.Close()
	return MaterialAtualizar
}

func Atualizamater(id int, categoria, tipo, descricao string) {
	db := db.Acessar()

	atualizamater, err := db.Prepare("update mater set categoria=$1, tipo=$2, descricao=$3 where id=$4")
	if err != nil {
		panic(err.Error())
	}
	atualizamater.Exec(categoria, tipo, descricao, id)
	defer db.Close()
}
