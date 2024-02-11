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

func buy_sell_05_memoized_helper(nums []int, n int, index int, isBuyFlag int, dp [][]int) int {
	if index >= n {
		return 0
	}

	if dp[index][isBuyFlag] != -1 {
		return dp[index][isBuyFlag]
	}
	// isBuyFlag == 1 => buy operation is to be done.
	if isBuyFlag == 1 {
		//case : we don;t buy at curr index => check if bought from next index
		not_buy_at_curr_index := buy_sell_05_memoized_helper(nums, n, index+1, isBuyFlag, dp)
		// case : we buy and spent money = nums[index] and now we want to sell
		buy_at_curr_index := buy_sell_05_memoized_helper(nums, n, index+1, 0, dp) - nums[index]
		dp[index][isBuyFlag] = max(not_buy_at_curr_index, buy_at_curr_index)
	} else {
		//case : we don;t sell at curr index => check if sell from next index
		not_sell_at_curr_index := buy_sell_05_memoized_helper(nums, n, index+1, isBuyFlag, dp)
		// case : we sell and got money = nums[index] and now we want to buy
		// coz of cooldown, we skip just next index and move to check from "index + 2"
		sell_at_curr_index := buy_sell_05_memoized_helper(nums, n, index+2, 1, dp) + nums[index]
		dp[index][isBuyFlag] = max(not_sell_at_curr_index, sell_at_curr_index)
	}

	return dp[index][isBuyFlag]
}

func buy_sell_stock_05_with_cooldown_infinite_transaction_allowed_memoized(nums []int, n int) int {
	dp := make([][]int, 2+n)
	for i := 0; i < 2+n; i++ {
		dp[i] = []int{-1, -1} // for buy and sell(not buy) states
	}

	buy := 1
	return buy_sell_05_memoized_helper(nums, n, 0, buy, dp)
}

func buy_sell_stock_05_with_cooldown_infinite_transaction_allowed_tabulation(nums []int, n int) int {
	dp := make([][]int, 2+n)
	for i := 0; i < 2+n; i++ {
		dp[i] = make([]int, 2)
	}

	for i := n - 1; i >= 0; i-- {
		for sell := 0; sell <= 1; sell++ {
			profit := 0
			if sell == 0 { // sell == 0 => buy scenario
				profit = max(0+dp[i+1][0], dp[i+1][1]-nums[i])
			}
			if sell == 1 { // sell == 1 => sell scenario
				profit = max(0+dp[i+1][1], dp[i+2][0]+nums[i])
			}
			dp[i][sell] = profit
		}
	}

	return dp[0][0]
}

func buy_sell_stock_05_with_cooldown_infinite_transaction_allowed(nums []int, n int) int {
	res_memoized := buy_sell_stock_05_with_cooldown_infinite_transaction_allowed_memoized(nums, n)
	res_tabulation := buy_sell_stock_05_with_cooldown_infinite_transaction_allowed_tabulation(nums, n)
	fmt.Println("memoized : ", res_memoized)
	fmt.Println("tabulation : ", res_tabulation)
	return res_memoized
}

func main() {
	n := readInt()
	nums := make([]int, n)
	// cooldown is after selling and of 1 day.
	for i := 0; i < n; i++ {
		nums[i] = readInt()
	}
	fmt.Println(buy_sell_stock_05_with_cooldown_infinite_transaction_allowed(nums, n))
}
