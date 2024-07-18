package modelo

import (
	"modulo/db"
)

type Nota_fiscal struct {
	Id                     int
	Notafiscaldofornecedor string
	Nomedoproduto          string
	Codigo                 int
	Tipodemovimentacao     string
	Horario                string
	Quantidade             int
	Observacao             string
}

func TrazerTodasNotasFiscai() []Nota_fiscal {
	db := db.Acessar()

	selectDeTodasAsNotasFiscais, err := db.Query("select * from notafiscal")
	if err != nil {
		panic(err.Error())
	}

	e := Nota_fiscal{}
	notafiscal := []Nota_fiscal{}

	for selectDeTodasAsNotasFiscais.Next() {
		var id int
		var notafiscaldofornecedor string
		var nomedoproduto string
		var codigo int
		var tipodemovimentacao string
		var horario string
		var quantidade int
		var observacao string

		err = selectDeTodasAsNotasFiscais.Scan(&id, &notafiscaldofornecedor, &nomedoproduto, &codigo, &tipodemovimentacao, &horario, &quantidade, &observacao)
		if err != nil {
			panic(err.Error())
		}
		e.Id = id
		e.Notafiscaldofornecedor = notafiscaldofornecedor
		e.Nomedoproduto = nomedoproduto
		e.Codigo = codigo
		e.Tipodemovimentacao = tipodemovimentacao
		e.Horario = horario
		e.Quantidade = quantidade
		e.Observacao = observacao
		notafiscal = append(notafiscal, e)
	}

	defer db.Close()
	return notafiscal
}

func InserirNotasfiscais(notafiscaldofornecedor string, nomedoproduto string, codigo int, tipodemovimentacao string, horario string, quantidade int, observacao string) {
	db := db.Acessar()

	insert, err := db.Prepare("insert into notafiscal (notafiscaldofornecedor, nomedoproduto, codigo, tipodemovimentacao, horario, quantidade, observacao) values ($1, $2, $3, $4, $5, $6, $7)")
	if err != nil {
		panic(err.Error())
	}
	insert.Exec(notafiscaldofornecedor, nomedoproduto, codigo, tipodemovimentacao, horario, quantidade, observacao)
	defer db.Close()
}

func DeletarNota(id string) {
	db := db.Acessar()

	deletarNotafiscal, err := db.Prepare("delete from notafiscal where id = $1")
	if err != nil {
		panic(err.Error())
	}
	deletarNotafiscal.Exec(id)
	defer db.Close()
}

func EditarNota(id string) Nota_fiscal {
	db := db.Acessar()

	atualizar, err := db.Query("select * from notafiscal where id = $1", id)
	if err != nil {
		panic(err.Error())
	}

	NotafiscalAtualizar := Nota_fiscal{}

	for atualizar.Next() {
		var id int
		var notafiscaldofornecedor string
		var nomedoproduto string
		var codigo int
		var tipodemovimentacao string
		var horario string
		var quantidade int
		var observacao string
		err = atualizar.Scan(&id, &notafiscaldofornecedor, &nomedoproduto, &codigo, &tipodemovimentacao, &horario, &quantidade, &observacao)
		if err != nil {
			panic(err.Error())
		}
		NotafiscalAtualizar.Id = id
		NotafiscalAtualizar.Notafiscaldofornecedor = notafiscaldofornecedor
		NotafiscalAtualizar.Nomedoproduto = nomedoproduto
		NotafiscalAtualizar.Codigo = codigo
		NotafiscalAtualizar.Tipodemovimentacao = tipodemovimentacao
		NotafiscalAtualizar.Horario = horario
		NotafiscalAtualizar.Quantidade = quantidade
		NotafiscalAtualizar.Observacao = observacao

	}

	defer db.Close()
	return NotafiscalAtualizar
}

func Atualizanota(id int, notafiscaldofornecedor string, nomedoproduto string, codigo int, tipodemovimentacao string, horario string, quantidade int, observacao string) {
	db := db.Acessar()

	atualizaProduto, err := db.Prepare("update notafiscal set notafiscaldofornecedor=$1, nomedoproduto=$2, codigo=$3, tipodemovimentacao=$4, horario=$5, quantidade=$6, observacao=$7 where id=$8")
	if err != nil {
		panic(err.Error())
	}
	atualizaProduto.Exec(notafiscaldofornecedor, nomedoproduto, codigo, tipodemovimentacao, horario, quantidade, observacao, id)
	defer db.Close()
}
