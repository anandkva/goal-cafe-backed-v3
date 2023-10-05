package main

import (
	"fmt"
	"goal-cafe-backend/config"
	"goal-cafe-backend/routes"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()

	config.ConnectDB()

	routes.AuthRoute(app)

	app.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "Hello GoalCafe App"})
	})

	if err := app.Run(":8080"); err == nil {
		fmt.Printf("Error:: %v", err)
	}

}
