package controle

import (
	"html/template"
	"net/http"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "Index", nil)
}

func Sobrenos(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "Sobrenos", nil)
}
