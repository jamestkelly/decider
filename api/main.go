package main

import (
	"decider/api/handlers"
	"decider/api/utils"
	"github.com/gorilla/mux"
	"net/http"
)

/*
main
@title Decider API Documentation
@version 1.0.0
@host {url}:5000
@BasePath /api
*/
func main() {
	router := mux.NewRouter()

	utils.Logger("Initialising routes for Decider API.", "INFO")
	api := router.PathPrefix("/api").Subrouter()
	utils.Logger("/api.", "INFO")
	api.HandleFunc("/register", handlers.RegisterUser).Methods("POST")
	utils.Logger("POST /register.", "INFO")
	api.HandleFunc("/validateToken", handlers.ValidateFirebaseToken).Methods("POST")
	utils.Logger("POST /validateToken.", "INFO")

	http.ListenAndServe(":5000", router)
}
