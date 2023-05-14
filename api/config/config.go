package config

import (
	"context"
	"decider/api/utils"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/db"
	"firebase.google.com/go/auth"
	"fmt"
	"github.com/joho/godotenv"
	"google.golang.org/api/option"
	"os"
)

var (
	FirebaseClient     *db.Client
	FirebaseAuthClient *auth.Client
)

/*
InitFirebase initializes the Firebase Admin SDK and sets up the FirebaseAuthClient.
*/
func InitFirebaseApp() *firebase.App {
	opt := option.WithCredentialsFile("../firebaseCredentials.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		utils.Logger(
			fmt.Sprintf("Error initialising Firebase app: %v", err),
			"ERROR",
		)
		os.Exit(1)
	}

	return app
}

func InitClient(app *firebase.App) (*db.Client, *auth.Client) {

	client, err := app.Database(context.Background())
	if err != nil {
		fmt.Printf("Error initializing Firebase Realtime Database client: %v\n", err)
	}


	authClient, err := app.Auth(context.Background())
	if err != nil {
		utils.Logger(
			fmt.Sprintf("Error initialising Firebase Auth: %v", err),
			"ERROR",
		)
		os.Exit(1)
	}

	return client, authClient


}

func Init() {
	app := InitFirebaseApp()
	FirebaseClient, FirebaseAuthClient = InitClient(app)

}

/*
LoadEnvironment loads configurations from the environment file.
*/
func LoadEnvironment() {
	err := godotenv.Load()
	if err != nil {
		utils.Logger(
			fmt.Sprintf("Error loading configurations from environment file: %v", err),
			"ERROR",
		)
		os.Exit(1)
	}
}
