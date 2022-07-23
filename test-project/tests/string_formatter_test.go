package tests

import (
	"test-project/helpers"
	"testing"
)

func TestFormatString(t *testing.T) {
	formattedString := helpers.FormatString("Test string: %s", "test-value")
	expectedString := "Test string: test-value"
	if formattedString != expectedString {
		t.Errorf("\nFormatted string has incorrect value.\nActual: %s\nExpected: %s", formattedString, expectedString)
	}
}
