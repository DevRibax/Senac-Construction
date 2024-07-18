package controle

import (
	"html/template"
	"log"
	"modulo/modelo"
	"net/http"
	"strconv"
)

var temple = template.Must(template.ParseGlob("templates/*.html"))

func Funcionarios(w http.ResponseWriter, r *http.Request) {

	temple.ExecuteTemplate(w, "Funcionarios", nil)

}

func Inserir(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		cargo := r.FormValue("cargo")
		cpf := r.FormValue("cpf")
		rg := r.FormValue("rg")
		email := r.FormValue("email")
		senha := r.FormValue("senha")

		senhaconverte, err := strconv.Atoi(senha)
		if err != nil {
			log.Println("Erro na conversão da senha:", err)
		}

		modelo.Insira(nome, cargo, cpf, rg, email, senhaconverte)
	}

	http.Redirect(w, r, "/lista_funcionarios", 301)
}

func ListaFuncionario(w http.ResponseWriter, r *http.Request) {
	funcionarioinserir := modelo.Trazerfuncionario()
	temple.ExecuteTemplate(w, "ListaFuncionario", funcionarioinserir)
}

func Deletarfunc(w http.ResponseWriter, r *http.Request) {
	deletefun := r.URL.Query().Get("id")
	modelo.Deletafuncionario(deletefun)
	http.Redirect(w, r, "/lista_funcionarios", 301)
}
