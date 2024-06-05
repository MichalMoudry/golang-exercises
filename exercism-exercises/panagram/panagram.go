package panagram

import "strings"

func IsPangram(input string) bool {
	input = strings.ToLower(input)
	input = strings.ReplaceAll(input, " ", "")
	input = strings.ReplaceAll(input, ",", "")
	input = strings.TrimRight(input, ".")
	for i := 'a'; i <= 'z'; i++ {
		if strings.IndexRune(input, i) == -1 {
			return false
		}
	}
	return true
}

func Run() {
	println("Is panagram (expected - true):", IsPangram("The quick brown fox jumps over the lazy dog."))
	println("Is panagram (expected - false):", IsPangram("a quick movement of the enemy will jeopardize five gunboats"))
	println("Is panagram (expected - false):", IsPangram("five boxing wizards jump quickly at it"))
	println("Is panagram (expected - false):", IsPangram("7h3 qu1ck brown fox jumps ov3r 7h3 lazy dog"))
}
