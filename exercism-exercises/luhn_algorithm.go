package main

import (
	"strconv"
	"strings"
)

func Valid(id string) bool {
	id = strings.ReplaceAll(id, " ", "")
	idLength := len(id)
	if idLength <= 1 {
		return false
	}
	idSplit := strings.Split(id, "")
	var sum int = 0
	parity := (idLength - 2) % 2
	for i := len(idSplit) - 1; i >= 0; i-- {
		number, err := strconv.Atoi(idSplit[i])
		if err != nil {
			return false
		}
		if i%2 == parity {
			number *= 2
			if number > 9 {
				number -= 9
			}
		}
		sum += number
	}
	if sum%10 == 0 {
		return true
	}
	return false
}

func main() {
	println("Is valid (expected - true):", Valid("4539 3195 0343 6467"))
	println("Is valid (expected - true):", Valid("055 444 285"))
	println("Is valid (expected - true):", Valid("234 567 891 234"))

	println("Is valid (expected - false):", Valid("8273 1232 7352 0569"))
	println("Is valid (expected - false):", Valid("0"))
}
