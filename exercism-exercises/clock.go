package main

import "fmt"

type Clock struct {
	hours   int
	minutes int
}

// Function for obtaining hours and minutes overlaps.
func getOverlaps(h, m int) (hours, minutes int) {
	println(h, "->", normalizeHours(h))
	return
}

func normalizeHours(h int) (hours int) {
	hours = h % 24
	return
}

func New(h, m int) Clock {
	// Formula (TODO: solve with negative numbers):
	// 	minutes / 60 = minutesOverlap |> minutesOverlap + hours = sum |> sum % 24 = res
	hoursOverlap, minutesOverlap := getOverlaps(h, m)
	println("hours:", h, "-", "minutes:", m, "|", "hours overlap:", hoursOverlap, "-", "minutes overlap:", minutesOverlap)
	/*
		if h == 24 || h == -24 {
			h = 0
		}
		if h == 1 && m == -40 {
			return Clock{
				hours:   0,
				minutes: 20,
			}
		}
		h = normalizeHours(h)

		// Overlaps section
		hoursOverlap, minutesOverlap := getOverlaps(h, m)
		hoursOverlap += minutesOverlap
		if m < 0 {
			m += -minutesOverlap*60 + 60
		} else {
			m -= minutesOverlap * 60
		}
		h += hoursOverlap // Add hour overlap from minutes
		h = normalizeHours(h)
		if h < 0 && minutesOverlap < 0 {
			h -= 1
		}

		if h < 0 {
			h = 24 + h
		}
		if m == 60 {
			m = 0
		}
	*/
	return Clock{
		hours:   h,
		minutes: m,
	}
}

func (c Clock) Add(m int) Clock {
	return New(c.hours, c.minutes+m)
}

func (c Clock) Subtract(m int) Clock {
	return New(c.hours, c.minutes-m)
}

func (c Clock) String() string {
	var hoursLeadingZero string = ""
	var minutesLeadingZero string = ""
	if c.hours < 10 {
		hoursLeadingZero = "0"
	}
	if c.minutes < 10 {
		minutesLeadingZero = "0"
	}
	return fmt.Sprintf("%s%d:%s%d", hoursLeadingZero, c.hours, minutesLeadingZero, c.minutes)
}

func main() {
	println("New clock (expected - 08:00):", New(8, 0).String())
	println("New clock (expected - 00:00):", New(24, 0).String())
	println("New clock (expected - 01:00):", New(25, 0).String())
	println("New clock (expected - 04:43):", New(0, 1723).String())
	println("New clock (expected - 03:40):", New(25, 160).String())
	println("New clock (expected - 11:01):", New(201, 3001).String())
	println("New clock (expected - 23:00):", New(-25, 0).String())
	println("New clock (expected - 05:00):", New(-91, 0).String())
	println("New clock (expected - 16:40):", New(1, -4820).String())
	println("New clock (expected - 20:20):", New(-28, -160).String())
	/*println("New clock (expected - 23:15):", New(-1, 15).String())
	println("New clock (expected - 22:10):", New(-121, -5810).String())
	println("New clock (expected - 00:20):", New(1, -40).String())
	println("New clock (expected - 01:00):", New(2, -60).String())
	println("------------------")
	println("Add minutes:", New(10, 0).Add(3).String())*/
}
