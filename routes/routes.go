package routes

import (

	"github.com/AbdulRafayZia/Gorilla-mux/api"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq" 
)

func Routes() *mux.Router{
	r := mux.NewRouter()
	r.HandleFunc("/login", api.LoginHandler).Methods("POST")
	r.HandleFunc("/protected", api.ProtectedHandler).Methods("POST")
	r.HandleFunc("/signup", api.CreateUserHandler).Methods("POST")
	r.HandleFunc("/staffLogin", api.StaffLogin).Methods("POST")

	return r

}
