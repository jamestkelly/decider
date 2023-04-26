package main

import (
	"log"
	"net/http"

	"github.com/yourusername/yourprojectname/pkg/api"
)

func main() {
	api := api.NewAPI() // Initialise API instance

	// start the server and listen for incoming requests
	log.Println("Starting server on port 8080")
	if err := http.ListenAndServe(":8080", api.Router); err != nil {
		log.Fatal(err)
	}
}
