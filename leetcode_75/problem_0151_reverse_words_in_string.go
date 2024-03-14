package main

import (
	"fmt"
	"log"
	"strings"
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

func build_proper_string(s string) string {
	// trim space
	// for each space, add word till that space to string builder and move till you encounter a non-space character.
	// add a space and repeat step  just above (update the start to when you encounter a non-space character)
	// at end, the last word would remain, add it to string builder too and return a procer string.

	// Trim Space
	sb, l := []byte(s), len(s)
	str := strings.Builder{}
	wi, wj := 0, 0
	// flag := 0
	i := 0
	for i < l && sb[i] == ' ' {
		i++
	}
	wi = i

	i = l - 1
	for i > 0 && sb[i] == ' ' {
		i--
	}
	wj = i + 1
	str.WriteString(s[wi:wj])
	s = str.String()
	// remove unwanted space in this string and build one space in word based string
	i = 0
	j := 0
	str = strings.Builder{}
	sb, l = []byte(s), len(s)

	for j < l {
		if sb[j] == ' ' {
			// fmt.Println(s[i:j])
			str.WriteString(s[i:j])
			for j < l && sb[j] == ' ' {
				j++
			}
			str.WriteByte(' ')
			i = j
		}
		j++
	}
	// fmt.Println(i, j)
	str.WriteString(s[i:j])

	return str.String()
}

func reverse(sb []byte, i, j int) {
	for i < j {
		temp := sb[i]
		sb[i] = sb[j]
		sb[j] = temp

		i++
		j--
	}

}

func reverse_words_in_string(s string) string {
	sb := []byte(build_proper_string(s))
	// fmt.Println(string(sb))
	l := len(sb)
	reverse(sb, 0, l-1)
	fmt.Println(string(sb))
	i, j := 0, 0
	for j < l {
		if sb[j] == ' ' {
			reverse(sb, i, j-1)
			i = j + 1
		}
		j++
	}
	reverse(sb, i, j-1)
	return string(sb)
}

func main() {
	s := "  hello world  "

	fmt.Println(reverse_words_in_string(s))
}
