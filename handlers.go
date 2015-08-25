package main

import (
	"fmt"
	"net/http"
    "github.com/gorilla/mux"
    "encoding/json"
)

//this file really really really really really needs refactoring. 


func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "WELCOME!")
}

func getExpenses(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	districtId := vars["districtId"]
	subActivityCode := vars["subActivityCode"]
	w.Header().Set("Access-Control-Allow-Origin", "*")	
	w.WriteHeader(http.StatusOK)
	rows, err := db.Query("SELECT id from sub_activities where code=?", subActivityCode)
	check(err)
	var expense Expense
	var expenses Expenses
	for rows.Next() {
		var subActivityId int64
		err := rows.Scan(&subActivityId)
		check(err)
		subActivityRows, err := db.Query("SELECT expense_id, restricted, unrestricted from activity_expenditure_types where district_id=? AND  sub_activity_id=?", districtId, subActivityId)
		check(err)
		for subActivityRows.Next() {
			var expenseId int64
			var restricted string
			var unrestricted string
			err := subActivityRows.Scan(&expenseId, &restricted, &unrestricted)
			check(err)
			expenseRows, err := db.Query("SELECT name, code FROM expenditure_types where id=?", expenseId)
			check(err)
			for expenseRows.Next() {
				var name string
				var code string
				err := expenseRows.Scan(&name, &code)
				check(err)
				expense = Expense{Id: expenseId, RestrictedExpenditure: restricted, UnrestrictedExpenditure: unrestricted, Name: name, Code: code}
				expenses = append(expenses, expense)
			}
		}

	}

	if err := json.NewEncoder(w).Encode(expenses); err != nil {
		check(err)
	}


}

func getSubActivity(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	districtId := vars["districtId"]
	subActivityCode := vars["subActivityCode"]
	w.Header().Set("Access-Control-Allow-Origin", "*")	
	w.WriteHeader(http.StatusOK)
	rows, err := db.Query("SELECT id, name from sub_activities where code=?", subActivityCode)
	check(err)
	var subActivity SubActivity
	for rows.Next() {
		var id int64
		var name string
		err := rows.Scan(&id, &name)
		check(err)
		subActivityRows, err := db.Query("SELECT expenditure FROM sub_activity_expenditure where district_id=? AND sub_activity_id=?", districtId, id)
		check(err)
		for subActivityRows.Next() {
			var expenditure string
			err := subActivityRows.Scan(&expenditure)
			check(err)
			subActivity = SubActivity{Id: id, Name: name, Expenditure: expenditure, Code: subActivityCode}
		}
	}
	if err := json.NewEncoder(w).Encode(subActivity); err != nil {
		check(err)
	}

}

func getSubActivities(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	districtId := vars["districtId"]
	activityCode := vars["activityCode"]
	w.Header().Set("Access-Control-Allow-Origin", "*")	
	w.WriteHeader(http.StatusOK)
	activityRows, err := db.Query("SELECT id from activities where code=?", activityCode)
	var activityId int64
	for activityRows.Next() {
		err := activityRows.Scan(&activityId)
		check(err)
	}
	rows, err := db.Query("SELECT id, name, code from sub_activities where activity_id=?", activityId)
	check(err)
	var subActivity SubActivity
	var subActivities SubActivities
	for rows.Next() {
		var id int64
		var name string
		var code string
		err := rows.Scan(&id, &name, &code)
		check(err)
		subActivityExpenditureRows, err := db.Query("SELECT expenditure from sub_activity_expenditure WHERE district_id=? AND sub_activity_id=?", districtId, id)
		check(err)
		for subActivityExpenditureRows.Next() {
			var expenditure string
			err := subActivityExpenditureRows.Scan(&expenditure)
			check(err)
			subActivity = SubActivity{Id: id, Name: name, Code: code, Expenditure: expenditure}
			subActivities = append(subActivities, subActivity)
		}
	}

	if err := json.NewEncoder(w).Encode(subActivities); err != nil {
		check(err)
	}
}

func getDistrictActivity(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	activityCode := vars["activityCode"]
	districtId := vars["districtId"]
	w.Header().Set("Access-Control-Allow-Origin", "*")	
	w.WriteHeader(http.StatusOK)
	rows, err := db.Query("SELECT  id, name from activities where code=?", activityCode)
	check(err)
	var activity Activity;
	for rows.Next() {
		var id int64
		var name string
		err := rows.Scan(&id, &name)
		check(err)
		activityExpenditureRows, err := db.Query("SELECT expenditure from activity_expenditure where district_id=? AND activity_id=?", districtId, id)
		check(err)
		for activityExpenditureRows.Next() {
			var expenditure string
			err := activityExpenditureRows.Scan(&expenditure)
			check(err)
			activity = Activity{Id: id, Name: name, Code: activityCode, Expenditure: expenditure}
		}
	}
	if err := json.NewEncoder(w).Encode(activity); err != nil {
		check(err)
	}

}

func getDistrictActivities(w http.ResponseWriter, r *http.Request) {	
	vars := mux.Vars(r)
	districtId := vars["districtId"]
	w.Header().Set("Access-Control-Allow-Origin", "*")	
	w.WriteHeader(http.StatusOK)
	rows, err := db.Query("SELECT activity_id, expenditure from activity_expenditure where district_id=?", districtId)
	check(err)
	var activity Activity;
	var activities Activities;
	for rows.Next() {
		var id int64
		var expenditure string
		err := rows.Scan(&id, &expenditure)
		check(err)
		codeNameRows, err1 := db.Query("SELECT code, name from activities where id=?", id)
		check(err1)
		for codeNameRows.Next() {
			var code string
			var name string
			err := codeNameRows.Scan(&code, &name)
			check(err)
			activity = Activity{Id: id, Name: name, Code: code, Expenditure: expenditure}
			activities = append(activities, activity)
		}
	}
	if err := json.NewEncoder(w).Encode(activities); err != nil {
		check(err)
	}
}

func getDistrict(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	districtId := vars["districtId"]
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")

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
	w.Header().Set("Content-Type", "application/jsonp;charset=UTF-8")
	w.WriteHeader(http.StatusOK)
    w.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Println("right function")
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