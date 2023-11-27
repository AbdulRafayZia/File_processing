package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/AbdulRafayZia/Gorilla-mux/internal/app/service"
	"github.com/AbdulRafayZia/Gorilla-mux/internal/app/utils"
	database "github.com/AbdulRafayZia/Gorilla-mux/internal/infrastructure/Database"
)

func GetUsersProcessses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Missing authorization header")
		return
	}
	tokenString = tokenString[len("Bearer"):]

	Claims, err := service.VerifyToken(tokenString)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Could not Get Claims")
		return
	}
	db:=database.OpenDB()
	defer db.Close()
	rows, err := db.Query("SELECT * FROM file_processing_data WHERE username = $1", Claims.Username)
	if err != nil {
		log.Fatal(err)
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
