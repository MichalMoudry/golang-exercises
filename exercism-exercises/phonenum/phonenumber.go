package phonenum

import (
	"errors"
	"fmt"
	"strings"
)

var (
	InvalidPhoneNumberErr = errors.New("invalid phone number")
)

func Number(phoneNumber string) (string, error) {
	replacer := strings.NewReplacer(
		"(", "",
		")", "",
		" ", "",
		"-", "",
		".", "",
		"+", "",
	)
	trimmedNumber := replacer.Replace(phoneNumber)
	numberLen := len(trimmedNumber)

	if numberLen == 11 && trimmedNumber[0] != '1' {
		return "", InvalidPhoneNumberErr
	} else if numberLen == 11 {
		trimmedNumber = trimmedNumber[1:]
	} else if numberLen != 10 {
		return "", InvalidPhoneNumberErr
	}
	return trimmedNumber, nil
}

func AreaCode(phoneNumber string) (string, error) {
	number, err := Number(phoneNumber)
	if err != nil {
		return "", err
	}
	return number[0:3], nil
}

func Format(phoneNumber string) (string, error) {
	num, err := Number(phoneNumber)
	if err != nil {
		return "", err
	}
	areaCode, err := AreaCode(num)
	if err != nil {
		return "", err
	}
	exchangeCode := num[3:6]
	return fmt.Sprintf("(%s) %s-%s", areaCode, exchangeCode, num[6:10]), nil
}

func Run() {
	// Test case 1
	input1 := "(223) 456-7890"
	num, err := Number(input1)
	fmt.Println(num, err, num == "2234567890")

	areaCode, err := AreaCode(input1)
	fmt.Println(areaCode, err, areaCode == "223")

	formatted, err := Format(input1)
	fmt.Println(formatted, err, formatted == "(223) 456-7890")

	fmt.Println("------------------------")
	// Test case 2
	input2 := "223.456.7890"
	num, err = Number(input2)
	fmt.Println(num, err, num == "2234567890")

	areaCode, err = AreaCode(input2)
	fmt.Println(areaCode, err, areaCode == "223")

	formatted, err = Format(input2)
	fmt.Println(formatted, err, formatted == "(223) 456-7890")
}
