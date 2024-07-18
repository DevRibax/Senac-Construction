package controle

import (
	"html/template"
	"modulo/modelo"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var templet = template.Must(template.ParseGlob("templates/*.html"))

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		templet.ExecuteTemplate(w, "Login", nil)
	} else if r.Method == "POST" {
		email := r.FormValue("email")
		senhaStr := r.FormValue("senha")
		senha, err := strconv.Atoi(senhaStr)
		if err != nil {
			data := map[string]string{"Error": "Invalid password format"}
			templet.ExecuteTemplate(w, "Login", data)
			return
		}
		if modelo.AuthenticateUser(email, senha) {
			http.Redirect(w, r, "/home", http.StatusFound)
		} else {
			data := map[string]string{"Error": "Email ou senha incorretos."}
			templet.ExecuteTemplate(w, "Login", data)
		}
	}
}

func Usuario(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		templet.ExecuteTemplate(w, "Usuario", nil)
	} else if r.Method == "POST" {
		email := r.FormValue("email")
		senhaStr := r.FormValue("senha")
		senha, err := strconv.Atoi(senhaStr)
		if err != nil {
			templet.ExecuteTemplate(w, "Usuario", "Invalid password format")
			return
		}
		modelo.Insert(email, senha)
		http.Redirect(w, r, "/login", http.StatusFound)
	}
}

func Home(c *gin.Context) {
	c.HTML(http.StatusOK, "tela.html", nil)
}
