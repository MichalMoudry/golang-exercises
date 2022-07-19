package main

import (
	"errors"
	"fmt"
)

// Record represents an expense record.
type Record struct {
	Day      int
	Amount   float64
	Category string
}

// DaysPeriod represents a period of days for expenses.
type DaysPeriod struct {
	From int
	To   int
}

// Filter returns the records for which the predicate function returns true.
func Filter(in []Record, predicate func(Record) bool) []Record {
	var res []Record
	for i := 0; i < len(in); i++ {
		if predicate(in[i]) {
			res = append(res, in[i])
		}
	}
	return res
}

// ByDaysPeriod returns predicate function that returns true when
// the day of the record is inside the period of day and false otherwise.
func ByDaysPeriod(p DaysPeriod) func(Record) bool {
	return func(r Record) bool {
		if r.Day >= p.From && r.Day <= p.To {
			return true
		} else {
			return false
		}
	}
}

// ByCategory returns predicate function that returns true when
// the category of the record is the same as the provided category
// and false otherwise.
func ByCategory(c string) func(Record) bool {
	return func(r Record) bool {
		if r.Category == c {
			return true
		} else {
			return false
		}
	}
}

// TotalByPeriod returns total amount of expenses for records
// inside the period p.
func TotalByPeriod(in []Record, p DaysPeriod) float64 {
	records := Filter(in, ByDaysPeriod(p))
	var res float64 = 0
	for i := 0; i < len(records); i++ {
		res += records[i].Amount
	}
	return res
}

// CategoryExpenses returns total amount of expenses for records
// in category c that are also inside the period p.
// An error must be returned only if there are no records in the list that belong
// to the given category, regardless of period of time.
func CategoryExpenses(in []Record, p DaysPeriod, c string) (float64, error) {
	records := Filter(in, ByCategory(c))
	if len(records) == 0 {
		return 0, errors.New(fmt.Sprintf("unknown category %s", c))
	}
	records = Filter(records, ByDaysPeriod(p))
	var res float64 = 0
	for i := 0; i < len(records); i++ {
		res += records[i].Amount
	}
	return res, nil
}

//-----------------------------------------
func Day1Records(r Record) bool {
	return r.Day == 1
}

func main() {
	records := []Record{
		{Day: 1, Amount: 15, Category: "groceries"},
		{Day: 11, Amount: 300, Category: "utility-bills"},
		{Day: 12, Amount: 28, Category: "groceries"},
	}
	filteredRecords := Filter(records, Day1Records)
	for i := 0; i < len(filteredRecords); i++ {
		println(filteredRecords[i].Day, "-", filteredRecords[i].Amount, "-", filteredRecords[i].Category)
	}

	println("------------------")

	records = []Record{
		{Day: 1, Amount: 15, Category: "groceries"},
		{Day: 11, Amount: 300, Category: "utility-bills"},
		{Day: 12, Amount: 28, Category: "groceries"},
		{Day: 26, Amount: 300, Category: "university"},
		{Day: 28, Amount: 1300, Category: "rent"},
	}
	period := DaysPeriod{From: 1, To: 15}
	filteredRecords = Filter(records, ByDaysPeriod(period))
	for i := 0; i < len(filteredRecords); i++ {
		println(filteredRecords[i].Day, "-", filteredRecords[i].Amount, "-", filteredRecords[i].Category)
	}

	println("------------------")

	filteredRecords = Filter(records, ByCategory("groceries"))
	for i := 0; i < len(filteredRecords); i++ {
		println(filteredRecords[i].Day, "-", filteredRecords[i].Amount, "-", filteredRecords[i].Category)
	}

	println("------------------")
	records = []Record{
		{Day: 15, Amount: 16, Category: "entertainment"},
		{Day: 32, Amount: 20, Category: "groceries"},
		{Day: 40, Amount: 30, Category: "entertainment"},
	}
	p1 := DaysPeriod{From: 1, To: 30}
	p2 := DaysPeriod{From: 31, To: 60}
	println("Total by period (expected - 16):", TotalByPeriod(records, p1))
	println("Total by period (expected - 50):", TotalByPeriod(records, p2))

	records = []Record{
		{Day: 1, Amount: 15, Category: "groceries"},
		{Day: 11, Amount: 300, Category: "utility-bills"},
		{Day: 12, Amount: 28, Category: "groceries"},
		{Day: 26, Amount: 300, Category: "university"},
		{Day: 28, Amount: 1300, Category: "rent"},
	}
	exp1, err1 := CategoryExpenses(records, p1, "entertainment")
	println("Category expenses (expected - 0, err):", exp1, err1.Error())
	// => 0, error(unknown category entertainment)
	exp2, err2 := CategoryExpenses(records, p1, "rent")
	println("Category expenses (expected - 1300, nil):", exp2, err2)
	// => 1300, nil
	exp3, err3 := CategoryExpenses(records, p2, "rent")
	println("Category expenses (expected - 0, nil):", exp3, err3)
	// => 0, nil
}
