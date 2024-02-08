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

func sub_set_sum(set []int, n, s int) bool {
	dp := make([][]bool, n+1)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, s+1)
	}
	dp[0][0] = true
	for i := 1; i < n+1; i++ {
		dp[i][0] = true
	}
	for i := 1; i < s+1; i++ {
		dp[0][i] = false
	}

	for i := 1; i < n+1; i++ {
		for j := 1; j < s+1; j++ {
			if set[i-1] <= j {
				dp[i][j] = dp[i-1][j-set[i-1]] || dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}

	return dp[n][s]
}

func main() {
	n, s := readInt(), readInt()
	set := make([]int, n)
	for i := 0; i < n; i++ {
		set[i] = readInt()
	}

	fmt.Println(sub_set_sum(set, n, s))
}
