package handlers

import (
	"decider/api/db"
	"decider/api/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func GetItem(c *gin.Context) {
	groupID := c.Param("groupId")
	itemID := c.Param("id")
	// TODO: Implement logic to get the item with the given groupID and itemID

	item, err := db.GetItemFromDataStore(groupID, itemID)
	if err != nil {
		// Handle error (e.g., item not found)
		c.IndentedJSON(http.StatusNotFound, gin.H{
			"error": "Item not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"groupId": groupID,
		"itemId":  itemID,
		"item":    item,
	})
}

func DeleteItem(c *gin.Context) {
	groupID := c.Param("groupId")
	itemID := c.Param("id")
	// TODO: Implement logic to delete the item with the given groupID and itemID
	c.IndentedJSON(http.StatusOK, gin.H{
		"groupId": groupID,
		"itemId":  itemID,
		"message": "DeleteItem",
	})
}

func CreateItem(c *gin.Context) {
	groupID := c.Param("groupId")
	itemID := c.Param("id")
	// TODO: Implement logic to create an item with the given groupID and itemID
	c.IndentedJSON(http.StatusOK, gin.H{
		"groupId": groupID,
		"itemId":  itemID,
		"message": "CreateItem",
	})
}

func EditItem(c *gin.Context) {
	groupID := c.Param("groupId")
	itemID := c.Param("id")
	// TODO: Implement logic to edit the item with the given groupID and itemID
	c.IndentedJSON(http.StatusOK, gin.H{
		"groupId": groupID,
		"itemId":  itemID,
		"message": "EditItem",
	})
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
