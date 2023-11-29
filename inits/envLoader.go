package inits

import (
	"log"

	"github.com/joho/godotenv"
)

// Load environment variables
func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}