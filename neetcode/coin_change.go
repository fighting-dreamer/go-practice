package main

import (
	"fmt"
	"log"
)

func coin_change(nums []int, n, t int) int {
	dp := make([][]int, n+1)

	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, t+1)
	}

	fmt.Println(len(dp), len(dp[0]))
	for i := 0; i < n+1; i++ {
		for j := 0; j < t+1; j++ {
			if j == 0 {
				dp[i][j] = 1
				continue
			}
			if i == 0 {
				dp[i][j] = 0
				continue
			}
			dp[i][j] = dp[i-1][j] // ways possible by not including ith coin for target sum j
			if j >= nums[i-1] {
				dp[i][j] += dp[i][j-nums[i-1]]
			}
		}
	}

	return dp[n][t]
}

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

func main() {
	n, t := readInt(), readInt()
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = readInt()
	}

	fmt.Println(coin_change(nums, n, t))
}
