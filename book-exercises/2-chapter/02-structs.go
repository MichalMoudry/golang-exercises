package main

import "fmt"

type Account struct {
	id      int
	owner   Person
	balance int32
}

type Person struct {
	id   int
	name string
}

func (a Account) toString() string {
	return fmt.Sprintf("Account n. %d, owner name: %s", a.id, a.owner.name)
}

type Item struct {
	Title        string
	Description  string
	Quantity     int
	PricePerUnit float32
}

func (i Item) total() int {
	return i.Quantity * int(i.PricePerUnit)
}

func main() {
	var testAccount = Account{
		id: 0,
		owner: Person{
			id:   0,
			name: "Michal",
		},
		balance: 20,
	}
	println(testAccount.toString())
	println(testAccount.id, testAccount.owner.name)

	item1 := Item{
		Title:        "LEGO set",
		Description:  "4000 pieces",
		Quantity:     1,
		PricePerUnit: 600,
	}
	item2 := Item{
		Title:        "Plushy",
		Description:  "plush toy",
		Quantity:     3,
		PricePerUnit: 5,
	}
	var items = []Item{item1, item2}
	var sum int
	for i := 0; i < len(items); i++ {
		println("Item n.", i, "-", items[i].Title, "| price:", items[i].total())
		sum += items[i].total()
	}
	println("----------------")
	println("Total:", sum)
}
