package main

import(
	"fmt"
	"database/sql"
    _ "github.com/go-sql-driver/mysql"
	"log"
	"io/ioutil"
	"strings"
	// "unicode"
)

func main() {
	db, err := sql.Open("mysql", "psanker:123@/education_data")
	err = db.Ping()

	db.SetMaxOpenConns(0)
	check(err)
	if err != nil {
		fmt.Println("Failed to prepare connection to database")
		log.Fatal("Error:", err.Error())
	}

	defer db.Close()

	content, err := ioutil.ReadFile("schoolEthnicity2014.csv")
	lines := strings.Split(string(content), "\r")
	//populate schools..
	rows,err := db.Query("CREATE TABLE IF NOT EXISTS schools ( id INT(11) NOT NULL AUTO_INCREMENT, name varchar(255) NOT NULL, PRIMARY KEY (id));")
	check(err)
	stmt, err := db.Prepare("CREATE TABLE IF NOT EXISTS demographics (id INT(11) NOT NULL AUTO_INCREMENT, school_id INT(11) NOT NULL, ethnicity ENUM('1','2','3','4','5','6','7','8','9'), gender ENUM('F', 'M'), kindergarten INT(11), grade_1 INT(11), grade_2 INT(11), grade_3 INT(11), grade_4 INT(11), grade_5 INT(11), grade_6 INT(11), grade_7 INT(11), grade_8 INT(11), grade_9 INT(11), grade_10 INT(11), grade_11 INT(11), grade_12 INT(11), ungr_elem INT(11), ungr_sec INT(11), total INT(11), adult INT(11), FOREIGN KEY(school_id) REFERENCES schools(id), PRIMARY KEY(id))")
	check(err)
	stmt.Exec()
	stmt, err = db.Prepare("CREATE TABLE IF NOT EXISTS districts (id INT(11) NOT NULL AUTO_INCREMENT, code VARCHAR(25), PRIMARY KEY (id));")
	check(err)
	stmt.Exec()
	stmt, err = db.Prepare("CREATE TABLE IF NOT EXISTS districts_schools_mapping (id INT(11) NOT NULL AUTO_INCREMENT, district_id INT(11) NOT NULL, school_id INT(11) NOT NULL, PRIMARY KEY (id), FOREIGN KEY(school_id) REFERENCES schools(id));")
	check(err)
	stmt.Exec()
	defer rows.Close()

	for i, row := range(lines) {
		if i > 1 {
			line := strings.TrimSpace(row)
			fmt.Println(line)
			splitStr := strings.Split(line, ",")
			if splitStr[3] != "" {
				schoolRows, err := db.Query("SELECT id from schools where name=?", splitStr[3])
				check(err)
				// var schoolId sql.Result
				hasNextRow := schoolRows.Next()
				if !hasNextRow {
					_, _ = db.Exec("INSERT INTO schools(name) VALUES (?)", splitStr[3])
				} else {
					//has a school row
					var id int
					err = schoolRows.Scan(&id)
					check(err)
					schoolRows.Close()
					splitStrLen := len(splitStr)
					startingIndex := splitStrLen - 19
					demographicRow, err := db.Query("SELECT id from demographics where id=?", id)
					hasNextRow = demographicRow.Next()
					demographicRow.Close()
					if !hasNextRow {
						_, err := db.Exec("INSERT INTO demographics(school_id, ethnicity, gender, kindergarten, grade_1, grade_2, grade_3, grade_4, grade_5, grade_6, grade_7, grade_8, grade_9, grade_10, grade_11, grade_12, ungr_elem, ungr_sec, total, adult) VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?, ?, ?)", id, splitStr[startingIndex], splitStr[startingIndex + 1], splitStr[startingIndex + 2],splitStr[startingIndex + 3], splitStr[startingIndex + 4], splitStr[startingIndex + 5], splitStr[startingIndex + 6], splitStr[startingIndex + 7],splitStr[startingIndex + 8], splitStr[startingIndex + 9], splitStr[startingIndex + 10], splitStr[startingIndex + 11], splitStr[startingIndex + 12], splitStr[startingIndex + 13], splitStr[startingIndex + 14], splitStr[startingIndex + 15], splitStr[startingIndex + 16], splitStr[startingIndex + 17], splitStr[startingIndex + 18])
						check(err)
					}
					districtRow, err := db.Query("SELECT id from districts where name=?", strings.TrimSpace(splitStr[2]));
					check(err)
					hasNextRow = districtRow.Next()
					districtRow.Close()
					if !hasNextRow {
						code := splitStr[0]
						code = code[2:6]
						districtResult, err := db.Exec("INSERT INTO districts(name, code) VALUES(?, ?)", strings.TrimSpace(splitStr[2]), code)
						check(err)
						districtId, err := districtResult.LastInsertId()
						check(err)
						districtSchoolRow, err := db.Query("SELECT id from districts_schools_mapping WHERE district_id=? AND school_id=?", districtId, id)
						hasNextRow = districtSchoolRow.Next()
						districtSchoolRow.Close()
						if !hasNextRow {
							_, err = db.Exec("INSERT INTO districts_schools_mapping(district_id, school_id) VALUES(?,?)", districtId, id)
							check(err)
						}
					}
				}
				schoolRows.Close()
			}
		}
	}
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}