package twofer

import "fmt"

func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}

func Run() {
	println(ShareWith("Bob"))
	println(ShareWith(""))
	println(ShareWith("Alice"))
}
