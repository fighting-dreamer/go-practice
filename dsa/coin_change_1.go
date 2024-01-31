package main

import (
	"fmt"
	"log"
	"math"
)

func coin_change_1(nums []int, n, t int) int {
	// minimum number of coins required for target sum t
	if t == 0 {
		return 0
	}
	dp := make([]int, t+1)
	dp[0] = 0
	for i := 0; i < t+1; i++ {
		for j := 0; j < n; j++ {
			if nums[j] <= i {
				// coin's value  less than target
				if dp[i-nums[j]] != math.MaxInt && dp[i] > dp[i-nums[j]]+1 {
					dp[i] = 1 + dp[i-nums[j]]
				}
			}
		}
	}

	if dp[t] == math.MaxInt {
		return 0 // no way to get that target sum.
	}
	return dp[t]
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
	fmt.Println(coin_change_1(nums, n, t))
}
