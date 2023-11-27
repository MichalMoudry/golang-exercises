package main

import "fmt"

// IntList is an abstraction of a list of integers which we can define methods on
type IntList []int

func (s IntList) Foldl(fn func(int, int) int, initial int) int {
	return 0
}

func (s IntList) Foldr(fn func(int, int) int, initial int) int {
	return 0
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
	return IntList{}
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
	return IntList{}
}

func main() {
	empty_slice := IntList{}
	simple_slice := IntList{1, 2, 3}

	fmt.Println(empty_slice.Length())
	fmt.Println(simple_slice.Length())
	fmt.Println(simple_slice.Reverse())
	fmt.Println(simple_slice.Append(IntList{5, 4}))
	fmt.Println(simple_slice.Append(IntList{5, 4}).Filter(func(i int) bool { return i%2 == 1 }))
}
