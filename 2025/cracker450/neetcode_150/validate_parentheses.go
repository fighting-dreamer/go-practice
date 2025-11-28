package main

import (
	"fmt"
	"log"

	"github.com/liyue201/gostl/ds/stack"
)

func readInt() int {
	var n int
	_, err := fmt.Scanf("%d", &n)
	if err != nil {
		log.Fatal("could not read input number", err)
	}
	return n
}

func readString() string {
	var s string
	_, err := fmt.Scan(&s)
	if err != nil {
		log.Fatal("could not read input string", err)
	}
	return s
}

func readDouble() float64 {
	var f float64
	_, err := fmt.Scanf("%f", &f)
	if err != nil {
		log.Fatal("could not read input float", err)
	}
	return f
}

func readChar() rune {
	var r rune
	_, err := fmt.Scanf("%c", &r)
	if err != nil {
		log.Fatal("Could not read input char", err)
	}
	return r
}

func isOpening(b byte) bool {
	return b == '(' || b == '[' || b == '{'
}

func isClosing(b byte) bool {
	return b == ')' || b == ']' || b == '}'
}

func matchingOpening(b byte) byte {
	mp := map[byte]byte{
		'}': '{',
		']': '[',
		')': '(',
	}

	return mp[b]
}

func validateParentheses(s string) bool {
	st := stack.New[byte]()

	for i := 0; i < len(s); i++ {
		fmt.Println(i, s[i])
		if isOpening(s[i]) {
			st.Push(s[i])
			continue
		}
		if isClosing(s[i]) {
			if matchingOpening(s[i]) == st.Top() {
				st.Pop()
			} else {
				return false // we found a closing parenthesis, but hte top parenthesis on stack was not matching.
			}
		}
	}

	return st.Empty() // gone through all string but still have some open parentheses not matching
}

func main() {
	s := readString()
	fmt.Println(s)
	fmt.Println(validateParentheses(s))
}
