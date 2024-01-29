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

func longest_palindromic_substring(s string) string {
	if s == "" {
		// return 0, ""
		return ""
	}

	l := len(s)
	dp := make([][]bool, l)
	for i := 0; i < l; i++ {
		dp[i] = make([]bool, l)
	}
	for i := 0; i < l; i++ {
		dp[i][i] = true
	}
	// maxL := 1
	resStart, resEnd := 0, 0
	for i := 0; i < l-1; i++ {
		if s[i] == s[i+1] {
			dp[i][i+1] = true
			// maxL = 2
			resStart = i
			resEnd = i + 1
		}
	}

	// starting with length of 3 to length of whole input string
	for ll := 3; ll <= l; ll++ {
		for i := 0; i <= l-ll; i++ {
			start := i
			end := start + ll - 1
			if s[start] == s[end] && dp[start+1][end-1] == true {
				dp[start][end] = true
				// maxL = ll
				resStart = start
				resEnd = end
			}
		}
	}
	// fmt.Println(resStart, resEnd)
	// return maxL, s[resStart : resEnd+1]
	return s[resStart : resEnd+1]
}

func main() {
	s := readString()
	fmt.Println(longest_palindromic_substring(s))
}
