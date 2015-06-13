package main

type District struct {
	Id int64 `json: "id"`
	Name string `json: "name"`
}

type Districts []District

type Activity struct {
	Id int64 `json: "id"`
	Code string `json: "code"`
	Expenditure string `json: "expenditure" `
	Name string `json : "name"`
}

type SubActivity struct {
	Id int64 `json: "id"`
	Code string `json: "code"`
	Expenditure string `json: "expenditure"`
	Name string `json: "name"`
}

type Expense struct {


}

type Activities []Activity
type SubActivities []SubActivity
