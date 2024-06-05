package raindrops

import (
	"fmt"
	"strings"
)

func Convert(number int) string {
	var sb strings.Builder
	if number%3 == 0 {
		sb.WriteString("Pling")
	}
	if number%5 == 0 {
		sb.WriteString("Plang")
	}
	if number%7 == 0 {
		sb.WriteString("Plong")
	}
	res := sb.String()
	if res == "" {
		res = fmt.Sprintf("%d", number)
	}
	return res
}

func Run() {
	println("Raindrops (expected - Plong):", Convert(28))
	println("Raindrops (expected - PlingPlang):", Convert(30))
	println("Raindrops (expected - 34):", Convert(34))
}
