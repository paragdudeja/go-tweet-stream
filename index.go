package main

import (
	"fmt"
	"go-tweet-stream/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// Main function
func main() {

	// Init the mux router
	router := mux.NewRouter()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router.HandleFunc("/tweets", handlers.HandleTweets).Methods("GET")

	// serve the app
	fmt.Println("Server at 8101")
	log.Fatal(http.ListenAndServe(":8101", router))
}
