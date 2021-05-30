package main

import (
	"go-api/config"
	"go-api/models"
	"go-api/router"
)

func main() {

	r := router.Routes()
	config.DatabaseConnection()
	models.Migrator()

	r.Run()
}
