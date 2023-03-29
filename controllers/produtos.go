package controllers

import (
	"log"
	"net/http"
	"strconv"
	"text/template"

	"github.com/Pedro-Cecilio/Aplicacao-Web/models"
)

var temp = template.Must(template.ParseGlob("templates/*html"))

func Index(w http.ResponseWriter, r *http.Request) {
	produtos := models.BuscaTodosProdutos()
	temp.ExecuteTemplate(w, "index.html", produtos)
}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "new.html", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco, err := strconv.ParseFloat(r.FormValue("preco"), 64)
		if err != nil {
			log.Println("Erro na conversão do preço:", err)
		}
		quantidade, err := strconv.Atoi(r.FormValue("quantidade"))
		if err != nil {
			log.Println("Erro na conversão da quantidade:", err)
		}
		models.CriaNovoProduto(nome, descricao, preco, quantidade)
	}
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idDoProduto, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		log.Println("Erro ao convertar o id:", err)
	}
	models.DeletaProduto(idDoProduto)
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	idDoProduto, _ := strconv.Atoi(r.URL.Query().Get("id"))
	Produto := models.RetornaProduto(idDoProduto)
	temp.ExecuteTemplate(w, "edit.html", Produto)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id, _ := strconv.Atoi(r.FormValue("id"))
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco, _ := strconv.ParseFloat(r.FormValue("preco"), 64)
		quantidade, _ := strconv.Atoi(r.FormValue("quantidade"))
		models.AtualizarProduto(nome, descricao, preco, quantidade, id)
		http.Redirect(w, r, "/", http.StatusMovedPermanently)
	}
}
