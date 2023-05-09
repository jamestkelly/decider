package main

import (
	"decider/api/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/healthCheck", handlers.HealthCheck)
	router.POST("/register", handlers.RegisterUser)
	router.POST("validateToken", handlers.ValidateFirebaseToken)
	router.Run("localhost:5050")
}
