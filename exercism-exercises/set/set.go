package set

import (
	"fmt"
	"slices"
	"strings"
)

// Implement Set as a collection of unique string values.
//
// For Set.String, use '{' and '}', output elements as double-quoted strings
// safely escaped with Go syntax, and use a comma and a single space between
// elements. For example, a set with 2 elements, "a" and "b", should be formatted as {"a", "b"}.
// Format the empty set as {}.

// Define the Set type here.
type Set struct {
	items []string
}

func New() Set {
	return Set{}
}

func NewFromSlice(l []string) Set {
	if len(l) == 0 {
		return New()
	}

	resItems := make([]string, 0)
	for i, v := range l {
		if i == 0 {
			resItems = append(resItems, v)
			continue
		}
		// Skip duplicate elements
		if !slices.Contains(resItems, v) {
			resItems = append(resItems, v)
		}
	}

	return Set{items: resItems}
}

func (s Set) String() string {
	length := len(s.items)
	if length == 0 {
		return "{}"
	}

	builder := strings.Builder{}
	builder.WriteString("{")
	for i, v := range s.items {
		if i == length-1 {
			builder.WriteString(fmt.Sprintf("\"%s\"", v))
		} else {
			builder.WriteString(fmt.Sprintf("\"%s\", ", v))
		}
	}
	builder.WriteString("}")
	return builder.String()
}

func (s Set) IsEmpty() bool {
	return len(s.items) == 0
}

func (s Set) Has(elem string) bool {
	return slices.Contains(s.items, elem)
}

func (s *Set) Add(elem string) {
	if !s.IsEmpty() && s.Has(elem) {
		return
	}

	currentItem := elem
	for i, v := range s.items {
		if v > currentItem {
			s.items[i] = currentItem
			currentItem = v
		}
	}
	s.items = append(s.items, currentItem)
}

func Subset(s1, s2 Set) bool {
	s1Length := len(s1.items)
	s2Length := len(s2.items)
	if s1Length == 0 {
		return true
	} else if s1Length > s2Length {
		return false
	}

	for i, v := range s1.items {
		if s2.items[i] != v {
			return false
		}
	}
	return true
}

func Disjoint(s1, s2 Set) bool {
	return Intersection(s1, s2).IsEmpty()
}

func Equal(s1, s2 Set) bool {
	if len(s1.items) != len(s2.items) {
		return false
	}

	for _, v := range s1.items {
		if !slices.Contains(s2.items, v) {
			return false
		}
	}

	return true
}

func Intersection(s1, s2 Set) Set {
	result := New()
	for _, v1 := range s1.items {
		if slices.Contains(s2.items, v1) {
			result.Add(v1)
		}
	}
	return result
}

func Difference(s1, s2 Set) Set {
	if len(s1.items) == 0 {
		return New()
	} else if len(s2.items) == 0 {
		return s1
	}

	resultSet := New()
	for _, v := range s1.items {
		if !slices.Contains(s2.items, v) {
			resultSet.Add(v)
		}
	}

	return resultSet
}

func Union(s1, s2 Set) Set {
	arr := make([]string, len(s1.items)+len(s2.items))

	for i, v := range s1.items {
		arr[i] = v
	}

	offset := len(s1.items)
	for i, v := range s2.items {
		arr[i+offset] = v
	}

	return NewFromSlice(arr)
}

func Run() {
	set := New()
	fmt.Println(set, set.IsEmpty())

	set = NewFromSlice([]string{"a", "b"})
	fmt.Println(set, set.IsEmpty())
	fmt.Println("Has b =>", set.Has("b"))

	set = NewFromSlice([]string{"a", "b", "d"})
	set.Add("c")
	fmt.Println("Add:", set)

	set = NewFromSlice([]string{"a", "a", "a", "b", "b", "c", "b"})
	fmt.Println(set)

	isSubset := Subset(set, NewFromSlice([]string{"a", "b", "c"}))
	fmt.Println("Is subset:", isSubset, Subset(New(), NewFromSlice([]string{"a"})))

	isEqual := Equal(NewFromSlice([]string{"b", "a"}), NewFromSlice([]string{"a", "b"}))
	fmt.Println("Is equal:", isEqual)

	intersection := Intersection(NewFromSlice([]string{"a", "b", "c", "d"}), NewFromSlice([]string{"a", "b", "d"}))
	fmt.Println("Intersection:", intersection)

	//diff := Difference(NewFromSlice([]string{"a", "b", "c", "d"}), NewFromSlice([]string{"a", "b", "d", "g"}))
	diff := Difference(NewFromSlice([]string{"c", "b", "a"}), NewFromSlice([]string{"b", "d"}))
	fmt.Println("Diff:", diff)

	disjoint := Disjoint(NewFromSlice([]string{"a", "b", "c", "d"}), NewFromSlice([]string{"x", "y", "z"}))
	fmt.Println("Disjoint:", disjoint)

	union := Union(NewFromSlice([]string{"a", "b", "c", "d"}), NewFromSlice([]string{"x", "y", "z"}))
	fmt.Println("Union", union)
}
