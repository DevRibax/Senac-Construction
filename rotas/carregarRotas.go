package rotas

import (
	"modulo/controle"
	"net/http"
)

func CarregarRotas() {

	http.HandleFunc("/", controle.Index)
	http.HandleFunc("/login", controle.Login)
	http.HandleFunc("/cad_funcionarios", controle.Funcionarios)
	http.HandleFunc("/cad_insp", controle.Insp)
	http.HandleFunc("/cad_material", controle.Material)
	http.HandleFunc("/notafiscal", controle.NotaFiscal)
	http.HandleFunc("/cad_fornecedor", controle.Fornecedor)
	http.HandleFunc("/editar_fornecedor", controle.Editarfornecedor)
	http.HandleFunc("/altereforn", controle.AlteraFornecedor)
	http.HandleFunc("/usuario", controle.Usuario)
	http.HandleFunc("/inseri", controle.Inseri)
	http.HandleFunc("/editar_mater", controle.Edit)
	http.HandleFunc("/delete", controle.Delete)
	http.HandleFunc("/atualmater", controle.Atualiza)
	http.HandleFunc("/mostra_cad", controle.Mostracad)
	http.HandleFunc("/sobrenos", controle.Sobrenos)
	http.HandleFunc("/mostra_estoque", controle.Mostraestoque)
	http.HandleFunc("/deletarestoque", controle.DeletarEstoque)
	http.HandleFunc("/editar_estoque", controle.EditarEstoque)
	http.HandleFunc("/update", controle.Update)
	http.HandleFunc("/colocar", controle.Colocar)
	http.HandleFunc("/adicionar", controle.Adicionar)
	http.HandleFunc("/listar_fornecedor", controle.ListarFornecedor)
	http.HandleFunc("/deleta_fornecedor", controle.Deletefornecedor)
	http.HandleFunc("/inserir", controle.Inserir)
	http.HandleFunc("/lista_funcionarios", controle.ListaFuncionario)
	http.HandleFunc("/deletefunc", controle.Deletarfunc)
	http.HandleFunc("/cad_registro", controle.Registro)
	http.HandleFunc("/lista_registro", controle.ListarRegistro)
	http.HandleFunc("/registrar", controle.Registrar)
	http.HandleFunc("/atualiza", controle.Atualizar)
	http.HandleFunc("/deleta_registro", controle.DeletaRegistro)
	http.HandleFunc("/editar_registro", controle.Editaregistro)
	http.HandleFunc("/pegarnota", controle.ColocarNotasFiscais)
	http.HandleFunc("/listar_notafiscal", controle.MostraNotasFiscais)
	http.HandleFunc("/deletanota", controle.Deletarnotafiscal)
	http.HandleFunc("/editar_notafiscal", controle.Editarnotafiscal)
	http.HandleFunc("/nota", controle.Nota)

	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("static"))))
}
