package config

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func DatabaseConnection() {
	settings := GetSettings()

	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v",
		settings.db.Host, settings.db.User, settings.db.Password,
		settings.db.Name, settings.db.Port)

	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db = conn
}

func GetDatabase() *gorm.DB {
	return db
}
