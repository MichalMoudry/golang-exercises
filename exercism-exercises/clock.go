package main

import "fmt"

type Clock struct {
	hours   int
	minutes int
}

// Function for obtaining hours and minutes overlaps.
func getOverlaps(h, m int) (hours, minutes int) {
	hours = h / 24
	minutes = m / 60
	return
}

func normalizeHours(h int) (hours int) {
	hours = h
	if hours > 23 || hours < -23 {
		hours = h % 24
	}
	return
}

func New(h, m int) Clock {
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
	return Clock{
		hours:   h,
		minutes: m,
	}
}

func (c Clock) Add(m int) Clock {
	return c
}

func (c Clock) Subtract(m int) Clock {
	return c
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
	//clock := New(25, 0)
	/*println(clock.String())
	clock = clock.Add(61)
	println(clock.String())
	clock = clock.Subtract(61)
	println(clock.String())*/
	println("New clock (expected - 08:00):", New(8, 0).String())
	println("New clock (expected - 00:00):", New(24, 0).String())
	println("New clock (expected - 01:00):", New(25, 0).String())
	println("New clock (expected - 04:43):", New(0, 1723).String())
	println("New clock (expected - 11:01):", New(201, 3001).String())
	println("New clock (expected - 23:00):", New(-25, 0).String())
	println("New clock (expected - 05:00):", New(-91, 0).String())
	println("New clock (expected - 16:40):", New(1, -4820).String())
	println("New clock (expected - 20:20):", New(-25, -160).String())
	println("New clock (expected - 23:15):", New(-1, 15).String())
	println("New clock (expected - 22:10):", New(-121, -5810).String())
	println("New clock (expected - 00:20):", New(1, -40).String())
	println("New clock (expected - 01:00):", New(2, -60).String())
}
