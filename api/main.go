package main

import (
	"decider/api/config"
	"decider/api/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	config.Init()
	router := gin.Default()
	apiGroup := router.Group("/api")
	{

		router.GET("/healthCheck", handlers.HealthCheck)
		router.POST("/register", handlers.RegisterUser)
		router.POST("/validateToken", handlers.ValidateFirebaseToken)

		apiGroup.GET("/group/:groupId/item/:id", handlers.GetItem)
		apiGroup.DELETE("/group/:groupId/item/:id", handlers.DeleteItem)
		apiGroup.POST("/group/:groupId/item/:id", handlers.CreateItem)
		apiGroup.PUT("/group/:groupId/item/:id", handlers.EditItem)

		apiGroup.GET("/serverItem", handlers.GetServerItems)
		apiGroup.GET("/serverItem/:id", handlers.GetServerItemByID)
		apiGroup.GET("/voteByItem/:id", handlers.GetVotesByItemID)
		apiGroup.PUT("/VoteByItem/:id", handlers.VoteByItemID)

		apiGroup.POST("/commentToItem/:id", handlers.AddCommentToItemByID)
		apiGroup.DELETE("/commentToItem/:id", handlers.RemoveCommentFromItemByID)

	}

	router.Run("localhost:5050")
}
