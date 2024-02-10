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

func find_max_by_selling_till_now(nums []int, n int) []int {
	sellMaxTillNow := make([]int, n)
	sellMaxTillNow[0] = 0
	minTillNow := nums[0]
	for i := 1; i < n; i++ {
		minTillNow = min(minTillNow, nums[i])

		if nums[i]-minTillNow > sellMaxTillNow[i-1] {
			sellMaxTillNow[i] = nums[i] - minTillNow
		} else {
			sellMaxTillNow[i] = sellMaxTillNow[i-1]
		}
	}
	return sellMaxTillNow
}

func find_max_by_buying_from_now(nums []int, n int) []int {
	buyMaxFromNow := make([]int, n)
	buyMaxFromNow[n-1] = 0
	maxFromNow := nums[n-1]
	for i := n - 2; i >= 0; i-- {
		maxFromNow = max(maxFromNow, nums[i])

		if maxFromNow-nums[i] > buyMaxFromNow[i+1] {
			buyMaxFromNow[i] = maxFromNow - nums[i]
		} else {
			buyMaxFromNow[i] = buyMaxFromNow[i+1]
		}
	}

	return buyMaxFromNow
}

func buy_sell_stock_03_two_transaction_allowed(nums []int, n int) int {
	// at most two transaction => there can be no transaction, one transaction, two transaction.
	// find transaction with max cost if would have to sell by today or today.
	// find the transaction with max cost if would have bought today to after.

	sellTillNow := find_max_by_selling_till_now(nums, n)
	buyFromNow := find_max_by_buying_from_now(nums, n)

	// fmt.Println(sellTillNow)
	// fmt.Println(buyFromNow)

	res := 0
	for i := 0; i < n; i++ {
		res = max(res, sellTillNow[i]+buyFromNow[i])
	}

	return res
}

func main() {
	// at most two transaction => there can be no transaction, one transaction, two transaction.
	n := readInt()
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = readInt()
	}

	fmt.Println(buy_sell_stock_03_two_transaction_allowed(nums, n))
}
