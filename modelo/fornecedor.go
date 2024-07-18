package modelo

import "modulo/db"

type Fornecedor struct {
	Id                 int
	Razaosocial        string
	Nomefantasia       string
	Cnpj               int
	Endereco           string
	Bairro             string
	Cidade             string
	Estado             string
	Cep                int
	Telefone           int
	Inscricaoestadual  int
	Inscricaomunicipal int
	Atividadeeconomica string
	Email              string
}

func InserirFornecedor(razaosocial string, nomefantasia string, cnpj int, endereco string, bairro string, cidade string, estado string, cep int, telefone int,
	inscricaoestadual int, inscricaomunicipal int, atividadeeconomica string, email string) {
	db := db.Acessar()

	Inserirfornecedor, err := db.Prepare("insert into dadosfornecedor (razaosocial, nomefantasia, cnpj, endereco, bairro, cidade, estado, cep, telefone, inscricaoestadual, inscricaomunicipal, atividadeeconomica, email) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)")
	if err != nil {
		panic(err.Error())
	}

	Inserirfornecedor.Exec(razaosocial, nomefantasia, cnpj, endereco, bairro, cidade, estado, cep, telefone, inscricaoestadual, inscricaomunicipal, atividadeeconomica, email)
	defer db.Close()

}

func ListaForn() []Fornecedor {
	db := db.Acessar()

	selectfornecedor, err := db.Query("select * from dadosfornecedor")
	if err != nil {
		panic(err.Error())
	}

	f := Fornecedor{}
	produtos := []Fornecedor{}

	for selectfornecedor.Next() {
		var id int
		var razaosocial string
		var nomefantasia string
		var cnpj int
		var endereco string
		var bairro string
		var cidade string
		var estado string
		var cep int
		var telefone int
		var inscricaoestadual int
		var inscricaomunicipal int
		var atividadeeconomica string
		var email string

		err = selectfornecedor.Scan(&id, &razaosocial, &nomefantasia, &cnpj, &endereco, &bairro, &cidade, &estado, &cep, &telefone, &inscricaoestadual, &inscricaomunicipal, &atividadeeconomica, &email)
		if err != nil {
			panic(err.Error())
		}

		f.Id = id
		f.Razaosocial = razaosocial
		f.Nomefantasia = nomefantasia
		f.Cnpj = cnpj
		f.Endereco = endereco
		f.Bairro = bairro
		f.Cidade = cidade
		f.Estado = estado
		f.Cep = cep
		f.Telefone = telefone
		f.Inscricaoestadual = inscricaoestadual
		f.Inscricaomunicipal = inscricaomunicipal
		f.Atividadeeconomica = atividadeeconomica
		f.Email = email

		produtos = append(produtos, f)
	}
	defer db.Close()
	return produtos
}

func Deletafornecedor(id string) {
	db := db.Acessar()

	deletarfornecedor, err := db.Prepare("delete from dadosfornecedor where id=$1")
	if err != nil {
		panic(err.Error())
	}

	deletarfornecedor.Exec(id)
	defer db.Close()

}

func Alterarforn(id string) Fornecedor {
	db := db.Acessar()

	alterar, err := db.Query("select * from dadosfornecedor where id = $1", id)
	if err != nil {
		panic(err.Error())
	}

	Fornatualizar := Fornecedor{}

	for alterar.Next() {
		var id int
		var razaosocial string
		var nomefantasia string
		var cnpj int
		var endereco string
		var bairro string
		var cidade string
		var estado string
		var cep int
		var telefone int
		var inscricaoestadual int
		var inscricaomunicipal int
		var atividadeeconomica string
		var email string

		err = alterar.Scan(&id, &razaosocial, &nomefantasia, &cnpj, &endereco, &bairro, &cidade, &estado, &cep, &telefone, &inscricaoestadual, &inscricaomunicipal, &atividadeeconomica, &email)
		if err != nil {
			panic(err.Error())
		}

		Fornatualizar.Id = id
		Fornatualizar.Razaosocial = razaosocial
		Fornatualizar.Nomefantasia = nomefantasia
		Fornatualizar.Cnpj = cnpj
		Fornatualizar.Endereco = endereco
		Fornatualizar.Bairro = bairro
		Fornatualizar.Cidade = cidade
		Fornatualizar.Estado = estado
		Fornatualizar.Cep = cep
		Fornatualizar.Telefone = telefone
		Fornatualizar.Inscricaoestadual = inscricaoestadual
		Fornatualizar.Inscricaomunicipal = inscricaomunicipal
		Fornatualizar.Atividadeeconomica = atividadeeconomica
		Fornatualizar.Email = email

	}

	defer db.Close()
	return Fornatualizar
}

func Atualizaforn(id int, razaosocial string, nomefantasia string, cnpj int, endereco string, bairro string, cidade string, estado string, cep int, telefone int,
	inscricaoestadual int, inscricaomunicipal int, atividadeeconomica string, email string) {
	db := db.Acessar()

	atualizaforn, err := db.Prepare("update dadosfornecedor set razaosocial=$1, nomefantasia=$2, cnpj=$3, endereco=$4, bairro=$5, cidade=$6, estado=$7, cep=$8, telefone=$9, inscricaoestadual=$10, inscricaomunicipal=$11, atividadeeconomica=$12, email=$13 where id=$14")
	if err != nil {
		panic(err.Error())
	}
	atualizaforn.Exec(razaosocial, nomefantasia, cnpj, endereco, bairro, cidade, estado, cep, telefone, inscricaoestadual, inscricaomunicipal, atividadeeconomica, email, id)
	defer db.Close()
}
