package config

import (
	"fmt"
	"os"
	"strconv"
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

	db := Database{
		Host:     db_host,
		Port:     db_port,
		Name:     db_name,
		User:     db_user,
		Password: db_password,
	}
	return db
}
