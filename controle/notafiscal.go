package controle

import (
	"html/template"
	"log"
	"modulo/modelo"
	"net/http"
	"strconv"
)

var templl = template.Must(template.ParseGlob("templates/*.html"))

func MostraNotasFiscais(w http.ResponseWriter, r *http.Request) {
	pegartodasasnotasfiscais := modelo.TrazerTodasNotasFiscai()
	templl.ExecuteTemplate(w, "MostraNotasFiscais", pegartodasasnotasfiscais)
}

func NotaFiscal(w http.ResponseWriter, r *http.Request) {
	templl.ExecuteTemplate(w, "NotaFiscal", nil)
}

func ColocarNotasFiscais(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		notafiscaldofornecedor := r.FormValue("notafiscaldofornecedor")
		nomedoproduto := r.FormValue("nomedoproduto")
		codigo := r.FormValue("codigo")
		tipodemovimentacao := r.FormValue("tipodemovimentacao")
		horario := r.FormValue("horario")
		quantidade := r.FormValue("quantidade")
		observacao := r.FormValue("observacao")

		convertercodigo, err := strconv.Atoi(codigo)
		if err != nil {
			log.Println("Error na conversão do código", err)
		}
		converterquantidade, err := strconv.Atoi(quantidade)
		if err != nil {
			log.Println("Error na conversão quantidade", err)
		}

		modelo.InserirNotasfiscais(notafiscaldofornecedor, nomedoproduto, convertercodigo, tipodemovimentacao, horario, converterquantidade, observacao)
	}

	http.Redirect(w, r, "/listar_notafiscal", 301)
}

func Deletarnotafiscal(w http.ResponseWriter, r *http.Request) {
	idDoNotafistal := r.URL.Query().Get("id")
	modelo.DeletarNota(idDoNotafistal)
	http.Redirect(w, r, "listar_notafiscal", 301)
}

func Editarnotafiscal(w http.ResponseWriter, r *http.Request) {
	idAtualizar := r.URL.Query().Get("id")
	notafiscal := modelo.EditarNota(idAtualizar)
	templl.ExecuteTemplate(w, "Editarnotafiscal", notafiscal)
}

func Nota(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		notafiscaldofornecedor := r.FormValue("notafiscaldofornecedor")
		nomedoproduto := r.FormValue("nomedoproduto")
		codigo := r.FormValue("codigo")
		tipodemovimentacao := r.FormValue("tipodemovimentacao")
		horario := r.FormValue("horario")
		quantidade := r.FormValue("quantidade")
		observacao := r.FormValue("observacao")

		idConvertidaParaInt, err := strconv.Atoi(id)
		if err != nil {
			log.Println("Erro na convesão do ID para int:", err)
		}

		convertercodigo, err := strconv.Atoi(codigo)
		if err != nil {
			log.Println("Error na conversão do código", err)
		}
		converterquantidade, err := strconv.Atoi(quantidade)
		if err != nil {
			log.Println("Error na conversão quantidade", err)
		}

		modelo.Atualizanota(idConvertidaParaInt, notafiscaldofornecedor, nomedoproduto, convertercodigo, tipodemovimentacao, horario, converterquantidade, observacao)
	}
	http.Redirect(w, r, "/listar_notafiscal", 301)
}
