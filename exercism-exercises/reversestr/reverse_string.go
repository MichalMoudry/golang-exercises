package reversestr

import "strings"

func Reverse(input string) string {
	if input == "" {
		return ""
	}
	var sb strings.Builder
	split := strings.Split(input, "")
	for i := len(split) - 1; i >= 0; i-- {
		sb.WriteString(split[i])
	}
	return sb.String()
}

func Run() {
	println(Reverse("cool"))
	println(Reverse("Hello, 世界"))
}
