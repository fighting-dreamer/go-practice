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

func coin_change_2(amt int, coins []int, n int) int {
	dp := make([][]int, 1+n)
	for i := 0; i < 1+n; i++ {
		dp[i] = make([]int, 1+amt)
	}
	dp[0][0] = 1 // no coin and no amount => 1 way
	for i := 1; i < 1+n; i++ {
		dp[i][0] = 0 // some coins, no amt => no ways
	}
	for i := 1; i < 1+amt; i++ {
		dp[0][i] = 0 // no coins, some amt => no ways
	}

	for i := 1; i < 1+n; i++ {
		for j := 1; j < 1+amt; j++ {
			dp[i][j] = dp[i-1][j]
			if coins[i-1] <= j {
				dp[i][j] += dp[i][j-coins[i-1]]
			}
		}
	}

	return dp[n][amt]
}

func main() {
	amt := readInt()
	n := readInt()
	coins := make([]int, n)
	for i := 0; i < n; i++ {
		coins[i] = readInt()
	}

	coin_change_2(amt, coins, n)
}
