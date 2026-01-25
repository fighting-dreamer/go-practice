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

func isMatch(openingP byte, closingP byte) bool {
	return (openingP == '(' && closingP == ')') || (openingP == '[' && closingP == ']') || (openingP == '{' && closingP == '}')
}

func isValidParenthesis(s string) bool {
	st := stack.New[byte]()
	for i := 0; i < len(s); i++ {
		if s[i] == ')' || s[i] == ']' || s[i] == '}' {
			if st.Empty() || !isMatch(st.Top(), s[i]) {
				return false
			}
			st.Pop()
		} else {
			st.Push(s[i])
		}
	}

	return st.Empty()
}

func main() {
	s := readString()
	fmt.Printf("isValidParenthesis(s): %v\n", isValidParenthesis(s))
}
