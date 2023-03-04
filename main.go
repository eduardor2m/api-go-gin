package main

import (
	"github.com/eduardor2m/api-go-gin/database"
	"github.com/eduardor2m/api-go-gin/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequest()
}
