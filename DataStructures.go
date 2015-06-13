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
	Id int64 `json: "id"`
	Code string `json: "code"`
	Name string `json: "name"`
	RestrictedExpenditure string `json: "restricted"`
	UnrestrictedExpenditure string `json: "unrestricted"`
}

type Activities []Activity
type SubActivities []SubActivity
type Expenses []Expense
