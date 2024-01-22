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

func best_time_to_buy_and_sell_stock_one_time(nums []int, n int) int {
	rightMax := nums[n-1]
	res := 0
	for i := n - 2; i >= 0; i-- {
		res = max(res, rightMax-nums[i])
		rightMax = max(rightMax, nums[i])
	}

	return res
}

func main() {
	n := readInt()
	nums := []int{}
	for i := 0; i < n; i++ {
		nums = append(nums, readInt())
	}

	fmt.Println(best_time_to_buy_and_sell_stock_one_time(nums, n))
}
