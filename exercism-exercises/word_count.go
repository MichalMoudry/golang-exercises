package main

import (
	"regexp"
	"strings"
)

type Frequency map[string]int

func WordCount(phrase string) Frequency {
	var res Frequency = make(Frequency)
	phrase = strings.ToLower(phrase)
	regex, _ := regexp.Compile("(\\w+([\"']\\w+|))")
	words := regex.FindAllString(phrase, -1)
	for _, word := range words {
		res[word] += 1
	}
	return res
}

func main() {
	res := WordCount("\"That's the password: 'PASSWORD 123'!\", cried the Special Agent.\nSo I fled.")
	for k, v := range res {
		println(k, v)
	}
	println("--------")
	res = WordCount("one,\ntwo,\nthree")
	for k, v := range res {
		println(k, v)
	}
	println("--------")
	res = WordCount("one,two,three")
	for k, v := range res {
		println(k, v)
	}
	println("--------")
	res = WordCount("car: carpet as java: javascript!!&@$%^&")
	for k, v := range res {
		println(k, v)
	}
	println("--------")
	res = WordCount("'First: don't laugh. Then: don't cry. You're getting it.'")
	for k, v := range res {
		println(k, v)
	}
}
