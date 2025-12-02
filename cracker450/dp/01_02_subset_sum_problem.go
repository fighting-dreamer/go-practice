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


func subset_sum(W int, coins []int, n int) bool {
	dp := make([][]bool, 1 + n)
	for i := 0; i < 1 + n; i++ {
		dp[i] = make([]bool, 1 + W)
	}

	for i := 0; i <= n; i++ {
		for j := 0; j <= W; j++ {
			if j == 0 {
				dp[i][j] = true
				continue
			}
			if i == 0 {
				dp[i][j] = false
				continue	
			}
			if j - coins[i - 1] >= 0 {
				dp[i][j] = dp[i - 1][j - coins[i - 1]] || dp[i - 1][j]
			}else {
				dp[i][j] = dp[i - 1][j]
			}
		}
	}

	return dp[n][W]
}

func subset_sum_recursive(S int, coins []int, n int, curr int, currSum int) bool {
	if currSum == 0 || curr == 0 {
		return true
	}

	left := subset_sum_recursive(S, coins, n, curr - 1, currSum)
	right := false
	if currSum > coins[curr] {
		right = subset_sum_recursive(S, coins, n, curr - 1, currSum - coins[curr])
	}

	return left || right
}

func main() {
	n := readInt()
	W := readInt()
	wt := make([]int, n)
	// val is same as wt

	for i := 0; i < n; i++ {
		wt[i] = readInt()
	}

	fmt.Printf("subset_sum(W, wt, n): %v\n", subset_sum(W, wt, n))
}
