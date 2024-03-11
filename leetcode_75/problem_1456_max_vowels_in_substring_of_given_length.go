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

func isVowel(ch byte) bool {
	switch {
	case ch == 'a':
		return true
	case ch == 'e':
		return true
	case ch == 'i':
		return true
	case ch == 'o':
		return true
	case ch == 'u':
		return true
	}
	return false
}

func find_substring_with_max_vowels(s string, k int) string {
	i, j := 0, 0
	n := len(s)
	vowelCount := 0
	maxVowelCount := 0
	res_l, res_r := 0, -1
	for i < n && j < n {
		if isVowel(s[j]) {
			vowelCount++
		}
		if j-i+1 < k {
			j++
		} else {
			if maxVowelCount < vowelCount {
				maxVowelCount = vowelCount
				res_l = i
				res_r = j
			}

			if isVowel(s[i]) {
				vowelCount--
			}
			i++
			j++
		}
	}

	return s[res_l : res_r+1]
}

func main() {
	s, k := readString(), readInt()
	// vowels = a,e,io,u
	// fmt.Println(isVowel("qwerty"[2]))
	fmt.Println(find_substring_with_max_vowels(s, k))
}
