package main

import (
	"fmt"

	"github.com/AndreaFabro/go-api-rest/models"
	"github.com/AndreaFabro/go-api-rest/routes"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{Id: 1, Nome: "Nome 1", Historia: "historia 1"},
		{Id: 2, Nome: "Nome 2", Historia: "historia 2"},
		{Id: 3, Nome: "Nome 3", Historia: "historia 3"},
	}

	fmt.Println("Iniciando o servidor rest com Go...")
	routes.HandleRequest()
}
