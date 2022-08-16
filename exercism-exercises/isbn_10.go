package main

import (
	"strings"
)

func IsValidISBN(isbn string) bool {
	isbn = strings.ReplaceAll(isbn, "-", "")
	if len(isbn) < 10 {
		return false
	}
	//var number int
	for _, v := range strings.Split(isbn, "") {
		println(v)
	}
	return true
}

func main() {
	strings := map[string]bool{
		"3-598-21508-8": true,
		"3-598-P1581-X": false,
		"3-598-2X507-9": false,
		"3598215088":    true,
		"359821507":     false,
		"00":            false,
		"3-598-21515-X": false,
		"":              false,
		"3132P34035":    false,
	}

	for key, value := range strings {
		println(key, "|", "Result =>", IsValidISBN(key) == value)
	}
}
