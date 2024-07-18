package controle

import (
	"html/template"
	"log"
	"modulo/modelo"
	"net/http"
	"strconv"
)

var templa = template.Must(template.ParseGlob("templates/*.html"))

func Fornecedor(w http.ResponseWriter, r *http.Request) {

	templa.ExecuteTemplate(w, "Fornecedor", nil)

}
func ListarFornecedor(w http.ResponseWriter, r *http.Request) {
	todosAsEmpresas := modelo.ListaForn()
	templa.ExecuteTemplate(w, "ListarFornecedor", todosAsEmpresas)

}

func Adicionar(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		razaosocial := r.FormValue("razaosocial")

		nomefantasia := r.FormValue("nomefantasia")

		cnpj := r.FormValue("cnpj")

		endereco := r.FormValue("endereco")

		bairro := r.FormValue("bairro")

		cidade := r.FormValue("cidade")

		estado := r.FormValue("estado")

		cep := r.FormValue("cep")

		telefone := r.FormValue("telefone")

		inscricaoestadual := r.FormValue("inscricaoestadual")

		inscricaomunicipal := r.FormValue("inscricaomunicipal")

		atividadeeconomica := r.FormValue("atividadeeconomica")

		email := r.FormValue("email")

		CnpjConvertidaParaInt, err := strconv.Atoi(cnpj)
		if err != nil {
			log.Println("Erro na conversão do cnpj:", err)
		}

		CepConvertidaParaInt, err := strconv.Atoi(cep)
		if err != nil {
			log.Println("Erro na conversão do cep:", err)
		}

		telefoneconverte, err := strconv.Atoi(telefone)
		if err != nil {
			log.Println("Erro na conversão do telefone:", err)
		}

		InscricaoEstadualConvertidaParaInt, err := strconv.Atoi(inscricaoestadual)
		if err != nil {
			log.Println("Erro na conversão do inscricaoEstadual:", err)
		}

		InscricaoMunicipalConvertidaParaInt, err := strconv.Atoi(inscricaomunicipal)
		if err != nil {
			log.Println("Erro na conversão do inscricaoMunicipal:", err)
		}

		modelo.InserirFornecedor(razaosocial, nomefantasia, CnpjConvertidaParaInt, endereco, bairro, cidade, estado, CepConvertidaParaInt, telefoneconverte, InscricaoEstadualConvertidaParaInt, InscricaoMunicipalConvertidaParaInt, atividadeeconomica, email)
	}
	http.Redirect(w, r, "/listar_fornecedor", 301)
}

func Deletefornecedor(w http.ResponseWriter, r *http.Request) {
	idDEmpresa := r.URL.Query().Get("id")
	modelo.Deletafornecedor(idDEmpresa)
	http.Redirect(w, r, "/listar_fornecedor", 301)
}

func Editarfornecedor(w http.ResponseWriter, r *http.Request) {
	idAtualizar := r.URL.Query().Get("id")
	fornecedor := modelo.Alterarforn(idAtualizar)
	templa.ExecuteTemplate(w, "Editarfornecedor", fornecedor)
}

func AlteraFornecedor(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		razaosocial := r.FormValue("razaosocial")
		nomefantasia := r.FormValue("nomefantasia")
		cnpj := r.FormValue("cnpj")
		endereco := r.FormValue("endereco")
		bairro := r.FormValue("bairro")
		cidade := r.FormValue("cidade")
		estado := r.FormValue("estado")
		cep := r.FormValue("cep")
		telefone := r.FormValue("telefone")
		inscricaoestadual := r.FormValue("inscricaoestadual")
		inscricaomunicipal := r.FormValue("inscricaomunicipal")
		atividadeeconomica := r.FormValue("atividadeeconomica")
		email := r.FormValue("email")

		idConvertidaParaInt, err := strconv.Atoi(id)
		if err != nil {
			log.Println("Erro na convesão do ID para int:", err)
		}

		CnpjConvertidaParaInt, err := strconv.Atoi(cnpj)
		if err != nil {
			log.Println("Erro na conversão do cnpj:", err)
		}

		CepConvertidaParaInt, err := strconv.Atoi(cep)
		if err != nil {
			log.Println("Erro na conversão do cep:", err)
		}

		telefoneconverte, err := strconv.Atoi(telefone)
		if err != nil {
			log.Println("Erro na conversão do telefone:", err)
		}

		InscricaoEstadualConvertidaParaInt, err := strconv.Atoi(inscricaoestadual)
		if err != nil {
			log.Println("Erro na conversão do inscricaoEstadual:", err)
		}

		InscricaoMunicipalConvertidaParaInt, err := strconv.Atoi(inscricaomunicipal)
		if err != nil {
			log.Println("Erro na conversão do inscricaoMunicipal:", err)
		}
		modelo.Atualizaforn(idConvertidaParaInt, razaosocial, nomefantasia, CnpjConvertidaParaInt, endereco, bairro, cidade, estado, CepConvertidaParaInt, telefoneconverte, InscricaoEstadualConvertidaParaInt, InscricaoMunicipalConvertidaParaInt, atividadeeconomica, email)
	}
	http.Redirect(w, r, "/listar_fornecedor", 301)
}
