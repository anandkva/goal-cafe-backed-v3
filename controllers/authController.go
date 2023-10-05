package controllers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"goal-cafe-backend/config"
	"goal-cafe-backend/models"
	"goal-cafe-backend/utils"
)

// Register handles user registration.
func Register(c *gin.Context) {
	var newUser models.User
	fmt.Printf("newUser: %v", newUser)
	if err := c.ShouldBindJSON(&newUser); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	collection := config.DB.Collection("users")

	// Check if the user with the provided email already exists
	existingUser := models.User{}
	err := collection.FindOne(context.TODO(), bson.M{"email": newUser.Email}).Decode(&existingUser)
	if err == nil {
		c.JSON(http.StatusConflict, gin.H{"code": 2, "message": "Email address already registered"})
		return
	}
	fmt.Printf("User::%v", newUser)
	_, err = collection.InsertOne(context.TODO(), newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	message := models.RegisterMessage{
		Code:    1,
		Message: "User created successfully",
		User:    newUser,
	}

	c.JSON(http.StatusCreated, message)
}

func Login(c *gin.Context) {
	var loginRequest models.LoginUser
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	collection := config.DB.Collection("users")
	user := models.User{}

	// Check if the user with the provided email exists
	err := collection.FindOne(context.TODO(), bson.M{"email": loginRequest.Email}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusUnauthorized, gin.H{"code": 2, "message": "User not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query the database"})
		}
		return
	}

	// Without hashing, simply compare the stored password with the login request password
	storedPassword := string(user.Password)
	loginPassword := loginRequest.Password

	if storedPassword != loginPassword {
		// Passwords don't match
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Generate a JWT token
	token, err := utils.GenerateJWTToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	// Send the token as a response
	c.JSON(http.StatusOK, gin.H{"code": 1, "user": user, "token": token})
}
