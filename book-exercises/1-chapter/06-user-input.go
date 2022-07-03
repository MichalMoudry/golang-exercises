package main

import "fmt"

func main() {
	/*
		var response, a2 string
		print("Enter an input: ")
		fmt.Scan(&response, &a2)
		fmt.Println("User typed:", response, "-", a2)
	*/
	var prefix string
	var num int
	fmt.Scanf("%3s%d", &prefix, &num)
	println("Prefix:", prefix, "|", "ID:", num)
}
