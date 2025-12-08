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

type Result struct {
	len        int
	startIndex int
	endIndex   int
	val        string
}

// consider if i is at index 0 and j at index 2 for string "ababa", the condition str[i] != str[j] terminate the loop before it start.
func problematic_longest_substring_with_no_repeating_char(str string) Result {
	chMp := map[byte]int{}
	// for j := 0; j < n; j++ {
	// 	if chMp[str[j]] is seen before {
	// 		check if length of non-repeating is possibly longest
	// 		if its longests, update result

	// 		move and remove entry for elements till you encounter the occurence of str[j] using index i
	// 	}
	// 	for str[j], add entry in map
	// }
	// // cehck if last segment from i-j is of max length

	j := 0
	i := 0
	n := len(str)
	res := 0
	var startIndex, endIndex int
	for ; j < n; j++ {
		if chMp[str[j]] == 1 {
			if j-i > res {
				res = j - i
				startIndex = i
				endIndex = j
			}
			for ; str[i] != str[j]; i++ {
				delete(chMp, str[i])
			}
		}
		chMp[str[j]] = 1
	}

	if j-i > res {
		res = j - i
		startIndex = i
		endIndex = j
	}

	return Result{
		len:        res,
		startIndex: startIndex,
		endIndex:   endIndex,
		val:        str[startIndex:endIndex],
	}
}

func longest_substring_with_no_repeating_char(str string) Result {
	lastSeen := map[byte]int{}
	j := 0
	i := 0
	n := len(str)
	res := 0
	startIndex := 0
	endIndex := 0

	for ; j < n; j++ {
		if found, ok := lastSeen[str[j]]; ok && found >= i {
			if j-i > res {
				res = j - i
				startIndex = i
				endIndex = j
			}

			// for ; i <= found; i++ {
			// 	delete(lastSeen, str[i])
			// } // we can just check "if found >= i" => ignoring values if found index is below i
			i = found + 1
		}
		lastSeen[str[j]] = j
	}

	if j-i > res {
		res = j - i
		startIndex = i
		endIndex = j
	}
	return Result{
		len:        res,
		startIndex: startIndex,
		endIndex:   endIndex,
		val:        str[startIndex:endIndex],
	}
}

func main() {
	str := readString()
	res := longest_substring_with_no_repeating_char(str)
	fmt.Printf("res: %v\n", res)
}
