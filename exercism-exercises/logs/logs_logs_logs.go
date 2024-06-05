package logs

import (
	"strings"
	"unicode/utf8"
)

// Application identifies the application emitting the given log.
func Application(log string) string {
	for _, value := range log {
		if value == '❗' {
			return "recommendation"
		} else if value == '🔍' {
			return "search"
		} else if value == '☀' {
			return "weather"
		}
	}
	return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	return strings.ReplaceAll(log, string(oldRune), string(newRune))
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	return utf8.RuneCountInString(log) <= limit
}

func Run() {
	println("Application (expected - recommendation):", Application("❗ recommended search product 🔍"))
	println("Application (expected - search):", Application("executed search 🔍"))
	println("Replace (expected - '? recommended product'):", Replace("❗ recommended product", '❗', '?'))
	println("Within limit (expected - true):", WithinLimit("hello❗", 6))
	println("Within limit (expected - true):", WithinLimit("exercism❗", 10))
}
