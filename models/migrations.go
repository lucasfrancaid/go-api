package models

import "go-api/config"

func Migrator() {
	db := config.GetDatabase()

	db.AutoMigrate(&Author{})
	db.AutoMigrate(&Book{})
}
