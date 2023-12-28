package main

import (
	"fmt"
	"strings"
)

func Number(phoneNumber string) (string, error) {
	replacer := strings.NewReplacer(
		"(", "",
		")", "",
		" ", "",
		"-", "",
	)
	return replacer.Replace(phoneNumber), nil
}

func AreaCode(phoneNumber string) (string, error) {
	number, err := Number(phoneNumber)
	if err != nil {
		return "", nil
	}
	return number[0:3], nil
}

func Format(phoneNumber string) (string, error) {
	panic("Please implement the Format function")
}

func main() {
	input1 := "(223) 456-7890"
	fmt.Println(Number(input1))
	fmt.Println(AreaCode(input1))
}
