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

func find_maximum_number_of_ways_for_coin_change(nums []int, n, s int) int {
	// we are given some coins, and we want to know the maximum number of ways to make the sum s from those
	// we can use same coin multiple times
	// value of ith coin is nums[i]
	// ith coin could be included or not included
	// the ways would be = ways with ith coin and ways without ith coin.

	dp := make([][]int, 1+n)
	for i := 0; i < 1+n; i++ {
		dp[i] = make([]int, 1+s)
	}

	for i := 1; i < 1+n; i++ {
		dp[i][0] = 1
	}
	for i := 1; i < 1+s; i++ {
		dp[0][i] = 0
	}

	dp[0][0] = 1 // very important.
	for i := 1; i < 1+n; i++ {
		for j := 1; j < 1+s; j++ {
			if nums[i-1] <= j {
				waysWithOut_ithCoin := dp[i-1][j]
				// waysWith_ithCoin := dp[i][j-nums[i-1]] // we are allowing the same coin to be re-used multiple time.
				waysWith_ithCoin := dp[i-1][j-nums[i-1]] // we do not want to re-use the same coin multiple times
				dp[i][j] = waysWithOut_ithCoin + waysWith_ithCoin
			} else {
				waysWithOut_ithCoin := dp[i-1][j]
				dp[i][j] = waysWithOut_ithCoin
			}
		}
	}

	for i := 0; i < 1+n; i++ {
		for j := 0; j < 1+s; j++ {
			fmt.Print(dp[i][j], " ")
		}
		fmt.Println()
	}

	return dp[n][s]
}

func main() {
	n, s := readInt(), readInt()
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = readInt()
	}

	fmt.Println(find_maximum_number_of_ways_for_coin_change(nums, n, s))
}
