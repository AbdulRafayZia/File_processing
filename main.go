package main

import (
	"fmt"
	"net/http"

	"github.com/AbdulRafayZia/Gorilla-mux/internal/infrastructure/routes"
	"github.com/gorilla/handlers"
	_ "github.com/lib/pq"
)

func main() {

	r := routes.Routes()
	port := 8080
	fmt.Printf("Server listening on :%d...\n", port)

	// Start the server
	headersOk := handlers.AllowedHeaders([]string{"Content-Type", "Authorization"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})
	http.ListenAndServe(":8080", handlers.CORS(originsOk, headersOk, methodsOk)(r))

	http.ListenAndServe(fmt.Sprintf(":%d", port), r)
}
