package main

import (
	"fmt"
	"net/http"
	// "github.com/AbdulRafayZia/Gorilla-mux/api"
	"github.com/gorilla/mux"
	"github.com/AbdulRafayZia/Gorilla-mux/login"

)




func main() {
	
	r := mux.NewRouter()

	// Define the route for the POST request
	// r.HandleFunc("/api/FileSummary", api.HandlePostRequest).Methods("POST")
	r.HandleFunc("/login", login.LoginHandler).Methods("GET")
	r.HandleFunc("/protected", login.ProtectedHandler).Methods("POST")
	// Start the server
	port := 8080
	fmt.Printf("Server listening on :%d...\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), r)
}
