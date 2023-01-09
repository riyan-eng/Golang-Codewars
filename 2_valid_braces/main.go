package main

import (
	"fmt"
)

func ValidBraces(str string) bool {
	if len(str)%2 != 0 {
		return false
	}

	st := []rune{}
	open := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
	}

	for _, r := range str {
		if closer, ok := open[r]; ok {
			st = append(st, closer)
			continue
		}

		l := len(st) - 1
		if l < 0 || r != st[l] {
			return false
		}
		st = st[:l]
	}
	return len(st) == 0
}

var m = map[string]string{
	"{": "}",
	"(": ")",
	"[": "]",
}

func ValidBraces2(str string) bool {
	s := make([]string, 0)
	for _, r := range str {
		r := string(r)
		if len(s) > 0 && m[s[len(s)-1]] == r {
			s = s[:len(s)-1]
		} else {
			s = append(s, r)
		}
	}
	return len(s) == 0
}

func main() {
	a := "(){}[]"
	b := "([{}])"
	c := "(}"
	d := "[(])"
	e := "[({})](]"

	fmt.Println(ValidBraces(a))
	fmt.Println(ValidBraces(b))
	fmt.Println(ValidBraces(c))
	fmt.Println(ValidBraces(d))
	fmt.Println(ValidBraces(e))

	fmt.Println(ValidBraces2(a))
	fmt.Println(ValidBraces2(b))
	fmt.Println(ValidBraces2(c))
	fmt.Println(ValidBraces2(d))
	fmt.Println(ValidBraces2(e))
}
