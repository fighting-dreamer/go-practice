package main

import "fmt"

func jump_game_2(arr []int, n int) int {
	dp := make([]int, n)
	dp[0] = 0
	for i := 1; i < n; i++ {
		dp[i] = n + 1 // large value
	}

	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if j+arr[j] >= i {
				dp[i] = min(dp[i], dp[j]+1)
			}
		}
	}

	return dp[n-1]
}

func main() {
	n := readInt()
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = readInt()
	}

	fmt.Printf("jump_game_2(arr, n): %v\n", jump_game_2(arr, n))
}
