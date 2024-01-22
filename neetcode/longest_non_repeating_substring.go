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

// try for : "tmmzuxt"

func length_of_non_repeating_longest_substring(s string) int {
	if s == "" {
		return 0
	}

	left := 0
	right := 1
	maxLength := 1
	mp := map[rune]int{}
	n := len(s)

	// invariant is at any time substring between [left, right] is having non-repeating.
	mp[rune(s[0])] = 1
	for left < right && right < n {
		if mp[rune(s[right])] == 1 {
			for mp[rune(s[right])] == 1 {
				delete(mp, rune(s[left]))
				left++
			}
		}
		mp[rune(s[right])] = 1
		maxLength = max(maxLength, right-left+1)

		fmt.Println(mp, maxLength, left, right)
		right++
	}

	return maxLength
}

func main() {
	s := readString()
	fmt.Println(length_of_non_repeating_longest_substring(s))
}
