package main

import (
	"github.com/robwestbrook/goblogart/inits"
	"github.com/robwestbrook/goblogart/models"
)

func init() {
	inits.LoadEnv()
	inits.DBInit()
}

func main() {
	// Auto migrate models
	inits.DB.AutoMigrate(&models.Post{})
	inits.DB.AutoMigrate(&models.User{})
}