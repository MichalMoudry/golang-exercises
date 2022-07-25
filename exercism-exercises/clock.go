package main

import "fmt"

type Clock struct {
	hours   int
	minutes int
}

// Function for obtaining hours and minutes overlaps.
func getOverlaps(h, m int) (hours, minutes int) {
	if h < 0 {
		h = -h
	}
	if m < 0 {
		m = -m
	}
	hours = h / 24
	minutes = m / 60
	return
}

func New(h, m int) Clock {
	areHoursNegative := h < 0
	areMinutesNegative := m < 0
	println(areHoursNegative, areMinutesNegative)
	/*if h > 23 {
		h -= 24 * (h / 24) // (h / 24) for continuous rolling = how many days is in hours
	}
	if m > 59 {
		newHours := m / 60
		h += newHours
		m -= 60 * newHours
		if h > 23 {
			h -= 24 * (h / 24)
		}
	}
	if h < 0 {
		h = 23 + (h + 1)
	}*/
	hoursOverlap, minutesOverlap := getOverlaps(h, m)
	println(hoursOverlap, minutesOverlap)
	hoursOverlap += minutesOverlap
	minutesOverlap = 0
	println(hoursOverlap, minutesOverlap)
	return Clock{
		hours:   h,
		minutes: m,
	}
}

func (c Clock) Add(m int) Clock {
	/*if m >= 60 {
		c.hours += m / 60
		c.minutes -= 60
	}
	c.minutes += m
	if c.hours == 24 {
		c.hours = 0
	}*/
	return c
}

func (c Clock) Subtract(m int) Clock {
	/*if c.hours == 0 {
		c.hours = 24
	}
	if m >= 60 {
		c.hours -= m / 60
		c.minutes -= 60
	}
	c.minutes += m*/
	return c
}

func (c Clock) String() string {
	var hoursLeadingZero string = ""
	var minutesLeadingZero string = ""
	if c.hours < 10 {
		hoursLeadingZero = "0"
	}
	if c.minutes <= 10 {
		minutesLeadingZero = "0"
	}
	return fmt.Sprintf("%s%d:%s%d", hoursLeadingZero, c.hours, minutesLeadingZero, c.minutes)
}

func main() {
	//clock := New(50, 184)
	New(50, -184)
	/*println(clock.String())
	clock = clock.Add(61)
	println(clock.String())
	clock = clock.Subtract(61)
	println(clock.String())*/
}
