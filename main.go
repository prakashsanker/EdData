package main 

import (
	"net/http"
	"log"
	
)



func main() {
	router := NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))

}

func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}