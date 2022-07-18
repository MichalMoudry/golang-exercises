package main

import "fmt"

type TemperatureUnit int

const (
	Celsius    TemperatureUnit = 0
	Fahrenheit TemperatureUnit = 1
)

func (tu TemperatureUnit) String() string {
	if tu == Celsius {
		return "°C"
	} else if tu == Fahrenheit {
		return "°F"
	} else {
		return ""
	}
}

type Temperature struct {
	degree int
	unit   TemperatureUnit
}

func (temp Temperature) String() string {
	return fmt.Sprintf("%d %s", temp.degree, temp.unit)
}

type SpeedUnit int

const (
	KmPerHour    SpeedUnit = 0
	MilesPerHour SpeedUnit = 1
)

func (su SpeedUnit) String() string {
	if su == KmPerHour {
		return "km/h"
	} else {
		return "mph"
	}
}

type Speed struct {
	magnitude int
	unit      SpeedUnit
}

func (speed Speed) String() string {
	return fmt.Sprintf("%d %s", speed.magnitude, speed.unit)
}

type MeteorologyData struct {
	location      string
	temperature   Temperature
	windDirection string
	windSpeed     Speed
	humidity      int
}

func (metData MeteorologyData) String() string {
	return fmt.Sprintf("%s: %s, Wind %s at %s, %d%% Humidity", metData.location, metData.temperature, metData.windDirection, metData.windSpeed, metData.humidity)
}

func main() {
	celsiusUnit := Celsius
	println(celsiusUnit.String())

	celsiusTemp := Temperature{
		degree: 21,
		unit:   Celsius,
	}
	println(celsiusTemp.String())

	windSpeedYesterday := Speed{
		magnitude: 22,
		unit:      MilesPerHour,
	}
	println(windSpeedYesterday.String())

	sfData := MeteorologyData{
		location: "San Francisco",
		temperature: Temperature{
			degree: 57,
			unit:   Fahrenheit,
		},
		windDirection: "NW",
		windSpeed: Speed{
			magnitude: 19,
			unit:      MilesPerHour,
		},
		humidity: 60,
	}
	println(sfData.String())
}
