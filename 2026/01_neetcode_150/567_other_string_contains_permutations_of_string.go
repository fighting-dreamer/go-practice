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

func check_s2_contains_s1_permutation(s1 string, s2 string) bool {
	n1 := len(s1)
	chMap1 := map[byte]int{}
	for i := 0; i < len(s1); i++ {
		chMap1[s1[i]]++
	}

	n2 := len(s2)
	chMap2 := map[byte]int{}
	for i, j := 0, 0; i < n2 && j < n2; {
		chMap2[s2[j]]++
		if j-i+1 < n1 {
			j++
		} else if j-i+1 == n1 {
			// compare char maps
			if len(chMap1) == len(chMap2) {
				flag := true
				for k, v := range chMap1 {
					if chMap2[k] != v {
						flag = false
						break
					}
				}
				if flag {
					return true // we have a permutation of s1 in s2
				}
			}

			chMap2[s2[i]]--
			if chMap2[s2[i]] == 0 {
				delete(chMap2, s2[i])
			}

			i++
			j++
		}
	}

	return false
}

func main() {
	s1 := readString()
	s2 := readString()

	fmt.Printf("check_s2_contains_s1_permutation(s1, s2): %v\n", check_s2_contains_s1_permutation(s1, s2))
}
