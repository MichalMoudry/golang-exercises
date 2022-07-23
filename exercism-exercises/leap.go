package main

func IsLeapYear(year int) bool {
	if year%4 == 0 {
		if year%100 == 0 && year%400 != 0 {
			return false
		}
		return true
	}
	return false
}

func main() {
	println("Is leap year (expected - false):", IsLeapYear(1997))
	println("Is leap year (expected - true):", IsLeapYear(1996))
	println("Is leap year (expected - false):", IsLeapYear(1900))
	println("Is leap year (expected - true):", IsLeapYear(2000))
}
