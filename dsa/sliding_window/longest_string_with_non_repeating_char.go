package main

import (
	"fmt"
	"log"
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

func longest_string_with_unique_chars(s string) string {
	resLen := 0
	start, end := 0, -1
	n := len(s)
	charMap := map[byte]int{}
	for i, j := 0, 0; i < n && j < n; {
		charMap[s[j]]++

		if charMap[s[j]] == 1 {
			// new unique char added
			fmt.Println(s[i : j+1])
			if j-i+1 > resLen {
				start = i
				end = j
				resLen = j - i + 1
			}
			j++
		} else {
			for s[i] != s[j] {
				delete(charMap, s[i])
				i++
			}

			charMap[s[i]]--
			i++
			fmt.Println(" -> ", s[i:j+1])
			if j-i+1 > resLen {
				start = i
				end = j
				resLen = j - i + 1
			}

			j++
		}
	}

	return s[start : end+1]
}

func main() {
	s := readString()
	fmt.Println(longest_string_with_unique_chars(s))
}
