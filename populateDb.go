package main

import (
	"fmt"
	"database/sql"
    _ "github.com/go-sql-driver/mysql"
	"log"
	"io/ioutil"
	"strings"
	"unicode"
)

func main() {

db, err := sql.Open("mysql", "psanker:123@/education_data")
err = db.Ping()

if err != nil {
	fmt.Println("Failed to prepare connection to database")
	log.Fatal("Error:", err.Error())
}

defer db.Close()


	content, err := ioutil.ReadFile("activities.csv")

	lines := strings.Split(string(content), "\r")

	//only work so long as I have one district
	rows, err := db.Query("SELECT id FROM districts")
	var districtId int
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&districtId)
		check(err)
	}


	var activityCode string
	for _ , line := range lines {
			line = strings.TrimSpace(line)
			splitStr := strings.Split(line, ",")
			strings.Replace(splitStr[1], "\"", "", -1)
			if isUpper(splitStr[1]) && !strings.Contains("Subtotal", splitStr[2]){
				//this is an activity
				activityCode = splitStr[0]
				_, err := db.Exec("INSERT INTO activities(code, name) VALUES(?, ?)", splitStr[0], splitStr[1])
				check(err)
			} else {
				//this is a sub activity
				check(err)
				if !strings.Contains(splitStr[0], "Subtotal") && !strings.Contains(splitStr[0], "Total Expenditures") {
					if activityCode != "" {
						rows, _ := db.Query("SELECT id from activities where code = ?", activityCode)
						var activityId int
						for rows.Next() {
							err = rows.Scan(&activityId)
							check(err)
							_, err = db.Exec("INSERT INTO sub_activities(activity_id, code, name) VALUES(?,?,?)", activityId, splitStr[0], splitStr[1])
							check(err)
						}
					}
				}

				rows, _ := db.Query("SELECT id from sub_activities where code= ?", splitStr[0])
				var sub_activityId int
				for rows.Next() {
					err := rows.Scan(&sub_activityId)
					check(err)
					if splitStr[len(splitStr) - 1] != "N/A" {
						_, err = db.Exec("INSERT INTO sub_activity_expenditure(district_id, sub_activity_id, expenditure) VALUES(?,?,?)", districtId, sub_activityId, splitStr[len(splitStr) - 1])
						check(err)	
					}
				}
			}
	}

	content, err = ioutil.ReadFile("expenditureTypes.csv")
	lines = strings.Split(string(content), "\r")

	for i, line := range lines {
		var expenseId int64
		if i > 1 {
			splitStr := strings.Split(line, ",")
			if splitStr[2] != "" {
				var existingId int64
				existingId = -1
				rows, _ := db.Query("SELECT id FROM expenditure_types WHERE EXISTS (SELECT * FROM expenditure_types where code=?)", splitStr[3])
				for rows.Next() {
					err := rows.Scan(&existingId)
					check(err)
				}
				if existingId == -1 {
					res, err := db.Exec("INSERT INTO expenditure_types(name, code) VALUES(?, ?) ", splitStr[2], splitStr[3])
					expenseId1, e := res.LastInsertId()
					check(e)
					check(err)
					expenseId = expenseId1
				} else {
					expenseId = existingId
				}

			}

			rows, _ := db.Query("SELECT id from sub_activities WHERE code=?", splitStr[0])
			var activityId int
			for rows.Next() {
				err := rows.Scan(&activityId)
				fmt.Println(splitStr[4])
				check(err)
				_, err = db.Exec("INSERT INTO activity_expenditure_types(district_id, sub_activity_id, expense_id, restricted, unrestricted) VALUES(?,?,?,?,?)", districtId, activityId,expenseId , splitStr[4], splitStr[5])
				check(err)
			}
		}
	}



}

func isUpper(str string) bool {
	for _, char := range str {
		if char != '-' && char != ' ' && char != '"' {
			if (!unicode.IsUpper(char)) {
				return false
			}	
		}
	}
	return true
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}


