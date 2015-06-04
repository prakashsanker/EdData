package main

import (
	"fmt"
	"database/sql"
    _ "github.com/go-sql-driver/mysql"
	"log"
	"os"
	"encoding/csv"
)

func main() {

db, err := sql.Open("mysql", "psanker:123@/education_data")
err = db.Ping()

if err != nil {
	fmt.Println("Failed to prepare connection to database")
	log.Fatal("Error:", err.Error())
}





defer db.Close()




	dat, err := os.Open("expenses.csv")
	check(err)

	defer dat.Close()

	reader := csv.NewReader(dat)

	reader.FieldsPerRecord = -1

	rawCSVData, err := reader.ReadAll()

	check(err)

	for _, each := range rawCSVData {
		fmt.Println(each)
	}




}

func check(e error) {
    if e != nil {
        panic(e)
    }
}


