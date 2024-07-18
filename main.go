package main

import (
	"fmt"
	"modulo/rotas"

	"net/http"
)

func main() {
	rotas.CarregarRotas()
	fmt.Println("SERVIDOR RODANDO NA PORTA 8080")
	http.ListenAndServe(":8080", nil)
}
