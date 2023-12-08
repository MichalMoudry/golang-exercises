package main

import (
	"fmt"
	"strings"
)

func Transform(in map[int][]string) map[string]int {
	result := make(map[string]int)
	for i, v := range in {
		for _, item := range v {
			result[strings.ToLower(item)] = i
		}
	}
	return result
}

func main() {
	testInput := map[int][]string{
		1: {"A", "E"},
		2: {"D", "G"},
	}
	fmt.Println(Transform(testInput))
}
