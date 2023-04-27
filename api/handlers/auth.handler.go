package handlers

import (
	"context"
	"decider/api/config"
	"decider/api/models"
	"decider/api/utils"
	"encoding/json"
	"firebase.google.com/go/auth"
	"fmt"
	"net/http"
)

/*
RegisterUser
Registers a user Firebase Authentication and parses the response.
*/
func RegisterUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		utils.Logger(
			fmt.Sprintf("Unable to decode user object due to error: %v", err),
			"ERROR",
		)
		http.Error(w, err.Error(), http.StatusBadRequest)
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
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := models.RegisterResponse{ // Send response
		Message: fmt.Sprintf("User %s registered successfully", u.UID),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

/*
LoginUser
Logs a user in via Firebase Authentication and parses the response.
*/
func LoginUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		utils.Logger(
			fmt.Sprintf("Unable to decode user object due to error: %v", err),
			"ERROR",
		)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Authenticate user with Firebase auth
	u, err := config.FirebaseAuthClient.GetUserByEmail(context.Background(), user.Email)
	if err != nil {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}
	_, err = config.FirebaseAuthClient.SignInWithEmailAndPassword(context.Background(), user.Email, user.Password)
	if err != nil {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	// Generate JWT token
	token, err := authClient.CustomToken(context.Background(), u.UID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Send response
	response := LoginResponse{Token: token}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
