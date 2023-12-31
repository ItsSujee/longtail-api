package models

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	host := os.Getenv("DBHOST")
	user := os.Getenv("DBUSER")
	pass := os.Getenv("DBPASS")
	dbname := os.Getenv("DBNAME")
	port := os.Getenv("DBPORT")
	dsn := "host=" + host + " user=" + user + " password=" + pass + " dbname=" + dbname + " port=" + port + " sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to Postgres database!")
	}

	if err != nil {
		return
	}

	DB = db
}
