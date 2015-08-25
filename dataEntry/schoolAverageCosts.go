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
	err = db.Pin()
	db.setMaxOpenConns(0)

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

		for _, rows := range(lines) {
			if i > 1 {
				line := strings.TrimSpace(row)
				splitStr := strings.Split(line, ",")
				
			}
		}
	}

}