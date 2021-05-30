package config

import (
	"fmt"
	"os"
	"strconv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	Host     string
	Port     int64
	Name     string
	User     string
	Password string
}

func GetDatabaseSettings() Database {
	db_host := os.Getenv("DB_HOST")
	db_name := os.Getenv("DB_NAME")
	db_user := os.Getenv("DB_USER")
	db_password := os.Getenv("DB_PASSWORD")
	db_port, portErr := strconv.ParseInt(os.Getenv("DB_PORT"), 10, 64)

	if portErr != nil {
		fmt.Println("Settings Error", portErr)
	}

	database := Database{
		Host:     db_host,
		Port:     db_port,
		Name:     db_name,
		User:     db_user,
		Password: db_password,
	}
	return database
}

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
