package handlers

import (
	"context"
	"decider/api/config"
	"decider/api/models"
	"decider/api/utils"
	"firebase.google.com/go/auth"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
RegisterUser
@Summary Register user.
@Description Registers a user Firebase Authentication and parses the response.
@Tags Users
@Success 200 {object} models.User
@Failure 404 {object} object
@Router /register [post]
*/
func RegisterUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		utils.Logger(
			fmt.Sprintf("Unable to decode user object due to error: %v", err),
			"ERROR",
		)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	params := (&auth.UserToCreate{}). // Create user in Firebase auth
						Email(user.Email).
						Password(user.Password)
	u, err := config.FirebaseAuthClient.CreateUser(context.Background(), params)
	if err != nil {
		utils.Logger(
			fmt.Sprintf("Unable to create user object due to error: %v", err),
			"ERROR",
		)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := models.RegisterResponse{ // Send response
		Message: fmt.Sprintf("User %s registered successfully", u.UID),
	}

	c.IndentedJSON(http.StatusOK, response)
}

/*
ValidateFirebaseToken
@Summary Validate user authentication
@Description Takes a token generated by a user on the client and verifies and decodes the supplied ID token for subsequent calls
to the API.
@Tags Users
@Success 200 {object} models.User
@Failure 400,500,401 {object} object
@Router /validateToken [post]
*/
func ValidateFirebaseToken(c *gin.Context) {
	var user models.User
	token := c.GetHeader("Authorization")
	if token == "" {
		c.AbortWithStatus(http.StatusUnauthorized)
		utils.Logger(
			fmt.Sprintf("Unauthorised user call to API: %v", http.StatusUnauthorized),
			"WARN",
		)
		return
	}

	// Verify the ID token
	idToken, err := config.FirebaseAuthClient.VerifyIDToken(context.Background(), token)
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		utils.Logger(
			fmt.Sprintf("Unauthorised user, unable to decode token due to error: %v", http.StatusUnauthorized),
			"WARN",
		)
		return
	}

	user.ID = idToken.UID
	user.Email = idToken.Claims["email"].(string) // TODO: Implement logic to do something with validated token

	// TODO: Return a response to the
	c.IndentedJSON(http.StatusOK, "yes")
}
