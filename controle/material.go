package controle

import (
	"html/template"
	"log"
	"modulo/modelo"
	"net/http"
	"strconv"
)

var tmp = template.Must(template.ParseGlob("templates/*.html"))

func Mostracad(w http.ResponseWriter, r *http.Request) {
	pegartodososprodutos := modelo.TrazerTodosOsMateriais()
	tmp.ExecuteTemplate(w, "Mostracad", pegartodososprodutos)

}

func Material(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "Material", nil)
}

func Inseri(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		tipo := r.FormValue("categoria")
		categoria := r.FormValue("tipo")
		descricao := r.FormValue("descricao")

		modelo.InserirMaterial(categoria, tipo, descricao)
	}

	http.Redirect(w, r, "/mostra_cad", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	iddomaterial := r.URL.Query().Get("id")
	modelo.DeletarMaterial(iddomaterial)
	http.Redirect(w, r, "/mostra_cad", 301)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	idatualizar := r.URL.Query().Get("id")
	material := modelo.EditarMaterial(idatualizar)
	temp.ExecuteTemplate(w, "Edit", material)
}

func Atualiza(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		categoria := r.FormValue("categoria")
		tipo := r.FormValue("tipo")
		descricao := r.FormValue("descricao")

		idConvertidaParaInt, err := strconv.Atoi(id)
		if err != nil {
			log.Println("Erro na convesão do ID para int:", err)
		}

		modelo.Atualizamater(idConvertidaParaInt, categoria, tipo, descricao)
	}

	http.Redirect(w, r, "/mostra_cad", 301)
}
