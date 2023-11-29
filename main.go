package main

import (
	"github.com/gin-gonic/gin"
	"github.com/robwestbrook/goblogart/controllers"
	"github.com/robwestbrook/goblogart/inits"
	"github.com/robwestbrook/goblogart/middlewares"
)

// Initialize app
func init() {
	inits.LoadEnv()
	inits.DBInit()
}

func main() {
	// Initialize gin framework
	r := gin.Default()

	// Define Post Routes
	r.POST("/", middlewares.RequireAuth, controllers.CreatePost)
	r.GET("/", controllers.GetPosts)
	r.GET("/:id", controllers.GetPost)
	r.PUT("/:id", controllers.UpdatePost)
	r.DELETE("/:id", controllers.DeletePost)

	// Define User Routes
	r.POST("/user", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/auth", controllers.Validate)
	r.GET("/users", controllers.GetUsers)
	r.GET("/logout", controllers.Logout)

	// Start the server
	r.Run()
}