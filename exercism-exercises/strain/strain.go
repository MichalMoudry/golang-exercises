package strain

import "fmt"

func Run() {
	isEven := Keep([]int{1, 2, 3, 4, 5}, func(i int) bool { return i%2 == 0 })
	fmt.Println(isEven)

	discard := Discard([]int{1, 2, 3, 4, 5}, func(i int) bool { return i%2 == 0 })
	fmt.Println(discard)
}

func Keep[T any](items []T, predicate func(T) bool) []T {
	result := make([]T, 0)
	for _, v := range items {
		if predicate(v) {
			result = append(result, v)
		}
	}

	return result
}

func Discard[T any](items []T, predicate func(T) bool) []T {
	result := make([]T, 0)
	for _, v := range items {
		if !predicate(v) {
			result = append(result, v)
		}
	}

	return result
}
