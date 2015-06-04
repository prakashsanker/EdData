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
	}

	rows, err = db.Query("SELECT DATABASE();")
	var test string

	for rows.Next() {
		rows.Scan(&test)
		fmt.Println(test)
	}

	var activityCode string
	for i, line := range lines {
		if i > 1 {
			splitStr := strings.Split(line, ",")
			if strings.Contains(splitStr[0], "-") {
				//this is an activity
				activityCode = splitStr[0]
				_, err := db.Exec("INSERT INTO activities(code, name) VALUES(?, ?)", splitStr[0], splitStr[1])
				check(err)
			} else {
				//this is a sub activity
				check(err)
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
				rows, _ := db.Query("SELECT id from sub_activities where code= ?", splitStr[0])
				var sub_activityId int
				for rows.Next() {
					err := rows.Scan(&sub_activityId)
					check(err)
					fmt.Println("EXPENDITURE")
					fmt.Println(splitStr[len(splitStr) - 1])
					fmt.Println(splitStr)
					if splitStr[len(splitStr) - 1] != "N/A" {
						_, err = db.Exec("INSERT INTO sub_activity_expenditure(district_id, sub_activity_id, expenditure) VALUES(?,?,?)", districtId, sub_activityId, splitStr[len(splitStr) - 1])
						check(err)	
					}

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


