package main 

import (
	"net/http"
	"fmt"
    _ "github.com/go-sql-driver/mysql"
    "database/sql"
)


func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func schoolHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, r.URL.Path)
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/school/", schoolHandler)
	http.ListenAndServe(":8080", nil)

	_, err := sql.Open("mysql", "psanker:psanker@/test?charset=utf8")
	checkErr(err)
	

}

func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}