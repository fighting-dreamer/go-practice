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

func buy_sell_stock_06_with_transaction_fee_infinite_transaction_allowed_tabulation(nums []int, n, fee int) int {
	// we have to pay transaction fee on one complete buy and sell
	// its like the effective benefit of a buy and sell is decreased by value of the transaction fee
	// we just incorporate as a cost factor and use it either on selling side or on buying side.(only one side)

	dp := make([][]int, 1+n)
	for i := 0; i < 1+n; i++ {
		dp[i] = make([]int, 2) // for buy or sell
	}

	for i := n - 1; i >= 0; i-- {
		for buy := 0; buy < 2; buy++ {
			if buy == 0 { // sell scenario
				selling_at_curr_index := dp[i+1][1] + (nums[i] - fee) // buy from next index
				selling_after_curr_index := dp[i+1][0]
				dp[i][0] = max(selling_after_curr_index, selling_at_curr_index)
			} else { // buy scenario
				buy_at_curr_index := dp[i+1][0] - nums[i] // sell from next index and pay the price of ith day.
				buy_after_curr_index := dp[i+1][1]
				dp[i][1] = max(buy_after_curr_index, buy_at_curr_index)
			}
		}
	}

	// the max we can make by buying(buy flag = 1) at or after 0th day
	return dp[0][1]
}

func main() {
	n, fee := readInt(), readInt()
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = readInt()
	}

	fmt.Println(buy_sell_stock_06_with_transaction_fee_infinite_transaction_allowed_tabulation(nums, n, fee))
}
