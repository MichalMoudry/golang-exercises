package main

import (
	"math"
	"strconv"
	"strings"
)

func IsNumber(n int) bool {
	numberAsString := strconv.Itoa(n)
	numberLength := len(numberAsString)
	numberSplit := strings.Split(numberAsString, "")
	var number float64
	sum := 0
	for _, value := range numberSplit {
		number, _ = strconv.ParseFloat(value, 32)
		sum += int(math.Pow(number, float64(numberLength)))
	}
	if sum == n {
		return true
	} else {
		return false
	}
}

func main() {
	println("Is Armstrong number (expected - true):", IsNumber(9))
	println("Is Armstrong number (expected - false):", IsNumber(10))
	println("Is Armstrong number (expected - true):", IsNumber(153))
	println("Is Armstrong number (expected - false):", IsNumber(154))
}
