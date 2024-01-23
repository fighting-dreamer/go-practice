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

/*
ADOBECODEBANC
ABC
*/

func minimum_window_substring(s, t string) string {
	res := ""
	resLeft, resRight := 0, len(s)-1
	resWindowLen := func() int {
		return resRight - resLeft + 1
	}

	left := 0
	right := 0
	windowLen := func() int {
		return right - left + 1
	}
	slen, tlen := len(s), len(t)
	have, need := 0, 0
	smap, tmap := map[byte]int{}, map[byte]int{}

	for i := 0; i < tlen; i++ {
		tmap[t[i]]++
	}
	need = len(tmap)

	for right < slen {
		_, ok := tmap[s[right]]
		if ok {
			if smap[s[right]]+1 == tmap[s[right]] {
				have++
			}
			smap[s[right]]++
		}

		if have < need {
			right++
		}

		if have == need {
			if resWindowLen() > windowLen() {
				res = s[left : right+1]
				resLeft = left
				resRight = right
				fmt.Println(res, resLeft, resRight, windowLen())
			}

			_, ok := tmap[s[left]]
			if ok {
				if smap[s[left]] == tmap[s[left]] {
					have--
				} else { // some issue in when to increment left if we have decremented have's value
					left++
				}
				smap[s[left]]--
			} else {
				left++
			}
		}
	}

	return res
}

func main() {
	s := readString()
	t := readString()

	fmt.Println(minimum_window_substring(s, t))
}
