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
	fmt.Fprintln(w, "District id : ", districtId)
}

func getDistricts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	rows, err := db.Query("SELECT id from districts")
	check(err)
	var district District
	for rows.Next() {
		var id int64
		err := rows.Scan(&id)
		check(err)
		test := "hey"
		district = District{Id: id, Activities: test}
	}

	if err := json.NewEncoder(w).Encode(district); err != nil {
		check(err)
	}


}

func check(e error) {
    if e != nil {
        panic(e)
    }
}