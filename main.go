package main

import (
	"fmt"
	"net/http"

	"github.com/spohess/routes"
)

func main() {
	fmt.Println("Carregando todas as rotas")
	routes.LoadRoutes()
	fmt.Println("Subindo servidor na porta 8000")
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Println("Servidor caiu com erro:", err)
	}
}
