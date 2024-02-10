package main

import (
	"fmt"
	"log"
)

// https://www.geeksforgeeks.org/swap-all-occurrences-of-two-characters-to-get-lexicographically-smallest-string/

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

func compare_swap(s string) string {
	ch := make([]int, 26)
	for i := 0; i < 26; i++ {
		ch[i] = -1
	}

	for i := 0; i < len(s); i++ {
		if ch[s[i]-'a'] == -1 {
			ch[s[i]-'a'] = i
		}
	}
	// fmt.Println(ch)

	var i, j int
	var flag int
	for i = 0; i < len(s); i++ {
		flag = 0
		for j = 0; j < int(s[i]-'a'); j++ {
			if ch[j] > ch[s[i]-'a'] {
				flag = 1
				break
			}
		}
		if flag == 1 {
			break
		}
	}
	// fmt.Println(i, j)
	if i < len(s)-1 {
		// can be used for swapping
		ch_i := byte(s[i])
		ch_j := byte(j + 'a')

		for k := 0; k < len(s); k++ {
			if s[k] == ch_j {
				s = s[:k] + string(ch_i) + s[k+1:]
			} else if s[k] == ch_i {
				s = s[:k] + string(ch_j) + s[k+1:]
			}
		}
	}

	return s
}

func main() {
	s := readString()

	fmt.Println(compare_swap(s))
}
