package main

import (
	"fmt"
)

// Main function.
func main() {
	// Print 10 sentences.
	for i := 0; i < 15; i++ {
		println(i, "test sentence")
		if i == 10 {
			break
		}
	}
	var run = true
	answer := ""
	fmt.Print("Enter an input: ")
	for run {
		fmt.Scan(&answer)
		println("You entered:", answer)
		if answer == "quit" {
			run = false
			//os.Exit(0)
		}
	}
	arr := []string{"arg1", "arg2", "arg3"}
	for i, s := range arr {
		fmt.Printf("index: %d, item: %s \n", i, s)
	}
}
