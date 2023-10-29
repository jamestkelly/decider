package handlers

import (
	"decider/api/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

/*
HealthCheck
@Summary Performs a basic health-check.
@Description Health-check API call to verify that the API server is running.
@Tags Users
@Success 200 {object} models.HealthCheckResponse
@Failure 404 {object} object
@Router /healthCheck [get]
*/
func HealthCheck(c *gin.Context) {
	payload := models.HealthCheckResponse{
		Date:    time.Now(),
		Message: "API is live.",
	}
	c.IndentedJSON(http.StatusOK, payload)
}
