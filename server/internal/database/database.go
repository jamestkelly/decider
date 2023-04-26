package database

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

var (
	FirebaseApp *firebase.App
	FirebaseDB  *firebase.Database
)

// InitializeFirebase initializes Firebase Realtime Database
func InitializeFirebase() {
	ctx := context.Background()
	opt := option.WithCredentialsFile("path/to/firebase/service/account/key.json") // TODO: Replace with correct creds config

	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Fatalf("error initializing Firebase app: %v", err)
	}

	db, err := app.Database(ctx)
	if err != nil {
		log.Fatalf("error initializing Firebase database: %v", err)
	}

	FirebaseApp = app
	FirebaseDB = db
}
