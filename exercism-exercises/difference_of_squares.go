package main

import "math"

func sum(number int) int {
	var res int = 0
	for i := 0; i < number; i++ {
		res += 1 + i
	}
	return res
}

func SquareOfSum(n int) int {
	return int(math.Pow(float64(sum(n)), 2))
}

func SumOfSquares(n int) int {
	var res float64 = 0
	for i := 0; i < n; i++ {
		res += math.Pow(float64(1+i), 2)
	}
	return int(res)
}

func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}

func main() {
	println("SquareOfSum (expected - 3025):", SquareOfSum(10))
	println("SumOfSquares (expected - 385):", SumOfSquares(10))
	println("Difference (expected - 2640):", Difference(10))
	println("Difference (expected - 170):", Difference(5))
}
