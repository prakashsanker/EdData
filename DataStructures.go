package main

type District struct {
	Id int64 `json: "id"`
	Activities string `json: "activities"`
}

type Districts []District

type Activity struct {
	Id int64 `json: "id"`
	Code string `json: "code"`
	Expenditure float64 `json: "expenditure" `
}

type Activities []Activity