package main

// IntList is an abstraction of a list of integers which we can define methods on
type IntList []int

func (s IntList) Foldl(fn func(int, int) int, initial int) int {
	return 0
}

func (s IntList) Foldr(fn func(int, int) int, initial int) int {
	return 0
}

func (s IntList) Filter(fn func(int) bool) IntList {
	for _, v := range s {
		if fn(v) {
			s = append(s, v)
		}
	}
	return s
}

func (s IntList) Length() int {
	return len(s)
}

func (s IntList) Map(fn func(int) int) IntList {
	for i := 0; i < s.Length(); i++ {
		s[i] = fn(s[i])
	}
	return s
}

func (s IntList) Reverse() IntList {
	var res IntList
	for i := s.Length() - 1; i >= 0; i-- {
		res = append(res, s[i])
	}
	return res
}

func (s IntList) Append(lst IntList) IntList {
	return append(s, lst...)
}

func (s IntList) Concat(lists []IntList) IntList {
	for _, v := range lists {
		s = s.Append(v)
	}
	return s
}

func mutiplyBy2(i int) int {
	return i * 2
}

func main() {
	var list IntList = IntList{5, 0, 3}
	println(len(list))
	list = list.Append(IntList{8, 7, 4, 3})
	println(len(list))
	println("----------------")
	list.Reverse()
	list.Map(mutiplyBy2)

	//var list2 IntList = IntList{1, 2, 3, 4, 5, 1, 3, 5}
}
