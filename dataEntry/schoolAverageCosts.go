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

		rows, err := db.Query("CREATE TABLE IF NOT EXISTS district_expenses(id INT(11) NOT NULL AUTO_INCREMENT, district_id INT(11) NOT NULL, expenditure INT(25), current_expense_ada INT(25), current_expense_per_ada INT(25), PRIMARY KEY(id), FOREIGN KEY(school_id) REFERENCES districts(id))")
		check(err)
		stmt, err := db.Prepare("CREATE TABLE IF NOT EXISTS districts (id INT(11) NOT NULL AUTO_INCREMENT, district_id INT(11) NOT NULL, PRIMARY KEY (id));")
		check(err)
		stmt.Exec()
		defer rows.Close()

		for i, row := range(lines) {
			if i > 0 {
				line := strings.TrimSpace(row)
				splitStr := strings.Split(line, ",")
				schoolName := splitStr[2]
				test := "SELECT id from schools where name LIKE \"%" + schoolName + "%\""
				fmt.Println(line)
				fmt.Println(splitStr)
				schoolRow, err := db.Query(test)
				check(err)
				hasNextRow := schoolRow.Next()
				if hasNextRow {
					//is a valid school
					var id int
					err = schoolRow.Scan(&id)
					length := len(splitStr)
					startingIndex := length - 4
					expenditure := strings.Replace(splitStr[startingIndex], " \" ", "", -1)
					currentExpenseAda := strings.Replace(splitStr[startingIndex + 1], "\"", "", -1)
					currentExpensePerAda := strings.Replace(splitStr[startingIndex + 2], "\"", "", -1)
					expenditure = strings.Replace(expenditure, "\"", "", -1)
					_, err = db.Exec("INSERT INTO district_expenses(school_id, expenditure, current_expense_ada, current_expense_per_ada) VALUES(?,?,?,?)", id, expenditure, currentExpenseAda, currentExpensePerAda)
					check(err)
				} else {
					fmt.Print("no school")
					fmt.Println("schoolName : " + schoolName)
				}
				schoolRow.Close()
			}
		}
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}