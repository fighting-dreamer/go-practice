package main

import (
	"fmt"
	"log"
	"math"
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

func find_minimum_coins_for_coin_change(coins []int, n, s int) int {
	// we want to re-use the coins
	// we will either include certain coin or not
	// we will always want to choose the minimum of those 2 options

	dp := make([][]int, 1+n)
	for i := 0; i < 1+n; i++ {
		dp[i] = make([]int, 1+s)
	}
	for i := 1; i < 1+n; i++ {
		dp[i][0] = 0
	}
	for i := 0; i < 1+s; i++ {
		dp[0][i] = math.MaxInt - 1
	}
	for i := 1; i < 1+s; i++ {
		if i%coins[0] == 0 {
			dp[1][i] = i / coins[0]
		} else {
			dp[1][i] = math.MaxInt - 1
		}
	}

	for i := 2; i < 1+n; i++ {
		for j := 1; j < 1+s; j++ {
			if j >= coins[i-1] {
				dp[i][j] = min(1+dp[i][j-coins[i-1]], dp[i-1][j])
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}

	for i := 0; i < 1+n; i++ {
		for j := 0; j < 1+s; j++ {
			if dp[i][j] >= math.MaxInt-1 {
				dp[i][j] = -1
			}
			fmt.Print(dp[i][j], " ")
		}
		fmt.Println()
	}

	return dp[n][s]
}

func main() {
	n, s := readInt(), readInt()
	coins := make([]int, n)
	for i := 0; i < n; i++ {
		coins[i] = readInt()
	}

	fmt.Println(find_minimum_coins_for_coin_change(coins, n, s))
}
