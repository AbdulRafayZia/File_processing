package main

import (
	// "encoding/json"
	"fmt"
	// "log"
	"net/http"

	// "github.com/AbdulRafayZia/Gorilla-mux/api"
	// "database/sql"

	"github.com/AbdulRafayZia/Gorilla-mux/login"
	

	"github.com/gorilla/mux"
	_ "github.com/lib/pq" 
)





func main() {
	// defer login.db.Close()

	r := mux.NewRouter()
	r.HandleFunc("/login", login.LoginHandler).Methods("POST")
	r.HandleFunc("/protected", login.ProtectedHandler).Methods("POST")
	r.HandleFunc("/signup",login.CreateUserHandler).Methods("POST")

	// Start the server
	port := 8080
	fmt.Printf("Server listening on :%d...\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), r)
}

