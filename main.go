package main

import (
	"net/http"

	"github.com/Pedro-Cecilio/Aplicacao-Web/routes"
	_ "github.com/lib/pq"
)

func main() {
	routes.CarregaRotas()
	http.ListenAndServe("localhost:8000", nil)
}
