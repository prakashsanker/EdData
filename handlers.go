package main

import (
	"fmt"
	"net/http"
    "github.com/gorilla/mux"
    "encoding/json"
    "strconv"
	"database/sql"
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
		var code string
		err := rows.Scan(&id, &name, &code)
		check(err)
		district = District{Id: id, Name: name}
	}
	if err := json.NewEncoder(w).Encode(district); err != nil {
		check(err)
	}
}

func getDistricts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/jsonp;charset=UTF-8")
    w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)	
	rows, err := db.Query("SELECT * from districts")
	check(err)
	var district District
	var districts Districts
	for rows.Next() {
		var id int64
		var name string
		var code string
		err := rows.Scan(&id, &name, &code)
		check(err)
		district = District{Id: id, Name: name, Code: code}
		districts = append(districts, district)
	}
	var expenditure string
	var currentExpenseAda string
	var currentExpensePerAda string
	for index, _ := range districts {
		var id = districts[index].Id
		err := db.QueryRow("SELECT expenditure, current_expense_ada, current_expense_per_ada FROM district_expenses where district_id=?", id).Scan(&expenditure, &currentExpenseAda, &currentExpensePerAda)
		if (err == nil) {
			districts[index].Expenditure = expenditure
			districts[index].CurrentExpenseADA = currentExpenseAda
			districts[index].CurrentExpensePerAda = currentExpensePerAda
		} else if (err != sql.ErrNoRows) {
			check(err)
		}
	}
	if err := json.NewEncoder(w).Encode(districts); err != nil {
		check(err)
	}
}

func getDistrictDemography(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	districtId := vars["districtId"]
	w.Header().Set("Content-Type", "application/jsonp;charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Access-Control-Allow-Origin","*")
	districtRows, err := db.Query("SELECT name from districts WHERE id=?", districtId)
	check(err)
	var districtName string
	for districtRows.Next() {
		err = districtRows.Scan(&districtName)
		check(err)
	}
	districtRows.Close()
	districtSchoolRows, err := db.Query("SELECT school_id from districts_schools_mapping WHERE district_id=?", districtId)
	check(err)
	var district District
	var schools Schools
	for districtSchoolRows.Next() {
		//for each school take the id and name
		var schoolId int64
		var schoolName string
		err = districtSchoolRows.Scan(&schoolId)
		check(err)
		err := db.QueryRow("SELECT name FROM schools where id=?", schoolId).Scan(&schoolName)
		check(err)
		//find the demographics
		demographicRows, err := db.Query("SELECT * from demographics where school_id=?", schoolId)
		check(err)
		var ethnicBreakdowns EthnicBreakdowns
		var school School
		for demographicRows.Next() {
				var id int64
				var schoolId int64
				var ethnicity string
				var gender string
				var kindergarten int64
				var grade1 int64
				var grade2 int64
				var grade3 int64
				var grade4 int64
				var grade5 int64
				var grade6 int64
				var grade7 int64
				var grade8 int64
				var grade9 int64
				var grade10 int64
				var grade11 int64
				var grade12 int64
				var ungradedElementary int64
				var ungradedSecondary int64
				var total int64
				var adult int64
				err := demographicRows.Scan(&id, &schoolId, &ethnicity, &gender, &kindergarten, &grade1, &grade2, &grade3, &grade4, &grade5, &grade6, &grade7, &grade8, &grade9, &grade10, &grade11, &grade12, &ungradedElementary, &ungradedSecondary, &total, &adult)
				check(err)
				schoolEthnicityBreakdown := EthnicBreakdown{Ethnicity: ethnicity, Gender: gender, Kindergarten: kindergarten, Grade1: grade1, Grade2: grade2, Grade3: grade3, Grade4: grade4, Grade5: grade5, Grade6: grade6, Grade7: grade7, Grade8: grade8, Grade9: grade9, Grade10: grade10, Grade11: grade11, Grade12: grade12, UngradedElementary: ungradedElementary, UngradedSecondary: ungradedSecondary, Total: total, Adult: adult}
				ethnicBreakdowns = append(ethnicBreakdowns, schoolEthnicityBreakdown)
			}
		school = School{Id: schoolId, Name: schoolName, EthnicInfo: ethnicBreakdowns}
		schools = append(schools, school)
		demographicRows.Close()
	}
	districtSchoolRows.Close()
	districIdNum, err := strconv.ParseInt(districtId, 10, 64)
	check(err)
	district = District{Id: districIdNum, Name: districtName, Schools: schools}
	if err := json.NewEncoder(w).Encode(district); err != nil {
		check(err)
	}
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}