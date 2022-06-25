package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

const Greeting string = "hello"

func main() {
	fmt.Println(Greeting)
	fmt.Println("current time:", time.Now())
	n := 100
	fmt.Println("random number:", rand.Intn(n))
	println("square root of 15:", math.Sqrt(15))
	var pi float32
	pi = math.Pi
	println("math.Pi:", pi)
	println("14 + 132 =", add(14, 132))
}

func add(x int, y int) int {
	return x + y
}
