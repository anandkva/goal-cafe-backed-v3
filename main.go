package main

import (
	"fmt"
	"goal-cafe-backend/config"
	"goal-cafe-backend/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()

	config.ConnectDB()

	routes.WelcomeRoute(app)
	routes.AuthRoute(app)

	fmt.Printf("hello World\n")

	if err := app.Run(":8080"); err == nil {
		fmt.Printf("Error:: %v", err)
	}
}
