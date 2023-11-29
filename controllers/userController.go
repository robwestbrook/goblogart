package controllers

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/robwestbrook/goblogart/inits"
	"github.com/robwestbrook/goblogart/models"
	"golang.org/x/crypto/bcrypt"
)

//
// User Signup
//
func Signup(ctx *gin.Context) {

	// Initialize a body struct for incoming data
	var body struct {
		Name			string
		Email			string
		Password	string
	}

	// Bind the JSON data to the body and check for errors
	if ctx.BindJSON(&body) != nil {
		ctx.JSON(400, gin.H{"error": "bad request"})
		return
	}

	// Hash the password and check for errors
	hash, err := bcrypt.GenerateFromPassword(
		[]byte(body.Password), 10)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err})
		return
	}

	// Bind the data to the user
	user := models.User{
		Name: body.Name,
		Email: body.Email,
		Password: string(hash),
	}

	// Create user in database and check for error
	result := inits.DB.Create(&user)
	if result.Error != nil {
		ctx.JSON(500, gin.H{"error": result.Error})
		return
	}

	// Return JSON
	ctx.JSON(200, gin.H{"data": user})
}

//
// Login User
//
func Login(ctx *gin.Context) {

	// Initialize a body struct for incoming data
	var body struct {
		Email			string
		Password	string
	}

	// Bind the JSON data to the body and check for errors
	if ctx.BindJSON(&body) != nil {
		ctx.JSON(400, gin.H{"error": "bad request"})
		return
	}

	// Initialize a user variable
	var user models.User

	// Find user in database by email and check for errors
	result := inits.DB.Where("email = ?", body.Email).First(&user)
	if result.Error != nil {
		ctx.JSON(500, gin.H{"error": "User not found"})
		return
	}

	// Validate password and return error if it doesn't match
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		ctx.JSON(401, gin.H{"error": "unauthorized"})
		return
	}

	// Generate a JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	// Sign the JWT token and check for errors
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		ctx.JSON(500, gin.H{"error": "error signing token"})
		return
	}

	// Create a cookie that expires in 30 days
	ctx.SetSameSite(http.SameSiteLaxMode)
	ctx.SetCookie(
		"Authorization", 
		tokenString,
		3600 *24*30,
		"",
		"localhost",
		false,
		true)

}

//
// Get Users and their posts from the database
//
func GetUsers(ctx *gin.Context) {

	// Create a slice to hold users of User model type
	var users []models.User

	// Get all users and their posts and check for errors
	err := inits.DB.Model(&models.User{}).Preload("Posts").Find(&users).Error
	if err != nil {
		fmt.Println(err)
		ctx.JSON(500, gin.H{"error": "Error getting users"})
		return
	}

	// Return JSON
	ctx.JSON(200, gin.H{"data": users})
}

//
// Use validation middleware to validate token and password
//
func Validate(ctx *gin.Context) {

	// Get the user from the context and check for errors
	user, err := ctx.Get("user")
	if err != false {
		ctx.JSON(500, gin.H{"error": err})
		return
	}

	// Return JSON
	ctx.JSON(200, gin.H{"data": "You are logged in", "user": user})
}

//
// Logout User
//
func Logout(ctx *gin.Context) {

	// Clear all cookie data
	ctx.SetSameSite(http.SameSiteLaxMode)
	ctx.SetCookie(
		"Authorization",
		"",
		-1,
		"",
		"localhost",
		false,
		true)
	ctx.JSON(200, gin.H{"data": "You are logged out"})
}