package routes

import (
	"goal-cafe-backend/controllers"

	"github.com/gin-gonic/gin"
)

func WelcomeRoute(router *gin.Engine) {
	router.GET("/", controllers.WelcomeController)
}
