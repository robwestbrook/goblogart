package inits

import (
	"log"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DBInit() {
	dbName := os.Getenv("DB_NAME")
	db, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if err != nil {
		log.Println("Failed to connect to database")
		panic("Failed to connect to database")
	}

	log.Println("Connected to database", dbName)
	DB = db
}