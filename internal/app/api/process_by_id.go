package api

import (
	"encoding/json"
	"log"
	"net/http"
	"github.com/AbdulRafayZia/Gorilla-mux/internal/app/service"

	"github.com/AbdulRafayZia/Gorilla-mux/internal/app/utils"
	database "github.com/AbdulRafayZia/Gorilla-mux/internal/infrastructure/Database"
	"github.com/gorilla/mux"
)

func GetProcessById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	w.Header().Set("Content-Type", "application/json")
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
	db:=database.OpenDB()
	defer db.Close()
	rows, err := db.Query("SELECT * FROM file_processing_data WHERE username = $1 AND id=$2", Claims.Username, id)
	if err != nil {
		http.Error(w, " Cannot find Process of this user with provided id ", http.StatusBadRequest)
		return
	}
	defer rows.Close()
	Record := make([]utils.ResponseBody, 0)
	for rows.Next() {
		var response utils.ResponseBody
		var ProcessId int

		err := rows.Scan(&ProcessId, &response.Filename, &response.TotalWords, &response.TotalLines, &response.TotalPuncuations, &response.TotalVowels, &response.ExecutionTime, &response.Routines, &response.Username)
		if err != nil {
			log.Fatal(err)
		}
		Record = append(Record, response)

	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Record)

}
