package phonenum

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var ErrInvalidPhoneNumber = errors.New("invalid phone number")

func Number(phoneNumber string) (string, error) {
	replacer := strings.NewReplacer("(", "", ")", "", " ", "", "-", "", ".", "", "+1", "")
	result := replacer.Replace(phoneNumber)
	resultLength := len(result)
	if resultLength > 11 {
		return "", ErrInvalidPhoneNumber
	} else if resultLength == 11 && result[0] != '1' {
		return "", ErrInvalidPhoneNumber
	} else if result[3] == '0' || result[3] == '1' { //Validate exchange codes
		return "", ErrInvalidPhoneNumber
	} else if _, err := strconv.Atoi(result); err != nil {
		return "", err
	}
	if resultLength == 11 {
		result = strings.TrimPrefix(result, "1")
	}
	return result, nil
}

func AreaCode(phoneNumber string) (string, error) {
	number, err := Number(phoneNumber)
	if err != nil {
		return "", err
	}
	return number[:3], nil
}

func Format(phoneNumber string) (string, error) {
	return phoneNumber, nil
}

func Run() {
	inputs := []string{
		"+1 (613)-995-0253",
		"613-995-0253",
		"1 613 995 0253",
		"613.995.0253",
		"22234567890",
		"321234567890",
		"523-abc-7890",
		"(023) 456-7890",
		"(223) 056-7890",
		"(123) 456-7890",
		"12234567890",
	}

	for _, v := range inputs {
		str, err := Number(v)
		if err != nil {
			fmt.Println(v, "=>", err)
		} else {
			fmt.Println(v, "=>", str)
		}
	}

	areaCodeInputs := []string{
		"(223) 456-7890",
	}

	fmt.Println("-- Area codes")
	for _, v := range areaCodeInputs {
		code, err := AreaCode(v)
		if err != nil {
			fmt.Println(v, "=> err:", err)
		} else {
			fmt.Println(v, "=>", code)
		}
	}

	fmt.Println("-- Format")
	formatInput := []string{
		"223.456.7890",
	}
	for _, v := range formatInput {
		res, err := Format(v)
		if err != nil {
			fmt.Println(v, "=> err:", err)
		} else {
			fmt.Println(v, "=>", res)
		}
	}
}
