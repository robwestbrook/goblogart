package main

import (
	"github.com/gin-gonic/gin"
	"github.com/robwestbrook/goblogart/controllers"
	"github.com/robwestbrook/goblogart/inits"
)

// Initialize app
func init() {
	inits.LoadEnv()
	inits.DBInit()
}

func main() {
	// Initialize gin framework
	r := gin.Default()

	// Define Routes
	r.POST("/", controllers.CreatePost)
	r.GET("/", controllers.GetPosts)
	r.GET("/:id", controllers.GetPost)
	r.PUT("/:id", controllers.UpdatePost)
	r.DELETE("/:id", controllers.DeletePost)

	// Start the server
	r.Run()
}