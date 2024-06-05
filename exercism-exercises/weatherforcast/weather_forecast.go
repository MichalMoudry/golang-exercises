// Package containing solution to Weather Forecast exercise.
package weatherforcast

// CurrentCondition represents current condition of a specific location.
var CurrentCondition string

// CurrentLocation represents current location.
var CurrentLocation string

// Forecast() return forcast for a specific location.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
