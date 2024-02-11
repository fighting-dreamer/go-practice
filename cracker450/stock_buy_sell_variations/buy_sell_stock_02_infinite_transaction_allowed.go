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

func buy_sell_stock_02_infinite_transaction_allowed(nums []int, n int) int {
	// in this case, as we can buy and sel infinite or as many : we can buy and sell on same day too
	// this means, we can just check if the next day the price of stock is more => w can add that as profit too.
	// as we are able to add a postive difference.
	res := 0
	for i := 1; i < n; i++ {
		if nums[i]-nums[i-1] > 0 {
			res += nums[i] - nums[i-1]
		}
	}

	return res
}

/*
input :
7
100 180 260 310 40 535 695
output :
865
*/
func main() {
	n := readInt()
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = readInt()
	}

	fmt.Println(buy_sell_stock_02_infinite_transaction_allowed(nums, n))
}
