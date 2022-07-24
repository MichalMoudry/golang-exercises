package main

import (
	"errors"
	"math"
)

func Square(number int) (uint64, error) {
	if number <= 0 || number > 64 {
		return 0, errors.New("Incorrect input.")
	}
	return uint64(math.Pow(2, float64(number-1))), nil
}

func Total() uint64 {
	var res uint64 = 0
	for i := 1; i <= 64; i++ {
		temp, _ := Square(i)
		res += temp
	}
	return res
}

func main() {
	res1, _ := Square(4)
	println("Grains on sqaure (expected - 8):", res1)
	/*res2, _ := Square(32)
	println("Grains on sqaure (expected - 2147483648):", res2)
	res3, _ := Square(64)
	println("Grains on sqaure (expected - 9223372036854775808):", res3)*/
	println(Total())
}
