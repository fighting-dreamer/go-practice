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

func valid_parentheses(s string) bool {
	st := stack.New[byte]()

	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			st.Push(s[i])
		} else {
			if (s[i] == ')' && st.Top() == '(') || (s[i] == ']' && st.Top() == '[') || (s[i] == '}' && st.Top() == '{') {
				st.Pop()
			} else {
				return false
			}
		}
	}

	return st.Empty()
}

func valid_parentheses_std(s string) bool {
	pairs := map[byte]byte{
		'}': '{',
		')': '(',
		']': '[',
	}

	st := make([]byte, 0)
	for _, char := range []byte(s) {
		openingP, ok := pairs[char]
		if !ok { // it is an opening one
			st = append(st, char)
			continue
		}

		if len(st) == 0 { // we got a closing bracket/parentheses and no opening one in the stack already. => invalid parentheses
			return false
		}

		if st[len(st)-1] != openingP {
			// mis matching of the opening parentheses => invalid parentheses
			return false
		}

		// pop the stack
		st = st[:len(st)-1]
	}

	return (len(st) == 0) // if the stack is not empty => some unbalanced parentheses => invalid parentheses
}

func main() {
	s := readString()
	fmt.Println(valid_parentheses_std(s))
}
