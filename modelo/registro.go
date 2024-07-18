package modelo

import "modulo/db"

type Registro struct {
	Id                 int
	Codigo             int
	TipoDeMovimentacao string
	Horario            string
	Observacao         string
}

func InserirRegistro(codigo int, tipodemovimentacao string, horario string, observacao string) {
	db := db.Acessar()

	Inseriregistro, err := db.Prepare("insert into registro (codigo, tipodemovimentacao, horario, observacao) values ($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}

	Inseriregistro.Exec(codigo, tipodemovimentacao, horario, observacao)
	defer db.Close()

}

func ListaReg() []Registro {
	db := db.Acessar()

	selectregistro, err := db.Query("select * from registro")
	if err != nil {
		panic(err.Error())
	}

	r := Registro{}
	registro := []Registro{}

	for selectregistro.Next() {
		var id int
		var codigo int
		var tipodemovimentacao string
		var horario string
		var observacao string

		err = selectregistro.Scan(&id, &codigo, &tipodemovimentacao, &horario, &observacao)
		if err != nil {
			panic(err.Error())
		}

		r.Id = id
		r.Codigo = codigo
		r.TipoDeMovimentacao = tipodemovimentacao
		r.Horario = horario
		r.Observacao = observacao

		registro = append(registro, r)
	}
	defer db.Close()
	return registro
}

func DeletaRegistro(id string) {
	db := db.Acessar()

	deletarfornecedor, err := db.Prepare("delete from registro where id=$1")
	if err != nil {
		panic(err.Error())
	}

	deletarfornecedor.Exec(id)
	defer db.Close()

}

func Editareg(id string) Registro {
	db := db.Acessar()

	atualizar, err := db.Query("select * from registro where id = $1", id)
	if err != nil {
		panic(err.Error())
	}

	atualizareg := Registro{}

	for atualizar.Next() {
		var id int
		var codigo int
		var tipodemovimentacao string
		var horario string
		var observacao string
		err = atualizar.Scan(&id, &codigo, &tipodemovimentacao, &horario, &observacao)
		if err != nil {
			panic(err.Error())
		}
		atualizareg.Id = id
		atualizareg.Codigo = codigo
		atualizareg.TipoDeMovimentacao = tipodemovimentacao
		atualizareg.Horario = horario
		atualizareg.Observacao = observacao
	}

	defer db.Close()
	return atualizareg
}

func Atualizareg(id int, codigo int, tipodemovimentacao string, horario string, observacao string) {
	db := db.Acessar()

	AtualizaProduto, err := db.Prepare("update registro set codigo=$1, tipodemovimentacao=$2, horario=$3, observacao=$4 where id=$5")
	if err != nil {
		panic(err.Error())
	}
	AtualizaProduto.Exec(codigo, tipodemovimentacao, horario, observacao, id)
	defer db.Close()
}
