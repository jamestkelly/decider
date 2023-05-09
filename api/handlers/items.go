package handlers

import (
	"decider/api/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func GetItem(c *gin.Context) {
	payload := models.HealthCheckResponse{
		Date:    time.Now(),
		Message: "test",
	}
	c.IndentedJSON(http.StatusOK, payload)
}
func CreateItem(c *gin.Context) {
	payload := models.HealthCheckResponse{
		Date:    time.Now(),
		Message: "test",
	}
	c.IndentedJSON(http.StatusOK, payload)
}
func DeleteItem(c *gin.Context) {
	payload := models.HealthCheckResponse{
		Date:    time.Now(),
		Message: "test",
	}
	c.IndentedJSON(http.StatusOK, payload)
}
func EditItem(c *gin.Context) {
	payload := models.HealthCheckResponse{
		Date:    time.Now(),
		Message: "test",
	}
	c.IndentedJSON(http.StatusOK, payload)
}
func GetServerItems(c *gin.Context) {
	payload := models.HealthCheckResponse{
		Date:    time.Now(),
		Message: "test",
	}
	c.IndentedJSON(http.StatusOK, payload)
}
func GetServerItemByID(c *gin.Context) {
	payload := models.HealthCheckResponse{
		Date:    time.Now(),
		Message: "test",
	}
	c.IndentedJSON(http.StatusOK, payload)
}
func GetVotesByItemID(c *gin.Context) {
	payload := models.HealthCheckResponse{
		Date:    time.Now(),
		Message: "test",
	}
	c.IndentedJSON(http.StatusOK, payload)
}
func VoteByItemID(c *gin.Context) {
	payload := models.HealthCheckResponse{
		Date:    time.Now(),
		Message: "test",
	}
	c.IndentedJSON(http.StatusOK, payload)
}
func AddCommentToItemByID(c *gin.Context) {
	payload := models.HealthCheckResponse{
		Date:    time.Now(),
		Message: "test",
	}
	c.IndentedJSON(http.StatusOK, payload)
}
func RemoveCommentFromItemByID(c *gin.Context) {
	payload := models.HealthCheckResponse{
		Date:    time.Now(),
		Message: "test",
	}
	c.IndentedJSON(http.StatusOK, payload)
}
