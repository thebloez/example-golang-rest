package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func findUsers(w http.ResponseWriter)  {

}

func returnAllUsers(w http.ResponseWriter, r *http.Request)  {
	var users Users
	var arrUser []Users
	var response Response

	db := connect()
	defer db.Close()

	rows, err := db.Query("SELECT id, first_name, last_name from person")
	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		if err := rows.Scan(&users.ID, &users.FirstName, &users.LastName); err != nil {
			log.Fatal(err.Error())
		} else {
			arrUser = append(arrUser, users)
		}
	}

	response.Status = 1
	response.Message = "Success"
	response.Data = arrUser

	w.Header().Set("Content-Type","Application/json")
	json.NewEncoder(w).Encode(response)
}
