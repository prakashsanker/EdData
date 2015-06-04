package main

import (
	"fmt"
	"database/sql"
    _ "github.com/go-sql-driver/mysql"
	"log"
	"io/ioutil"
	"strings"
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
		log.Println(districtId)
	}

	for i, line := range lines {
		if i > 1 {
			splitStr := strings.Split(line, ",")
			var activityCode string
			if strings.Contains(splitStr[0], "-") {
				//this is an activity
				activityCode = splitStr[0]
				stmt1, _ := db.Prepare("INSERT INTO activities(code) VALUES(?)")
				stmt2, _ := db.Prepare("INSERT INTO activities(name) VALUES(?)")
				stmt1.Exec(splitStr[0])
				stmt2.Exec(splitStr[1])
			} else {
				//this is a sub activity
				stmt1, _ := db.Prepare("INSERT INTO sub_activities(code) VALUES(?)")
				stmt2, _ := db.Prepare("INSERT INTO sub_activities(name) VALUES(?)")
				stmt1.Exec(splitStr[0])
				stmt2.Exec(splitStr[1])
				if activityCode != "" {
					rows, _ := db.Query("SELECT id from activities where code = ?", activityCode)
					var activityId int
					for rows.Next() {
						err := rows.Scan(&activityId)
						check(err)
						stmt3, err := db.Prepare("INSERT INTO sub_activities(activity_id) VALUES(?)")
						stmt3.Exec(activityId)
					}
				}
				rows, _ := db.Query("SELECT id from sub_activities where code= ?", splitStr[0])
				var sub_activityId int
				for rows.Next() {
					err := rows.Scan(&sub_activityId)
					check(err)
					stmt5, _ := db.Prepare("INSERT INTO sub_activity_expenditure(district_id) VALUES(?)")
					stmt6, _ := db.Prepare("INSERT INTO sub_activity_expenditure(sub_activity_id) VALUES(?)")
					stmt7, _ := db.Prepare("INSERT INTO sub_activity_expenditure(expenditure) VALUES(?)")
					stmt5.Exec(districtId)
					stmt6.Exec(sub_activityId)
					stmt7.Exec(splitStr[2])
				}
			}
		}
	}




}

func check(e error) {
    if e != nil {
        panic(e)
    }
}


