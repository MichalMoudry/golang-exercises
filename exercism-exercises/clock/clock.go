package clock

import "fmt"

type Clock struct {
	hours   int
	minutes int
}

func New(h, m int) Clock {
	// Formula: minutes / 60 = minutesOverlap |> minutesOverlap + hours = sum |> sum % 24 = res
	minutesOverlap := m / 60
	h = (minutesOverlap + h) % 24
	m -= (minutesOverlap * 60)
	if m < 0 {
		m = 60 + m
		h--
	}
	if h < 0 {
		h = 24 + h
	}
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

func Run() {
	println("New clock (expected - 08:00):", New(8, 0).String())
	println("New clock (expected - 00:00):", New(24, 0).String())
	println("New clock (expected - 01:00):", New(25, 0).String())
	println("New clock (expected - 04:43):", New(0, 1723).String())
	println("New clock (expected - 03:40):", New(25, 160).String())
	println("New clock (expected - 11:01):", New(201, 3001).String())
	println("New clock (expected - 23:00):", New(-25, 0).String())
	println("New clock (expected - 05:00):", New(-91, 0).String())
	println("New clock (expected - 16:40):", New(1, -4820).String())
	//println("New clock (expected - 20:20):", New(-28, -160).String())
	println("New clock (expected - 23:15):", New(-1, 15).String())
	println("New clock (expected - 22:10):", New(-121, -5810).String())
	println("New clock (expected - 00:20):", New(1, -40).String())
	println("New clock (expected - 01:00):", New(2, -60).String())
	println("------------------")
	println("Add minutes:", New(10, 0).Add(3).String())
	println("Subtract minutes:", New(0, 3).Subtract(4).String())
}
