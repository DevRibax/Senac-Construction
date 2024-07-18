package controle

import (
	"html/template"
	"log"
	"modulo/modelo"
	"net/http"
	"strconv"
)

var templ = template.Must(template.ParseGlob("templates/*.html"))

func Insp(w http.ResponseWriter, r *http.Request) {

	templ.ExecuteTemplate(w, "Insp", nil)

}

func Mostraestoque(w http.ResponseWriter, r *http.Request) {
	pegartodososestoques := modelo.TrazerTodosOsEstoques()
	templ.ExecuteTemplate(w, "Mostraestoque", pegartodososestoques)
}

func Colocar(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		produto := r.FormValue("produto")
		codigodebarras := r.FormValue("codigodebarras")
		quantidade := r.FormValue("quantidade")
		condicao := r.FormValue("condição")
		observacao := r.FormValue("observação")

		convertercodigo, err := strconv.Atoi(codigodebarras)
		if err != nil {
			log.Println("Error na conversão do código", err)
		}
		converterquantidade, err := strconv.Atoi(quantidade)
		if err != nil {
			log.Println("Error na conversão quantidade", err)
		}

		modelo.InserirEstoque(produto, converterquantidade, convertercodigo, condicao, observacao)
	}

	http.Redirect(w, r, "/mostra_estoque", 301)
}

func DeletarEstoque(w http.ResponseWriter, r *http.Request) {
	idDoEstoque := r.URL.Query().Get("id")
	modelo.DeletarEstoque(idDoEstoque)
	http.Redirect(w, r, "/mostra_estoque", 301)
}

func EditarEstoque(w http.ResponseWriter, r *http.Request) {
	idAtualizar := r.URL.Query().Get("id")
	estoque := modelo.EditarEstoque(idAtualizar)
	templ.ExecuteTemplate(w, "EditarEstoque", estoque)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		produto := r.FormValue("produto")
		codigodebarras := r.FormValue("codigodebarras")
		quantidade := r.FormValue("quantidade")
		condicao := r.FormValue("condição")
		observacao := r.FormValue("observação")

		idConvertidaParaInt, err := strconv.Atoi(id)
		if err != nil {
			log.Println("Erro na convesão do ID para int:", err)
		}

		convertercodigo, err := strconv.Atoi(codigodebarras)
		if err != nil {
			log.Println("Error na conversão do código", err)
		}
		converterquantidade, err := strconv.Atoi(quantidade)
		if err != nil {
			log.Println("Error na conversão quantidade", err)
		}

		modelo.Atualizaestoque(idConvertidaParaInt, produto, convertercodigo, converterquantidade, condicao, observacao)
	}
	http.Redirect(w, r, "/mostra_estoque", 301)
}
