package isogram

import "strings"

func IsIsogram(word string) bool {
	word = strings.ReplaceAll(word, "-", "")
	word = strings.ReplaceAll(word, " ", "")
	word = strings.ToLower(word)
	split := strings.Split(word, "")
	var firstIndex int
	var lastIndex int
	for _, value := range split {
		firstIndex = strings.Index(word, value)
		lastIndex = strings.LastIndex(word, value)
		if firstIndex != lastIndex {
			return false
		}
	}
	return true
}

func Run() {
	println("lumberjacks (expected - true):", IsIsogram("lumberjacks"), "\n")
	println("background (expected - true):", IsIsogram("background"), "\n")
	println("downstream (expected - true):", IsIsogram("downstream"), "\n")
	println("six-year-old (expected - true):", IsIsogram("six-year-old"), "\n")
	println("pass (expected - false):", IsIsogram("pass"), "\n")
	println("isograms (expected - false):", IsIsogram("isograms"), "\n")
	println("Alphabet (expected - false):", IsIsogram("Alphabet"), "\n")
	println("Emily Jung Schwartzkopf (expected - true):", IsIsogram("Emily Jung Schwartzkopf"), "\n")
}
