package controle

import (
	"html/template"
	"log"
	"modulo/modelo"
	"net/http"
	"strconv"
)

var templi = template.Must(template.ParseGlob("templates/*.html"))

func Registro(w http.ResponseWriter, r *http.Request) {
	templi.ExecuteTemplate(w, "Registro", nil)

}
func ListarRegistro(w http.ResponseWriter, r *http.Request) {
	todosAsEmpresas := modelo.ListaReg()
	templi.ExecuteTemplate(w, "ListarRegistro", todosAsEmpresas)

}

func Registrar(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		codigo := r.FormValue("codigo")
		tipodemovimentacao := r.FormValue("tipodemovimentacao")
		horario := r.FormValue("horario")
		observacao := r.FormValue("observacao")

		codigoConvertidaParaInt, err := strconv.Atoi(codigo)
		if err != nil {
			log.Println("Erro na convesão de quantidade:", err)
		}

		modelo.InserirRegistro(codigoConvertidaParaInt, tipodemovimentacao, horario, observacao)
	}
	http.Redirect(w, r, "/lista_registro", 301)
}

func DeletaRegistro(w http.ResponseWriter, r *http.Request) {
	idDEmpresa := r.URL.Query().Get("id")
	modelo.DeletaRegistro(idDEmpresa)
	http.Redirect(w, r, "/lista_registro", 301)
}

func Editaregistro(w http.ResponseWriter, r *http.Request) {
	idatualizar := r.URL.Query().Get("id")
	registro := modelo.Editareg(idatualizar)
	templi.ExecuteTemplate(w, "Editaregistro", registro)
}

func Atualizar(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		codigo := r.FormValue("codigo")
		tipodemovimentacao := r.FormValue("tipodemovimentacao")
		horario := r.FormValue("horario")
		observacao := r.FormValue("observacao")

		idConvertidaParaInt, err := strconv.Atoi(id)
		if err != nil {
			log.Println("Erro na convesão do ID para int:", err)
		}

		codigoConvertidaParaInt, err := strconv.Atoi(codigo)
		if err != nil {
			log.Println("Erro na convesão de quantidade:", err)
		}

		modelo.Atualizareg(idConvertidaParaInt, codigoConvertidaParaInt, tipodemovimentacao, horario, observacao)
	}
	http.Redirect(w, r, "/lista_registro", 301)
}
