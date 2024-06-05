package techplace

import "strings"

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return "Welcome to the Tech Palace, " + strings.ToUpper(customer)
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	var builder strings.Builder
	var border string = strings.Repeat("*", numStarsPerLine)
	builder.WriteString(border)
	builder.WriteString("\n")
	builder.WriteString(welcomeMsg)
	builder.WriteString("\n")
	builder.WriteString(border)
	return builder.String()
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	return strings.TrimSpace(strings.ReplaceAll(oldMsg, "*", ""))
}

func Run() {
	println(WelcomeMessage("Judy"))
	println(AddBorder("Welcome!", 10))
	message := `
**************************
*    BUY NOW, SAVE 10%   *
**************************
`
	println(CleanupMessage(message))
}
