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

func unbounded_knapsack(profit []int, wt []int, n int, cap int) int {
	dp := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, cap+1)
	}

	dp[0][0] = 0
	for i := 1; i < cap+1; i++ {
		dp[0][i] = 0
	}
	for i := 1; i < n+1; i++ {
		dp[i][0] = 0
	}

	for i := 1; i < n+1; i++ {
		for j := 1; j < cap+1; j++ {
			if wt[i-1] > j {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = max(profit[i-1]+dp[i][j-wt[i-1]], dp[i-1][j])
			}
		}
	}

	return dp[n][cap]
}

func main() {
	n, cap := readInt(), readInt()
	profit := make([]int, n)
	wt := make([]int, n)
	for i := 0; i < n; i++ {
		profit[i] = readInt()
	}
	for i := 0; i < n; i++ {
		wt[i] = readInt()
	}

	fmt.Println(unbounded_knapsack(profit, wt, n, cap))
}
