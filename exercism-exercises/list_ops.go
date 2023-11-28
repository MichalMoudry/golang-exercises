package main

import (
	"fmt"
)

// IntList is an abstraction of a list of integers which we can define methods on
type IntList []int

func (s IntList) Foldl(fn func(int, int) int, initial int) int {
	length := len(s)
	if length == 0 {
		return initial
	}
	println("-------")
	fmt.Println(s)
	var current int
	for i := 0; i < length; i++ {
		if !((i+1)%2 == 0) {
			current = s[i]
			if i == length-1 {
				initial = fn(initial, current)
				continue
			}
			initial += fn(current, s[i+1])
		}
	}
	println("-------")
	return initial
}

func (s IntList) Foldr(fn func(int, int) int, initial int) int {
	result := 0
	length := len(s)
	if length == 0 {
		return result
	}
	/*println("----------")
	maxLength := len(s) + 1
	slc := make(IntList, maxLength)
	for i, v := range s {
		slc[i] = v
	}
	slc[maxLength-1] = initial
	var next int
	for i := maxLength - 1; i >= 0; i-- {
		if i != 0 {
			next = s[i-1]
			println(slc[i], next, fn(slc[i], next))
		} else {

		}
	}
	println("------------")*/
	return result
}

func (s IntList) Filter(fn func(int) bool) IntList {
	result := IntList{}
	for i := 0; i < len(s); i++ {
		if fn(s[i]) {
			result = append(result, s[i])
		}
	}
	return result
}

func (s IntList) Length() int {
	length := 0
	for range s {
		length += 1
	}
	return length
}

func (s IntList) Map(fn func(int) int) IntList {
	result := make(IntList, len(s))
	for i, v := range s {
		result[i] = fn(v)
	}
	return result
}

func (s IntList) Reverse() IntList {
	result := make(IntList, len(s))
	for i, j := 0, len(result)-1; j >= 0; i, j = i+1, j-1 {
		result[j] = s[i]
	}
	return result
}

func (s IntList) Append(lst IntList) IntList {
	return append(s, lst...)
}

func (s IntList) Concat(lists []IntList) IntList {
	result := s
	for _, v := range lists {
		result = append(result, v...)
	}
	return result
}

func main() {
	empty_slice := IntList{}
	simple_slice := IntList{1, 2, 3}

	fmt.Println(empty_slice.Length())
	fmt.Println(simple_slice.Length())
	fmt.Println(simple_slice.Reverse())
	fmt.Println(simple_slice.Append(IntList{5, 4}))
	fmt.Println(simple_slice.Append(IntList{5, 4}).Filter(func(i int) bool { return i%2 == 1 }))
	fmt.Println(simple_slice.Map(func(i int) int { return i + 1 }))

	foldr_slice := IntList{1, 2, 3, 4, 5}
	fmt.Println(
		foldr_slice.Foldl(
			func(i1, i2 int) int {
				fmt.Println("Adding:", i1, "+", i2)
				return i1 + i2
			},
			5,
		),
	)
}
