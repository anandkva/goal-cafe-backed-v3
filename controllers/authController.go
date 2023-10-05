package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"

	"goal-cafe-backend/config"
	"goal-cafe-backend/models"
	"goal-cafe-backend/services"
	"goal-cafe-backend/utils"
)

const (
	statusCreated = http.StatusCreated
	statusOK      = http.StatusOK
)

func Register(ctx *gin.Context) {
	var newUser models.User
	if err := ctx.ShouldBindJSON(&newUser); err != nil {
		utils.HandleBadRequest(ctx, err)
		return
	}

	collection := config.DB.Collection("users")

	if utils.UserExists(collection, newUser.Email) {
		utils.HandleConflict(ctx, "Email address already registered")
		return
	}

	hashedPassword, err := utils.HashPassword(newUser.Password)
	if err != nil {
		utils.HandleInternalServerError(ctx, "Failed to create user")
		return
	}

	newUser.Password = hashedPassword

	if err := utils.InsertUser(collection, newUser); err != nil {
		utils.HandleInternalServerError(ctx, "Failed to create user")
		return
	}

	response := models.RegisterMessage{
		Code:    1,
		Message: "User created successfully",
		User:    newUser,
	}

	ctx.JSON(statusCreated, response)
}

func Login(ctx *gin.Context) {
	var loginRequest models.LoginUser
	if err := ctx.ShouldBindJSON(&loginRequest); err != nil {
		utils.HandleBadRequest(ctx, err)
		return
	}

	user, err := services.GetUserByEmail(loginRequest.Email)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			utils.HandleUnauthorizedError(ctx, "User not found")
		} else {
			utils.HandleInternalServerError(ctx, "Failed to query the database")
		}
		return
	}

	if !utils.ComparePasswords(loginRequest.Password, user.Password) {
		utils.HandleUnauthorizedError(ctx, "Invalid credentials")
		return
	}

	token, err := utils.GenerateJWTToken(*user)
	if err != nil {
		utils.HandleInternalServerError(ctx, "Failed to generate token")
		return
	}

	response := gin.H{"code": 1, "user": user, "token": token}
	ctx.JSON(statusOK, response)
}
