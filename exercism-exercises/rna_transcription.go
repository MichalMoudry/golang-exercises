package main

import (
	"fmt"
	"strings"
)

var dnaToRnaMap map[rune]rune = map[rune]rune{
	'G': 'C',
	'C': 'G',
	'T': 'A',
	'A': 'U',
}

func ToRNA(dna string) string {
	if dna == "" {
		return ""
	}
	strBuilder := strings.Builder{}
	for _, v := range dna {
		strBuilder.WriteRune(dnaToRnaMap[v])
	}
	return strBuilder.String()
}

func main() {
	fmt.Println(dnaToRnaMap)
	fmt.Println(ToRNA("T"))
}
