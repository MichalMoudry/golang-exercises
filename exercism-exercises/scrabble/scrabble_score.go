package scrabble

import "strings"

var letterValues map[string]int = map[string]int{
	"a": 1, "e": 1, "i": 1, "o": 1, "u": 1, "l": 1, "n": 1, "r": 1, "s": 1, "t": 1,
	"d": 2, "g": 2,
	"b": 3, "c": 3, "m": 3, "p": 3,
	"f": 4, "h": 4, "v": 4, "w": 4, "y": 4,
	"k": 5,
	"j": 8, "x": 8,
	"q": 10, "z": 10,
}

func Score(word string) int {
	if word == "" {
		return 0
	}
	word = strings.ToLower(word)
	wordSplit := strings.Split(word, "")
	var res int = 0
	for _, char := range wordSplit {
		res += letterValues[char]
	}
	return res
}

func Run() {
	println("Score (expected - 14):", Score("cabbage"))
	println("Score (expected - 41):", Score("OxyphenButazone"))
	println("Score (expected - 8):", Score("pinata"))
}
