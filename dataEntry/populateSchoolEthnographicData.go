package main

import (
	"fmt"
    csvSource "gopkg.in/Clever/optimus.v3/sources/csv"
    "gopkg.in/Clever/optimus.v3"
	Transformer "gopkg.in/Clever/optimus.v3/transformer"
	Transforms "gopkg.in/Clever/optimus.v3/transforms"
    "os"
    "errors"
)

func main() {
	f, err := os.Open("activities.csv")
	check(err)
	csvTable := csvSource.New(f)
	fmt.Println(csvTable)
	transformer := Transformer.New(csvTable)
	eachTransformFunc := Transforms.Each(printRow)
	transformer.Apply(eachTransformFunc)
	rows := csvTable.Rows()
	fmt.Println(rows)
}

func printRow(row optimus.Row) error {
	fmt.Println("NO")
	return errors.New("Each not working")
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}