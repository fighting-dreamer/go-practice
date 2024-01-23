package main

// STAR

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

func bruteforce_longest_repeating_character_with_replacement(s string, k int) int {
	// for each substring, we find the character with max count
	// we check if the length of that sub-string - max_repeating_char_in_that_substring <= k
	// coz we can replace the other characters
	n := len(s)
	mp := map[rune]int{}
	maxVal := 0
	maxRes := 0
	for l := 1; l <= n; l++ {
		// fmt.Println("ran for ", l)
		for i := 0; i <= n-l; i++ {
			start := i
			end := i + l

			for k2 := range mp {
				delete(mp, k2)
			}
			maxVal = 0
			for start < end {
				mp[rune(s[start])]++
				maxVal = max(maxVal, mp[rune(s[start])])
				start++
			}

			if l-maxVal <= k {
				maxRes = max(maxRes, l)
			}
			// fmt.Println(i, i+l, maxRes)
		}
	}

	return maxRes
}

func longest_repeating_character_with_replacement(s string, k int) int {
	// return bruteforce_longest_repeating_character_with_replacement(s, k)

	// we understand that the windowLength - maximum_occuring_char_in_that_window should be less than or equal to k
	// start with L = 0, R = 0
	// if the above computation is satisfied, increment R
	// if computation is satisfied, increment L

	left := 0
	right := 0
	windowLen := func() int { return right - left + 1 }

	res := 0
	count := make([]int, 26)
	maxOccuringCharCount := 0 // such names are against go's naming convention, but it improve readability.
	n := len(s)

	charIndex := func(c byte) int {
		return int(c - 'A')
	}

	for ; right < n; right++ { // using already defined right
		count[charIndex(s[right])]++
		maxOccuringCharCount = max(maxOccuringCharCount, count[charIndex(s[right])])
		for windowLen()-maxOccuringCharCount > k { // using windowLen func makes it easier to understand and improve readability.
			count[charIndex(s[left])]--
			left++
		}

		res = max(res, windowLen())
	}

	return res
}

func main() {
	s := readString()
	k := readInt()
	fmt.Println(longest_repeating_character_with_replacement(s, k))
}
