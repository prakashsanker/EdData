package main

import (
	"fmt"
	"net/http"
    "github.com/gorilla/mux"
    "encoding/json"
)




func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "WELCOME!")
}

func getDistrict(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	districtId := vars["districtId"]
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	rows, err := db.Query("SELECT * from districts WHERE id=?", districtId)
	check(err)
	var district District
	for rows.Next() {
		var name string
		var id int64
		err := rows.Scan(&id, &name)
		check(err)
		district = District{Id: id, Name: name}
	}
	if err := json.NewEncoder(w).Encode(district); err != nil {
		check(err)
	}
}

func getDistricts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	rows, err := db.Query("SELECT * from districts")
	check(err)
	var district District
	var districts Districts
	for rows.Next() {
		var id int64
		var name string
		err := rows.Scan(&id, &name)
		check(err)
		district = District{Id: id, Name: name}
		districts = append(districts, district)
	}
	if err := json.NewEncoder(w).Encode(districts); err != nil {
		check(err)
	}
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}