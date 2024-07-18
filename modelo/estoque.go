package modelo

import "modulo/db"

type Estoque struct {
	Id             int
	Produto        string
	Codigodebarras int
	Quantidade     int
	Condicao       string
	Observacao     string
}

func TrazerTodosOsEstoques() []Estoque {
	db := db.Acessar()

	selectDeTodosOsEstoques, err := db.Query("select * from estoque")
	if err != nil {
		panic(err.Error())
	}

	e := Estoque{}
	estoques := []Estoque{}

	for selectDeTodosOsEstoques.Next() {
		var id int
		var produto string
		var codigodebarras int
		var quantidade int
		var condicao string
		var observacao string
		err = selectDeTodosOsEstoques.Scan(&id, &produto, &codigodebarras, &quantidade, &condicao, &observacao)
		if err != nil {
			panic(err.Error())
		}
		e.Id = id
		e.Produto = produto
		e.Codigodebarras = codigodebarras
		e.Quantidade = quantidade
		e.Condicao = condicao
		e.Observacao = observacao
		estoques = append(estoques, e)
	}

	defer db.Close()
	return estoques
}

func InserirEstoque(produto string, codigodebarras int, quantidade int, condicao string, observacao string) {
	db := db.Acessar()

	insert, err := db.Prepare("insert into estoque (produto, codigodebarras, quantidade, condição, observação) values ($1, $2, $3, $4, $5)")
	if err != nil {
		panic(err.Error())
	}
	insert.Exec(produto, codigodebarras, quantidade, condicao, observacao)
	defer db.Close()
}

func DeletarEstoque(id string) {
	db := db.Acessar()

	deleteestoque, err := db.Prepare("delete from estoque where id = $1")
	if err != nil {
		panic(err.Error())
	}
	deleteestoque.Exec(id)
	defer db.Close()
}

func EditarEstoque(id string) Estoque {
	db := db.Acessar()

	atualizar, err := db.Query("select * from estoque where id = $1", id)
	if err != nil {
		panic(err.Error())
	}

	EstoqueAtualizar := Estoque{}

	for atualizar.Next() {
		var id int
		var produto string
		var codigodebarras int
		var quantidade int
		var condicao string
		var observacao string
		err = atualizar.Scan(&id, &produto, &codigodebarras, &quantidade, &condicao, &observacao)
		if err != nil {
			panic(err.Error())
		}
		EstoqueAtualizar.Id = id
		EstoqueAtualizar.Produto = produto
		EstoqueAtualizar.Codigodebarras = codigodebarras
		EstoqueAtualizar.Quantidade = quantidade
		EstoqueAtualizar.Condicao = condicao
		EstoqueAtualizar.Observacao = observacao
	}

	defer db.Close()
	return EstoqueAtualizar
}

func Atualizaestoque(id int, produto string, codigodebarras int, quantidade int, condicao string, observacao string) {
	db := db.Acessar()

	atualizaProduto, err := db.Prepare("update estoque set produto=$1, codigodebarras=$2, quantidade=$3, condição=$4, observação=$5 where id=$6")
	if err != nil {
		panic(err.Error())
	}
	atualizaProduto.Exec(produto, codigodebarras, quantidade, condicao, observacao, id)
	defer db.Close()
}
