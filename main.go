package main

import (
	"github.com/mdhenriques/api-go/database"
	"github.com/mdhenriques/api-go/routes"
)

// @title API Exemplo
// @version 1.0
// @description Esta é uma API de exemplo com autenticação JWT
// @host localhost:8080
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	database.Connect()
	r := routes.SetupRouter()
	r.Run()
}
