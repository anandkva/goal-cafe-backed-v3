package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func WelcomeController(controller *gin.Context) {

	welcomeMessage := "Welcome GoalCafe Projects"

	controller.JSON(http.StatusOK, welcomeMessage)
}
