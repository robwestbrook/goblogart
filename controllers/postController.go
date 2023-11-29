package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/robwestbrook/goblogart/inits"
	"github.com/robwestbrook/goblogart/models"
)

//
// Add a post to the database
//
func CreatePost(ctx *gin.Context) {

	// Define a struct
	var body struct {
		Title		string
		Body		string
		Likes		int
		Draft		bool
		Author	string
	}

	// Bind body var to JSON
	ctx.BindJSON(&body)

	// Add info to post variable
	post := models.Post{
		Title: body.Title,
		Body: body.Body,
		Likes: body.Likes,
		Draft: body.Draft,
		Author: body.Author,
	}

	fmt.Println(post)

	// Create post in database and check for errors
	result := inits.DB.Create(&post)
	if result.Error != nil {
		ctx.JSON(500, gin.H{"error": result.Error})
		return
	}

	// Return JSON
	ctx.JSON(200, gin.H{"data": post})
}

//
// Get all Posts from database
//
func GetPosts(ctx *gin.Context) {

	// Initialize a slice to hold posts with Post model type
	var posts []models.Post

	// Get all posts from database and check for error
	result := inits.DB.Find(&posts)
	if result.Error != nil {
		ctx.JSON(500, gin.H{"error": result.Error})
		return
	}

	// Return JSON
	ctx.JSON(200, gin.H{"data": posts})
}

//
// Get a single Post from database
//
func GetPost(ctx *gin.Context) {

	// Initial a post variable with Post model type
	var post models.Post

	// Get single post from database and check for error
	result := inits.DB.First(&post, ctx.Param("id"))
	if result.Error != nil {
		ctx.JSON(500, gin.H{"error": result.Error})
		return
	}

	// Return JSON
	ctx.JSON(200, gin.H{"data": post})
}

//
// Update a Post in the database
//
func UpdatePost(ctx *gin.Context) {

	// Initialize a var with the Post model type
	var body struct {
		Title			string
		Body			string
		Likes			int
		Draft			bool
		Author		string
	}

	// Bind body to JSON
	ctx.BindJSON(&body)

	// Initialize a var with Post model type
	var post models.Post

	// Get the post and load into post variable
	// and check for error
	result := inits.DB.First(&post, ctx.Param("id"))
	if result.Error != nil {
		ctx.JSON(500, gin.H{"error": result.Error})
		return
	}

	// Update post
	inits.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body: body.Body,
		Likes: body.Likes,
		Draft: body.Draft,
		Author: body.Author,
	})

	// Return JSON
	ctx.JSON(200, gin.H{"data": post})
}

//
// Delete a Post from the database
//
func DeletePost(ctx *gin.Context) {

	// Get ID from content
	id := ctx.Param("id")

	// Delete post from database
	inits.DB.Delete(&models.Post{}, id)

	// Return JSON
	ctx.JSON(200, gin.H{"data": "Post deleted successfully"})
}