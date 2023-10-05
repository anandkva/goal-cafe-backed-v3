package routes

import (
	"goal-cafe-backend/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoute(router *gin.Engine) {
	authRoute := router.Group("/auth")
	authRoute.POST("/register", controllers.Register)
	authRoute.POST("/login", controllers.Login)
}
