package main

type District struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
	Schools Schools `json:"schools"`
	Code string `json:"code"`
	Expenditure string `json:"expenditure"`
	CurrentExpenseADA string `json:"currentExpenseAda"`
	CurrentExpensePerAda string `json:"currentExpensePerAda"`
}

type Districts []District

type Activity struct {
	Id int64 `json:"id"`
	Code string `json:"code"`
	Expenditure string `json:"expenditure" `
	Name string `json:"name"`
}

type SubActivity struct {
	Id int64 `json:"id"`
	Code string `json:"code"`
	Expenditure string `json:"expenditure"`
	Name string `json:"name"`
}

type Expense struct {
	Id int64 `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
	RestrictedExpenditure string `json:"restricted"`
	UnrestrictedExpenditure string `json:"unrestricted"`
}

type School struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
	EthnicInfo EthnicBreakdowns `json:"ethnicInfo"`
}

type EthnicBreakdown struct {
	Ethnicity string `json:"ethnicity"`
	Gender string `json:"ethnicity"`
	Kindergarten int64 `json:"kindergarten"`
	Grade1 int64 `json:"grade1"`
	Grade2 int64 `json:"grade2"`
	Grade3 int64 `json:"grade3"`
	Grade4 int64 `json:"grade4"`
	Grade5 int64 `json:"grade5"`
	Grade6 int64 `json:"grade6"`
	Grade7 int64 `json:"grade7"`
	Grade8 int64 `json:"grade8"`
	Grade9 int64 `json:"grade9"`
	Grade10 int64 `json:"grade10"`
	Grade11 int64 `json:"grade11"`
	Grade12 int64 `json:"grade12"`
	UngradedElementary int64 `json:"ungrElem"`
	UngradedSecondary int64 `json:"ungrSec"`
	Total int64 `json:"total"`
	Adult int64 `json:"adult"`
}

type Activities []Activity
type SubActivities []SubActivity
type Expenses []Expense
type EthnicBreakdowns []EthnicBreakdown
type Schools []School
