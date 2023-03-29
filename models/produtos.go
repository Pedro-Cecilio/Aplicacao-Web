package models

import (
	"log"

	"github.com/Pedro-Cecilio/Aplicacao-Web/db"
)

type Produtos struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func BuscaTodosProdutos() []Produtos {
	db := db.ConectaComBancoDeDados()
	selectAll, err := db.Query("SELECT * FROM produtos ORDER BY produtosId ASC")
	if err != nil {
		panic(err.Error())
	}
	p := Produtos{}
	produtos := []Produtos{}
	for selectAll.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = selectAll.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}
		p.Id = id
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

		produtos = append(produtos, p)
	}
	defer db.Close()
	return produtos
}

func CriaNovoProduto(nome, descricao string, preco float64, quantidade int) {
	db := db.ConectaComBancoDeDados()
	inserirProduto, err := db.Prepare("INSERT INTO produtos(nome, descricao, preco, quantidade) VALUES ($1, $2, $3, $4)")
	if err != nil {
		log.Println("Erro ao inserir produto:", err)
	}

	inserirProduto.Exec(nome, descricao, preco, quantidade)

	defer db.Close()
}

func DeletaProduto(idDoProduto int){
	db := db.ConectaComBancoDeDados()
	
	deletar, err := db.Prepare("DELETE FROM produtos WHERE produtosId = $1")
	if err != nil {
		log.Println("Erro ao preparar requisição:", err)
	}

	deletar.Exec(idDoProduto)
	defer db.Close()
}

func AtualizarProduto(nome, descricao string, preco float64, quantidade, idDoProduto int){
	db := db.ConectaComBancoDeDados()

	atualizar, err := db.Prepare("UPDATE produtos SET nome = $1, descricao = $2, preco = $3, quantidade = $4 WHERE produtosId = $5 ")
	if err != nil {
		log.Println("Erro ao preparar atualização:", err)
	}

	atualizar.Exec(nome, descricao, preco, quantidade, idDoProduto)
	defer db.Close()
}

func RetornaProduto(idDoProduto int) (produto Produtos){
	db := db.ConectaComBancoDeDados()
	produtoEspecifico, err := db.Query("SELECT * FROM produtos WHERE produtosId = $1", idDoProduto)
	if err != nil {
		log.Println("Erro na requisição:", err)
	}
	produtoParaAtualizar := Produtos{}
	for produtoEspecifico.Next(){
		var nome, descricao string
		var preco float64
		var quantidade, id int

		produtoEspecifico.Scan(&id, &nome, &descricao, &preco, &quantidade)
		produtoParaAtualizar.Id = id
		produtoParaAtualizar.Nome = nome
		produtoParaAtualizar.Descricao = descricao
		produtoParaAtualizar.Preco = preco
		produtoParaAtualizar.Quantidade = quantidade
	}
	defer db.Close()
	return produtoParaAtualizar
}