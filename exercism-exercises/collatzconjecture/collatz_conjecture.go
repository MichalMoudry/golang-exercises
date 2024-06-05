package collatzconjecture

import (
	"errors"
	"fmt"
)

var (
	NegativeNumberErr = errors.New("only positive numbers are allowed")
)

func CollatzConjecture(n int) (int, error) {
	if n <= 0 {
		return 0, NegativeNumberErr
	}

	var steps int
	for n != 1 {
		if n%2 == 0 {
			n = n / 2
		} else {
			n = n*3 + 1
		}
		steps += 1
	}
	return steps, nil
}

func Run() {
	res1, err1 := CollatzConjecture(1000000)
	fmt.Println(res1, "|", err1)
}
