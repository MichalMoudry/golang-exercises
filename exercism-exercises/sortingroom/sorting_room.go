package sortingroom

import (
	"fmt"
	"strconv"
)

// DescribeNumber should return a string describing the number.
func DescribeNumber(f float64) string {
	return fmt.Sprintf("This is the number %.1f", f)
}

type NumberBox interface {
	Number() int
}

// DescribeNumberBox should return a string describing the NumberBox.
func DescribeNumberBox(nb NumberBox) string {
	return fmt.Sprintf("This is a box containing the number %.1f", float64(nb.Number()))
}

type FancyNumber struct {
	n string
}

func (i FancyNumber) Value() string {
	return i.n
}

type FancyNumberBox interface {
	Value() string
}

// ExtractFancyNumber should return the integer value for a FancyNumber
// and 0 if any other FancyNumberBox is supplied.
func ExtractFancyNumber(fnb FancyNumberBox) int {
	switch value := fnb.(type) {
	case FancyNumber:
		val, err := strconv.Atoi(value.n)
		if err != nil {
			return 0
		}
		return val
	default:
		return 0
	}
}

// DescribeFancyNumberBox should return a string describing the FancyNumberBox.
func DescribeFancyNumberBox(fnb FancyNumberBox) string {
	switch value := fnb.(type) {
	case FancyNumber:
		val, err := strconv.Atoi(value.n)
		if err != nil {
			panic("Cannot convert this number!")
		}
		return fmt.Sprintf("This is a fancy box containing the number %.1f", float64(val))
	default:
		return "This is a fancy box containing the number 0.0"
	}
}

// DescribeAnything should return a string describing whatever it contains.
func DescribeAnything(i interface{}) string {
	switch value := i.(type) {
	case int, float64:
		val, ok := value.(float64)
		if !ok {
			valInt, _ := value.(int)
			return DescribeNumber(float64(valInt))
		}
		return DescribeNumber(val)
	case NumberBox:
		return DescribeNumberBox(value)
	case FancyNumberBox:
		return DescribeFancyNumberBox(value)
	default:
		return "Return to sender"
	}
}

func Run() {
	println(DescribeNumber(-12.345))
	println(ExtractFancyNumber(FancyNumber{"10"}))
	println(DescribeFancyNumberBox(FancyNumber{"10"}))
	println(DescribeAnything(42))
}
