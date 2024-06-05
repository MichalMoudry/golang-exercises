package carsassemble

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return (float64(productionRate) / successRate) / 100
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	var perHour float64 = CalculateWorkingCarsPerHour(productionRate, successRate)
	return int(perHour) / 60
}

// CalculateCost works out the cost of producing the given number of cars
func CalculateCost(carsCount int) uint {
	var surplus int = carsCount % 10
	var pairs int = (carsCount - surplus) / 10
	return uint(pairs)*95000 + uint(surplus)*10000
}

func Run() {
	println(CalculateWorkingCarsPerHour(1547, 90))
	println(CalculateWorkingCarsPerHour(221, 100))
	println(CalculateWorkingCarsPerMinute(1105, 90))
	println(CalculateCost(37))
}
