package partyrobot

import (
	"fmt"
	"strings"
)

// Welcome greets a person by name.
func Welcome(name string) string {
	return fmt.Sprintf("Welcome to my party, %s!", name)
}

// HappyBirthday wishes happy birthday to the birthday person and exclaims their age.
func HappyBirthday(name string, age int) string {
	return fmt.Sprintf("Happy birthday %s! You are now %d years old!", name, age)
}

// AssignTable assigns a table to each guest.
func AssignTable(name string, table int, neighbor, direction string, distance float64) string {
	var tableBuilder strings.Builder
	tableBuilder.WriteString(Welcome(name))
	tableBuilder.WriteString("\n")
	tableBuilder.WriteString(fmt.Sprintf("You have been assigned to table %03d. Your table is %s, exactly %2.1f meters from here.", table, direction, distance))
	tableBuilder.WriteString("\n")
	tableBuilder.WriteString(fmt.Sprintf("You will be sitting next to %s.", neighbor))
	return tableBuilder.String()
}

func Run() {
	println(Welcome("Michal"))
	println(HappyBirthday("Frank", 55))
	println(AssignTable("Christiane", 101, "Frank", "on the left", 23.7834298))
	println(AssignTable("Chihiro", 22, "Akachi Chikondi", "straight ahead", 9.239438))
}
