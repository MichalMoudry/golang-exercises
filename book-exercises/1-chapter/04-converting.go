package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) >= 2 {
		arg := os.Args[1]
		fmt.Printf("Argument 1: %s\n", arg)
		fmt.Printf("Argument 1 type: %T\n", arg)
		no1, err := strconv.Atoi(arg)
		if err != nil {
			panic("Couldn't convert to number.")
		} else {
			println("10 +", no1, "=", 10+no1)
		}
	} else {
		println("No arguments were added")
	}
}
