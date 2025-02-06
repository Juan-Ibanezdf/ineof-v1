package main

import (
	"log"

	api "github.com/Juan-Ibanezdf/ineof-v1/cmd"
	// Pacotes do Swagger
	_ "github.com/Juan-Ibanezdf/ineof-v1/docs"
)

// @title Ineof API
// @version 1.0
// @description API para gerenciamento de campanhas, equipamentos e dados LIDAR.
// @host localhost:8080
// @BasePath /api/v1
func main() {
	// Inicia o servidor
	server := api.NewServer()
	err := server.Start()
	if err != nil {
		log.Fatal("Erro ao iniciar o servidor:", err)
	}
}
