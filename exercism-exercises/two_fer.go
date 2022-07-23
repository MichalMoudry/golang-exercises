package main

import "fmt"

func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}

func main() {
	println(ShareWith("Bob"))
	println(ShareWith(""))
	println(ShareWith("Alice"))
}
