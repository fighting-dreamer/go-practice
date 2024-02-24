package main

import (
	"fmt"
	"log"
)

func minimum_window_substring(s, t string) string {
	// check the minimum length sub-string which have all characters of `t` in it.
	t_charMap := map[byte]int{}
	ls, lt := len(s), len(t)
	for i := 0; i < lt; i++ {
		t_charMap[t[i]]++
	}
	s_charMap := map[byte]int{}
	remaining := len(t_charMap)
	i, j := 0, 0
	res_i, res_j, resLen := 0, 0, ls+1

	for i < ls && j < ls {
		if _, ok := t_charMap[s[j]]; ok {
			s_charMap[s[j]]++
			if s_charMap[s[j]] == t_charMap[s[j]] {
				remaining--
			}
		}

		for remaining == 0 && i <= j {
			// we have achieved a window which have all the required characters atleast.
			if j-i+1 < resLen {
				res_i, res_j = i, j
				resLen = j - i + 1
			}

			if _, ok := t_charMap[s[i]]; ok {
				if s_charMap[s[i]] == t_charMap[s[i]] {
					remaining++
				}
				s_charMap[s[i]]--
			}
			i++

		}
		j++
	}

	return s[res_i : res_j+1]
}

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

func main() {
	s, t := readString(), readString()
	// t to be part of s
	fmt.Println(minimum_window_substring(s, t))
}
