package helpers

import "fmt"

func FormatString(format, input string) string {
	return fmt.Sprintf(format, input)
}
