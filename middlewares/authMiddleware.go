package middlewares

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/robwestbrook/goblogart/inits"
	"github.com/robwestbrook/goblogart/models"
)

func RequireAuth(ctx *gin.Context) {

	// Get token from cookie and check for errors
	tokenString, err := ctx.Cookie("Authorization")
	if err != nil {
		ctx.JSON(401, gin.H{"error": "unauthorized"})
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	// Check token and check for errors
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		// Get secret from environment variable
		return []byte(os.Getenv("SECRET")), nil
	})

	// Check if token is expired
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			ctx.JSON(401, gin.H{"error": "unauthorized"})
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// Create a variable of User model type
	var user models.User

		// Get the user by id found in token and check
		// for valid user
		inits.DB.First(&user, int(claims["id"].(float64)))
		if user.ID == 0 {
			ctx.JSON(401, gin.H{"error": "unauthorized"})
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// Set the user into the context
		ctx.Set("user", user)
		fmt.Println(claims["foo"], claims["nbf"])
	} else {

		// Token is not present or invalid
		ctx.AbortWithStatus(http.StatusUnauthorized)
	}

	// Complete middleware function and proceed
	ctx.Next()
}