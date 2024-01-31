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

func check_anagram(tm, sm map[byte]int) bool {
	fmt.Println("check : ")
	fmt.Println(sm, " : ", tm)
	if !(len(tm) == len(sm)) {
		return false
	}
	for k, v := range tm {
		if v != sm[k] {
			return false
		}
	}

	return true
}

func all_anagrams(s, t string) []string {
	res := []string{}
	k := len(t)
	chmt := map[byte]int{} // char map for string t
	for i := 0; i < k; i++ {
		chmt[t[i]]++
	}
	fmt.Println("chmt : ", chmt)
	chms := map[byte]int{} // char map for string s
	n := len(s)
	for i, j := 0, 0; i < n && j < n; {
		chms[s[j]]++
		if j-i+1 < k {
			j++
		} else if j-i+1 == k {
			if check_anagram(chmt, chms) {
				res = append(res, s[i:j+1])
			}
			if chms[s[i]]--; chms[s[i]] == 0 {
				delete(chms, s[i])
			}

			i++
			j++
		}
	}

	return res
}

func main() {
	s, t := readString(), readString()

	fmt.Println(all_anagrams(s, t))
}
