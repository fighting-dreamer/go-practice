package main

import (
	"fmt"
	"log"
	"strings"

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

func reverse(s string) string {
	i, j := 0, len(s)-1
	sb := []byte(s)
	for i < j {
		t := sb[i]
		sb[i] = sb[j]
		sb[j] = t

		i++
		j--
	}

	return string(sb)
}

func remove_stars_from_string(s string) string {
	l := len(s)
	st := stack.New[byte]()
	for i := 0; i < l; i++ {
		if s[i] == '*' {
			st.Pop()
		} else {
			st.Push(s[i])
		}
	}
	str := strings.Builder{}
	for !st.Empty() {
		str.WriteByte(st.Pop())
	}

	return reverse(str.String())
}

func main() {
	s := readString()
	fmt.Println(remove_stars_from_string(s))
}
