package hamming

import (
	"errors"
)

func Distance(a, b string) (int, error) {
	aLength := len(a)
	bLength := len(b)
	if a == "" && b != "" {
		return 0, errors.New("disallow empty first strand")
	} else if b == "" && a != "" {
		return 0, errors.New("disallow empty second strand")
	} else if aLength > bLength {
		return 0, errors.New("disallow first strand longer")
	} else if aLength < bLength {
		return 0, errors.New("disallow second strand longer")
	}
	if a == b {
		return 0, nil
	} else {
		var res int
		for index, char := range a {
			if char != rune(b[index]) {
				res += 1
			}
		}
		return res, nil
	}
}

func Run() {
	res1, err1 := Distance("", "G")
	println(res1, err1.Error())

	res2, _ := Distance("GGACTGAAATCTG", "GGACTGAAATCTG")
	println(res2)

	res3, _ := Distance("GGACGGATTCTG", "AGGACGGATTCT")
	println(res3)
}
