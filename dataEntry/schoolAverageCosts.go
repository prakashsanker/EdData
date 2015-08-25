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
		log.Fatal("Error: ", err.Error())

		defer db.Close()

		content, err := ioutil.ReadFile("currentExpense1314.csv")

		lines := strings.Split(string(content), "\r")

		rows, err := db.Query("CREATE TABLE IF NOT EXISTS school_expenses(id INT(11) NOT NULL AUTO_INCREMENT, school_id INT(11) NOT NULL, expenditure INT(25), current_expense_ada INT(25), current_expense_per_ada INT(25), PRIMARY KEY(id), FOREIGN KEY(school_id) REFERENCES schools(id))")
		check(err)
		defer rows.Close()

		for i, row := range(lines) {
			if i > 0 {
				line := strings.TrimSpace(row)
				splitStr := strings.Split(line, ",")
				schoolName := splitStr[3]
				schoolRow, err := db.Query("SELECT id from schools where name=?", schoolName)
				check(err)
				hasNextRow := schoolRow.Next()
				if hasNextRow {
					//is a valid school
					var id int
					err = schoolRow.Scan(&id)
					_, err = db.Exec("INSERT INTO school_expenses(school_id, expenditure, current_expense_ada, current_expense_per_ada) VALUES(?,?,?,?)", id, splitStr[4], splitStr[5], splitStr[6])
					check(err)
				}
				schoolRow.Close()
			}
		}
	}
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}