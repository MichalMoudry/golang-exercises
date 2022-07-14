package main

import (
	"time"
)

func parseTime(layout, date string) time.Time {
	res, error := time.Parse(layout, date)
	if error != nil {
		panic(error.Error())
	}
	return res
}

// Schedule returns a time.Time from a string containing a date
func Schedule(date string) time.Time {
	return parseTime("1/2/2006 15:04:05", date)
}

// HasPassed returns whether a date has passed
func HasPassed(date string) bool {
	return parseTime("January 2, 2006 15:04:05", date).Before(time.Now())
}

// IsAfternoonAppointment returns whether a time is in the afternoon
func IsAfternoonAppointment(date string) bool {
	hour := parseTime("Monday, January 2, 2006 15:04:05", date).Hour()
	if hour >= 12 && hour < 18 {
		return true
	} else {
		return false
	}
}

// Description returns a formatted string of the appointment time
func Description(date string) string {
	dateAsTime := parseTime("1/2/2006 15:04:05", date)
	return dateAsTime.Format("You have an appointment on Monday, January 2, 2006, at 15:04.")
}

// AnniversaryDate returns a Time with this year's anniversary
func AnniversaryDate() time.Time {
	return time.Date(2020, 9, 15, 0, 0, 0, 0, time.UTC)
}

func main() {
	println("Schedule (expected - 2019-07-25 13:45:00 +0000 UTC):", Schedule("7/25/2019 13:45:00").String())
	println("Has passed (expected - true):", HasPassed("July 25, 2019 13:45:00"))
	println("Is afternoon appointment (expected - true):", IsAfternoonAppointment("Thursday, July 25, 2019 13:45:00"))
	println(Description("7/25/2019 13:45:00"))
	println(AnniversaryDate().String())
}
