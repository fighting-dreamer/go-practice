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

func palindromic_substrings(s string) []string {
	if s == "" {
		return []string{}
	}
	res := map[string]bool{}
	l := len(s)
	dp := make([][]bool, l)
	for i := 0; i < l; i++ {
		dp[i] = make([]bool, l)
	}
	for i := 0; i < l; i++ {
		dp[i][i] = true
		res[s[i:i+1]] = true
	}
	for i := 0; i < l-1; i++ {
		if s[i] == s[i+1] {
			dp[i][i+1] = true
			res[s[i:i+2]] = true
		}
	}

	for i := 3; i <= l; i++ {
		for j := 0; j <= l-i; j++ {
			start := j
			end := start + i - 1

			if s[start] == s[end] && dp[start+1][end-1] == true {
				dp[start][end] = true
				res[s[start:end+1]] = true
			}
		}
	}
	ress := []string{}
	for k, _ := range res {
		ress = append(ress, k)
	}
	return ress
}

func main() {
	s := readString()
	ss := palindromic_substrings(s)
	fmt.Println(ss)
}
