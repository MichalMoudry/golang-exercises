package vehiclepurchase

import "strings"

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	if kind == "car" || kind == "truck" {
		return true
	} else {
		return false
	}
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
	var builder strings.Builder
	if option1 < option2 {
		builder.WriteString(option1)
	} else {
		builder.WriteString(option2)
	}
	builder.WriteString(" is clearly the better choice.")
	return builder.String()
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	var onePercent = originalPrice / 100
	if age < 3 {
		return onePercent * 80
	} else if age >= 3 && age < 10 {
		return onePercent * 70
	} else {
		return onePercent * 50
	}
}

func Run() {
	println("Needs license (expected - true):", NeedsLicense("car"))
	println("Needs license (expected - false):", NeedsLicense("bike"))
	println("Choose vehicle (expected - 'Toyota Corolla is clearly the better choice.'):", ChooseVehicle("Wuling Hongguang", "Toyota Corolla"))
	println("Choose vehicle (expected - 'Volkswagen Beetle is clearly the better choice.'):", ChooseVehicle("Volkswagen Beetle", "Volkswagen Golf"))
	println("Calculate price (expected - 800):", int(CalculateResellPrice(1000, 1)))
	println("Calculate price (expected - 700):", int(CalculateResellPrice(1000, 5)))
	println("Calculate price (expected - 500):", int(CalculateResellPrice(1000, 15)))
}
