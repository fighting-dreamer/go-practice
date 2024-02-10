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

func buy_sell_stock_01_only_one_transaction_allowed(nums []int, n int) int {
	// find the max till now from last and keep track ofthe max diff we can found.
	maxTillNow := nums[n-1]
	res := 0
	for i := n - 2; i >= 0; i-- {
		maxTillNow = max(maxTillNow, nums[i])
		res = max(res, maxTillNow-nums[i])
	}

	return res
}

func main() {
	// we are allowed to make only one transaction, we have to maximize the profit.
	n := readInt()
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = readInt()
	}
	fmt.Println(buy_sell_stock_01_only_one_transaction_allowed(nums, n))
}
