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
	inits.DB.AutoMigrate(&models.Post{})
}