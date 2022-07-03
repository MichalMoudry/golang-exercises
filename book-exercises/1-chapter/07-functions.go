package main

import "fmt"

func main() {
	log("10 + 15 =")
	println(add(10, 15))
	sum, product := subtract(10, 5)
	println(sum, product)
}

func add(first int, second int) (sum int) {
	sum = first + second
	return
}

func subtract(first int, second int) (sum int, product int) {
	sum = first - second
	product = first * second
	return
}

func log(message string) {
	fmt.Println(message)
}
