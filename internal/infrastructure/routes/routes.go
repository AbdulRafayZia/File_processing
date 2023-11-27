package routes

import (

	"github.com/AbdulRafayZia/Gorilla-mux/internal/app/api"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq" 
)

func Routes() *mux.Router{
	r := mux.NewRouter()
	r.HandleFunc("/login", api.LoginHandler).Methods("POST")
	r.HandleFunc("/fileProcess", api.ProcessFile).Methods("POST")
	r.HandleFunc("/signup", api.CreateUserHandler).Methods("POST")
	r.HandleFunc("/staffLogin", api.StaffLogin).Methods("POST")
	r.HandleFunc("/user_processes", api.GetUsersProcessses).Methods("GET")
	r.HandleFunc("/get_process/{id}", api.GetProcessById).Methods("GET")
	r.HandleFunc("/statistics", api.Statistics).Methods("POST")





	return r

}
