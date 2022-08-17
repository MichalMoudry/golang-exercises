package main

import (
	"strconv"
	"strings"
)

func IsValidISBN(isbn string) bool {
	isbn = strings.ReplaceAll(isbn, "-", "")
	if len(isbn) != 10 {
		return false
	}
	var number int
	var err error
	var cumulativeNumber int
	var index int = 10
	for _, value := range strings.Split(isbn, "") {
		if value == "X" {
			number = 10
		} else {
			number, err = strconv.Atoi(value)
		}
		if err != nil || (index != 1 && value == "X") {
			return false
		}
		cumulativeNumber += (index * number)
		index -= 1
	}
	return cumulativeNumber%11 == 0
}

func main() {
	strings := map[string]bool{
		"3-598-P1581-X": false,
		"3-598-2X507-9": false,
		"3598215088":    true,
		"359821507":     false,
		"00":            false,
		"3-598-21515-X": false,
		"":              false,
		"3132P34035":    false,
		"3-598-21507-X": true,
		"3-598-21508-8": true,
		"359821507X":    true,
		"98245726788":   false,
	}

	for key, value := range strings {
		println(key, "|", "Want:", value, "|", "Got:", IsValidISBN(key), "|", "Result =>", IsValidISBN(key) == value)
	}
}
