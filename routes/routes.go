package routes

import (

	"github.com/AbdulRafayZia/Gorilla-mux/login"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq" 
)

func Routes() *mux.Router{
	r := mux.NewRouter()
	r.HandleFunc("/login", login.LoginHandler).Methods("POST")
	r.HandleFunc("/protected", login.ProtectedHandler).Methods("POST")
	r.HandleFunc("/signup", login.CreateUserHandler).Methods("POST")
	return r

}
