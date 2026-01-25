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

func longest_non_repeating_substring(s string) int {
	n := len(s)
	res := 0
	membership := map[byte]int{}
	start_index := 0
	for i := start_index; i < n; i++ {
		lastSeen, ok := membership[s[i]]
		fmt.Println(string(s[i]), lastSeen, ok)
		if ok {
			// character is present
			for j := start_index; j <= lastSeen; j++ {
				delete(membership, s[j])
			}
			start_index = lastSeen + 1
			res = max(res, len(membership))
		} else {
			// character is not present
			membership[s[i]] = i
			res = max(res, len(membership))
		}
	}

	return res
}

func main() {
	s := readString()
	fmt.Printf("longest_non_repeating_substring(s): %v\n", longest_non_repeating_substring(s))
}
