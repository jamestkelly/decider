package main

import (
	"decider/api/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/healthCheck", handlers.HealthCheck)
	router.POST("/register", handlers.RegisterUser)
	router.POST("/validateToken", handlers.ValidateFirebaseToken)
	apiGroup := router.Group("/api") 
	{
		apiGroup.GET("/item", handlers.GetItem)
		apiGroup.DELETE("/item", handlers.DeleteItem)
		apiGroup.POST("/item", handlers.CreateItem)
		apiGroup.PUT("/item", handlers.EditItem)

		apiGroup.GET("/serverItem", handlers.GetServerItems)
		apiGroup.GET("/serverItem/:id", handlers.GetServerItemByID)
		apiGroup.GET("/voteByItem/:id", handlers.GetVotesByItemID)
		apiGroup.PUT("/VoteByItem/:id", handlers.VoteByItemID)

		apiGroup.POST("/commentToItem/:id", handlers.AddCommentToItemByID)
		apiGroup.DELETE("/commentToItem/:id", handlers.RemoveCommentFromItemByID)

	}

	router.Run("localhost:5050")
}
