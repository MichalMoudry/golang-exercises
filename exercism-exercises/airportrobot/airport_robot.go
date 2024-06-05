package airportrobot

import "fmt"

type Greeter interface {
	LanguageName() string
	Greet(visitor string) string
}

func SayHello(visitor string, greeter Greeter) string {
	return fmt.Sprintf("I can speak %s: %s", greeter.LanguageName(), greeter.Greet(visitor))
}

type Italian struct{}

func (Italian) LanguageName() string {
	return "Italian"
}

func (Italian) Greet(visitor string) string {
	return fmt.Sprintf("Ciao %s!", visitor)
}

type Portuguese struct{}

func (Portuguese) LanguageName() string {
	return "Portuguese"
}

func (Portuguese) Greet(visitor string) string {
	return fmt.Sprintf("Olá %s!", visitor)
}

func Run() {
	italianRobot := Italian{}
	fmt.Println(italianRobot.LanguageName())
	fmt.Println(italianRobot.Greet("Michal"))
	fmt.Println(SayHello("Tomaso Giulio Micheli", italianRobot))

	portugeseRobot := Portuguese{}
	fmt.Println(SayHello("Manuela Alberto", portugeseRobot))
}
