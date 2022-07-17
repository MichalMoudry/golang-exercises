package main

import (
	"errors"
	"fmt"
)

var ErrScaleMalfunction = errors.New("sensor error")

type WeightFodder interface {
	FodderAmount() (float64, error)
}

type Fodder struct {
	amount float64
}

func (f *Fodder) FodderAmount() (float64, error) {
	if f.amount < 0 {
		return 0, errors.New("negative fodder")
	}
	return f.amount, nil
}

// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {
	if cows == 0 {
		return 0, errors.New("division by zero")
	} else if cows < 0 {
		return 0, errors.New(fmt.Sprintf("silly nephew, there cannot be %d cows", cows))
	}
	fodderAmount, err := weightFodder.FodderAmount()
	if err != nil {
		if fodderAmount > 0 && err == ErrScaleMalfunction {
			fodderAmount *= 2
			return fodderAmount / float64(cows), nil
		} else if fodderAmount < 0 && err.Error() == "non-scale error" {
			return 0, errors.New("non-scale error")
		} else if fodderAmount < 0 && err.Error() == "negative fodder" {
			return 0, errors.New("negative fodder")
		}
		return 0, errors.New(err.Error())
	} else if fodderAmount < 0 {
		return 0, errors.New("negative fodder")
	}
	return fodderAmount / float64(cows), nil
}

func main() {
	fodder := Fodder{
		amount: 10,
	}
	res, err := DivideFood(&fodder, 2)
	println(int(res), err)
}
