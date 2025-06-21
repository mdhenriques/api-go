package main

import (
	"github.com/mdhenriques/api-go/database"
	"github.com/mdhenriques/api-go/routes"
)

func main() {
	database.Connect()
	r := routes.SetupRouter()
	r.Run()
}
