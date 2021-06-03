package models

import "github.com/lucasfrancaid/go-api/config"

func Migrator() {
	db := config.GetDatabase()

	db.AutoMigrate(&Author{})
	db.AutoMigrate(&Book{})
}
