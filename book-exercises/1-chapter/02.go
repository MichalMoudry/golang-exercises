package main

import (
	"fmt"
)

func main() {
	// Group of variables
	var (
		firstName = "Michal"
		lastName  = "Moudrý"
	)
	age := 23
	fmt.Printf("first name: %s\nlast name: %s\nage: %d\n\n", firstName, lastName, age)

	// Assignment
	var (
		players       = 3
		replay        = false
		namePlayerOne = "chris"
	)
	fmt.Println(players)
	fmt.Println(replay)
	fmt.Println(namePlayerOne)
}
