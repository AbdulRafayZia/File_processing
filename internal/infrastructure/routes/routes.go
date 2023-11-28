package routes

import (
	"github.com/AbdulRafayZia/Gorilla-mux/internal/app/api"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func Routes() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/user/login", api.LoginHandler).Methods("POST")
	r.HandleFunc("/fileProcess", api.ProcessFile).Methods("POST")
	r.HandleFunc("/user/signup", api.CreateUserHandler).Methods("POST")
	r.HandleFunc("/staff/staffLogin", api.StaffLogin).Methods("POST")
	r.HandleFunc("/user/user_processes", api.GetUsersProcessses).Methods("GET")
	r.HandleFunc("/user/get_process/{id}", api.GetProcessById).Methods("GET")
	r.HandleFunc("/staff/statistics", api.Statistics).Methods("POST")
	r.HandleFunc("/staff/get_all_processes", api.GetAllProcesses).Methods("GET")

	return r

}
