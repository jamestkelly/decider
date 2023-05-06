package config

import (
	"context"
	"decider/api/utils"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"fmt"
	"google.golang.org/api/option"
)

var (
	FirebaseApp        *firebase.App
	FirebaseAuthClient *auth.Client
)

/*
*
InitFirebase
Lorem Ipsum
*/
func InitFirebase() {
	opt := option.WithCredentialsFile("../firebaseCredentials.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		utils.Logger(
			fmt.Sprintf("Error initialising Firebase app: %v", err),
			"ERROR",
		)
	}

	auth, err := FirebaseApp.Auth(context.Background())
	if err != nil {
		utils.Logger(
			fmt.Sprintf("Error initialising Firebase Auth: %v", err),
			"ERROR",
		)
	}

	FirebaseApp = app
	FirebaseAuthClient = auth
}
