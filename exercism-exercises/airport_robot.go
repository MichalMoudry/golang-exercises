package main

import "fmt"

type Greeter interface {
	LanguageName() string
	Greet(visitor string) string
}

func SayHello(visitor string, greeter Greeter) string {
	return greeter.Greet(visitor)
}

type Italian struct{}

func (Italian) LanguageName() string {
	return "Italian"
}

func (Italian) Greet(visitor string) string {
	return fmt.Sprintf("Cialo %s!", visitor)
}

type Portuguese struct{}

func (Portuguese) LanguageName() string {
	return "Portuguese"
}

func (Portuguese) Greet(visitor string) string {
	return fmt.Sprintf("Olá %s!", visitor)
}

func main() {
	italianRobot := Italian{}
	fmt.Println(italianRobot.LanguageName())
	fmt.Println(italianRobot.Greet("Michal"))
}
