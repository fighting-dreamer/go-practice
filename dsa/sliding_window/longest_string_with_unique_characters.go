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

func longest_string_with_unique_chars(s string, k int) string {
	charMap := map[byte]int{}
	n := len(s)
	start, end := 0, -1 // -1 coz if there wa no real result, we still be putting the first char as result otherwise.
	resLen := 0
	uniqueCharInMap := 0

	for i, j := 0, 0; i < n && j < n; {
		charMap[s[j]]++
		if charMap[s[j]] == 1 {
			// it is a new unique char
			uniqueCharInMap++
		}

		if uniqueCharInMap < k {
			j++
		} else {
			if uniqueCharInMap == k {
				fmt.Println(s[i : j+1])
				if j-i+1 > resLen {
					resLen = j - i + 1
					start = i
					end = j
				}
				j++

			} else {
				// got a char that is makes the string not made of unique chars.
				for uniqueCharInMap > k {
					charMap[s[i]]--
					if charMap[s[i]] == 0 {
						uniqueCharInMap--
						delete(charMap, s[i])
					}
					i++
				}

				fmt.Println("-> ", s[i:j+1])
				if j-i+1 > resLen {
					resLen = j - i + 1
					start = i
					end = j
				}

				j++

			}
		}
	}

	return s[start : end+1]
}

func main() {
	s := readString()
	k := readInt()
	fmt.Println("RES : ", longest_string_with_unique_chars(s, k))
}
