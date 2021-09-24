package main

import (
	"context"
	"log"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
)

func initializeApp() *firebase.App {
	app, err := firebase.NewApp(context.Background(), nil)
	if err != nil {
		log.Fatalf("error initializing app %v\n", err)
	}

	return app
}

func initializeAuth(app *firebase.App) *auth.Client {
	client, err := app.Auth(context.Background())
	if err != nil {
		log.Printf("Error initializing authentication client %v\n", err)
	}
	return client
}