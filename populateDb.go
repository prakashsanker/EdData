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


	content, err := ioutil.ReadFile("expenses4.csv")

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
			functionCode := splitStr[0]
			if strings.Contains(functionCode, "Subtotal") {

			} else {
				stmt, err := db.Prepare("INSERT INTO functions(code) VALUES(?)")
				check(err)
				stmt.Exec(functionCode)
			}
			stmt, err := db.Prepare("INSERT INTO functions(name) VALUES(?)")
			expenseType := splitStr[1]
			check(err)
			stmt.Exec(expenseType)
			check(err)
			expenditure := splitStr[2]
			stmt, _ := db.Prepare("INSERT INTO subtype_expenditure(district_id) VALUES(?)")
			stmt.Exec(districtId)
			stmt, _ := db.Prepare("INSERT INTO subtype_expenditure(expenditure) VALUES(?)")
			stmt.Exec(expenditure)
			stmt, _ := db.Prepare("INSERT INTO subtype_expenditure(subtype_id) VALUES(?)")
			stmt.Exec(i-2)



		}
	}




}

func check(e error) {
    if e != nil {
        panic(e)
    }
}


