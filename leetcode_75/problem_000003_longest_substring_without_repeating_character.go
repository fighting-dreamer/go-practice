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

func longest_substring_without_repeating_characters(s string) string {
	charMap := map[byte]int{}
	for i := 0; i < 26; i++ {
		charMap[byte('a'+i)] = 0
	}
	i, j, n := 0, 0, len(s)

	res_i, res_j, resSize := 0, 0, 0

	for i < n && j < n {
		charMap[s[j]]++
		// a character is repeating or not repeating.
		if charMap[s[j]] == 1 {
			// adding a unique character in the map
			if j-i+1 > resSize {
				resSize = j - i + 1
				res_i, res_j = i, j
			}
			j++
		} else {
			// we have a duplicate for s[j]
			for s[i] != s[j] {
				charMap[s[i]]--
				i++
			}
			charMap[s[i]]--
			i++

			if j-i+1 > resSize {
				resSize = j - i + 1
				res_i, res_j = i, j
			}
			j++
		}
	}

	return s[res_i : res_j+1]
}

func main() {
	s := readString()

	fmt.Println(longest_substring_without_repeating_characters(s))
}
