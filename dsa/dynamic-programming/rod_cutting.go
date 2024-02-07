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

func rod_cutting(price []int, l int) int {
	// rod cutting it totally similar to unbounded knapsack
	// the wt is same as length
	// the price is same as profit
	// the cap is same as length here, it can change if the user want to give a rod of diff length to consider.
	wt := make([]int, l)
	for i := 0; i < l; i++ {
		wt[i] = i + 1
	}

	dp := make([][]int, l+1)
	for i := 0; i < l+1; i++ {
		dp[i] = make([]int, l+1)
	}
	dp[0][0] = 0
	for i := 1; i < l+1; i++ {
		dp[i][0] = 0
		dp[0][i] = 0
	}

	for i := 1; i < l+1; i++ {
		for j := 1; j < l+1; j++ {
			if wt[i-1] > j {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = max(price[i-1]+dp[i][j-wt[i-1]], dp[i-1][j])
			}
		}
	}

	return dp[l][l]
}

func main() {
	l := readInt()
	price := make([]int, l)
	for i := 0; i < l; i++ {
		price[i] = readInt()
	}

	fmt.Println(rod_cutting(price, l))
}
