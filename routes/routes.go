package routes

import (

	"github.com/AbdulRafayZia/Gorilla-mux/handler"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq" 
)

func Routes() *mux.Router{
	r := mux.NewRouter()
	r.HandleFunc("/login", handler.LoginHandler).Methods("POST")
	r.HandleFunc("/protected", handler.ProtectedHandler).Methods("POST")
	r.HandleFunc("/signup", handler.CreateUserHandler).Methods("POST")
	return r

}
