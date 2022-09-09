package main

import (
	"fmt"
)

func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}
	stack := []rune{}
	m := map[rune]rune{
		'(': ')',
		'{': '}',
		'[': ']',
	}
	for _, e := range s {
		if value, ok := m[e]; ok {
			stack = append(stack, value)
			continue
		}
		l := len(stack) - 1
		if l < 0 || stack[l] != e {
			return false
		}
		stack = stack[:l]
	}
	return len(stack) == 0
}

func main() {
	b := isValid("()")
	fmt.Println(b)
}
