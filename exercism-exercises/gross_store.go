package main

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	var units map[string]int = map[string]int{
		"quarter_of_a_dozen": 3,
		"half_of_a_dozen":    6,
		"dozen":              12,
		"small_gross":        120,
		"gross":              144,
		"great_gross":        1728,
	}
	return units
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return make(map[string]int)
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	value, exists := units[unit]
	if !exists {
		return false
	} else {
		_, isItemInBill := bill[item]
		if !isItemInBill {
			bill[item] = value
		} else {
			bill[item] += value
		}
		return true
	}
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	value, exists := bill[item]
	unitValue, unitExists := units[unit]
	if !exists || !unitExists {
		return false
	} else {
		if value < unitValue {
			return false
		} else if value == unitValue {
			delete(bill, item)
			return true
		} else {
			bill[item] -= unitValue
			return true
		}
	}
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	quantity, exists := bill[item]
	if !exists {
		return 0, false
	} else {
		return quantity, true
	}
}

func main() {
	for key, v := range Units() {
		println(key, ":", v)
	}
	bill := NewBill()
	units := Units()
	ok := AddItem(bill, units, "carrot", "dozen")
	println(ok)
	println(RemoveItem(bill, units, "carrot", "dozen"))
}
