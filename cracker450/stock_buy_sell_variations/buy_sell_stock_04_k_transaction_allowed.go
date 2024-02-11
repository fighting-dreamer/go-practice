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

func buy_sell_stock_04_at_most_k_transaction_allowed(nums []int, n, k int) int {
	dp := make([][]int, 1+k) // the max we can get for i number of transaction allowed in j number of days.
	for i := 0; i < 1+k; i++ {
		dp[i] = make([]int, 1+n)
	}

	// there is a dp[1+k][1+n] matrix.
	// for k = 0 : the dp[0][i..] value is 0
	// for n = 0 : the dp[i..][0] value is 0
	for i := 0; i < 1+n; i++ {
		dp[0][i] = 0
	}
	for i := 0; i < 1+k; i++ {
		dp[i][0] = 0
	}

	for i := 1; i < 1+k; i++ {
		for j := 1; j < 1+n; j++ {
			dp[i][j] = dp[i][j-1] // max sum till i transactions in j-1 days.
			for d := 0; d < j; d++ {
				// dp[i-1][d] => transaction one less than i and till dth day.
				dp[i][j] = max(dp[i][j], nums[j-1]-nums[d]+dp[i-1][d])
			}
		}
	}

	return dp[k][n]
}

func main() {
	n, k := readInt(), readInt()

	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = readInt()
	}

	fmt.Println(buy_sell_stock_04_at_most_k_transaction_allowed(nums, n, k))
}
