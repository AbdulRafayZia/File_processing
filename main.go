package main

import (
	"fmt"
	"net/http"

	"gorilla/api"


	
	"github.com/gorilla/mux"

)



// YourFunction is the function that processes the input and returns a result

// HandlePostRequest handles the POST request and calls YourFunction

func main() {
	
	r := mux.NewRouter()

	// Define the route for the POST request
	r.HandleFunc("/api/FileSummary", api.HandlePostRequest).Methods("POST")

	// Start the server
	port := 8080
	fmt.Printf("Server listening on :%d...\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), r)
}