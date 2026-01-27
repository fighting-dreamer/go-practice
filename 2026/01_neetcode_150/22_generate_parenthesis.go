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

func generate_parenthesis_dfs(current string, open int, close int, max int, res *[]string) {
	if len(current) == 2*max {
		*res = append(*res, current)
		return
	}
	if open < max {
		generate_parenthesis_dfs(current+"(", open+1, close, max, res)
	}
	if close < open {
		generate_parenthesis_dfs(current+")", open, close+1, max, res)
	}
}

func generate_parenthesis(n int) []string {
	res := []string{}

	generate_parenthesis_dfs("", 0, 0, n, &res)

	return res
}

func main() {
	n := readInt()
	fmt.Printf("generate_parenthesis(n): %v\n", generate_parenthesis(n))
}
