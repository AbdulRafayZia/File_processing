package main

import (
	"fmt"
	"net/http"

	"github.com/AbdulRafayZia/Gorilla-mux/internal/infrastructure/routes"
	_ "github.com/lib/pq"
)

func main() {
	// defer login.db.Close()

	r:=routes.Routes()

	// Start the server
	port := 8080
	fmt.Printf("Server listening on :%d...\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), r)
}
