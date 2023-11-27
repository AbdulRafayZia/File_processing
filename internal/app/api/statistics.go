package api

import (
	"encoding/json"
	"fmt"

	"net/http"

	"github.com/AbdulRafayZia/Gorilla-mux/internal/app/service"
	"github.com/AbdulRafayZia/Gorilla-mux/internal/app/utils"
	database "github.com/AbdulRafayZia/Gorilla-mux/internal/infrastructure/Database"
)

type StatsRequest struct {
	Filename string `json:"filename"`
}


func Statistics(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var request StatsRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Unable to get data", http.StatusBadRequest)
		return
	}
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		w.WriteHeader(http.StatusUnauthorized)
		http.Error(w, " Missing Authorization Bearer ", http.StatusUnauthorized)

		return
	}
	tokenString = tokenString[len("Bearer"):]

	Claims, err := service.VerifyToken(tokenString)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		http.Error(w, " Could not get Claims ", http.StatusUnauthorized)

		return
	}
	validrole :=service.CheckRole(Claims.Role)
	if !validrole {
		http.Error(w, "Not a Staff Member", http.StatusUnauthorized)
		return

	}

	db:=database.OpenDB()
	defer db.Close()

	rows, err := db.Query("SELECT AVG(execution_time) FROM file_processing_data WHERE filename = $1", request.Filename)

	if err != nil {
		http.Error(w, " Cannot find Process of this user with provided Filename ", http.StatusBadRequest)
		return
	}
	
	defer rows.Close()
	var avgExecutionTime utils.ExecutionData

	// Check for rows
	for rows.Next() {
		
		// Scan the average execution time value into the variable
		if err := rows.Scan(&avgExecutionTime.AveragTime); err != nil {

			http.Error(w, " Cannot Scan the Average time from Rows", http.StatusInternalServerError)
			fmt.Println(err)

			return
		}
	}
	if err := rows.Err(); err != nil {

		http.Error(w, "Error in iterating over the Row", http.StatusInternalServerError)
		return

	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(avgExecutionTime)
}
