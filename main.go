package main 

import (
	"net/http"
	"log"
	"github.com/rs/cors"
)



func main() {
	router := NewRouter()
	handler := cors.Default().Handler(router)

	log.Fatal(http.ListenAndServe(":8100", handler))

}

func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}