package models

import (
	"os"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
)

var DB *gorm.DB

func ConnectDatabase() {
	host := os.Getenv("DBHOST")
	user := os.Getenv("DBUSER")
	pass := os.Getenv("DBPASS")
	dbname := os.Getenv("DBNAME")
	port := os.Getenv("DBPORT")
	dsn := "host="+host+" user="+user+" password="+pass+" dbname="+dbname+" port="+port+" sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
			panic("Failed to connect to database!")
	}

	err = db.AutoMigrate(&BusStop{})
	
	if err != nil {
			return
	}

	DB = db
}