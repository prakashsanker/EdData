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
	fmt.Println("WHAT")

	if err != nil {
		fmt.Println("HELLO")
		fmt.Println("Failed to prepare connection to database")
		log.Fatal("Error: ", err.Error())
	}

		defer db.Close()

		content, err := ioutil.ReadFile("currentExpense1314.csv")
		check(err)

		lines := strings.Split(string(content), "\r")

		rows, err := db.Query("CREATE TABLE IF NOT EXISTS district_expenses(id INT(11) NOT NULL AUTO_INCREMENT, district_id INT(11) NOT NULL, expenditure INT(25), current_expense_ada INT(25), current_expense_per_ada INT(25), PRIMARY KEY(id), FOREIGN KEY(district_id) REFERENCES districts(id))")
		check(err)
		defer rows.Close()

		for i, row := range(lines) {
			if i > 0 {
				line := strings.TrimSpace(row)
				splitStr := strings.Split(line, ",")
				districtName := strings.Replace(splitStr[2], "\"", "", -1)
				test := "SELECT id from districts where name LIKE \"%" + districtName + "%\""
				districtRow, err := db.Query(test)
				check(err)
				hasNextRow := districtRow.Next()
				if hasNextRow {
					//is a valid district
					var id int
					err = districtRow.Scan(&id)
					length := len(splitStr)
					startingIndex := length - 4
					expenditure := strings.Replace(splitStr[startingIndex], " \" ", "", -1)
					currentExpenseAda := strings.Replace(splitStr[startingIndex + 1], "\"", "", -1)
					currentExpensePerAda := strings.Replace(splitStr[startingIndex + 2], "\"", "", -1)
					expenditure = strings.Replace(expenditure, "\"", "", -1)
					_, err = db.Exec("INSERT INTO district_expenses(district_id, expenditure, current_expense_ada, current_expense_per_ada) VALUES(?,?,?,?)", id, expenditure, currentExpenseAda, currentExpensePerAda)
					check(err)
				} else {
					fmt.Print("no district")
					fmt.Println("district name : " + districtName)
				}
				districtRow.Close()
			}
		}
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}