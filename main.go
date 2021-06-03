package main

import (
	"github.com/lucasfrancaid/go-api/config"
	"github.com/lucasfrancaid/go-api/models"
	"github.com/lucasfrancaid/go-api/router"
)

func main() {

	r := router.Routes()
	config.DatabaseConnection()
	models.Migrator()

	r.Run()
}
