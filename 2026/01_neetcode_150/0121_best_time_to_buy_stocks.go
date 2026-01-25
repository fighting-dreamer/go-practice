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

// return hte max profit from one txn (buy and sell operation)
func best_time_to_buy_and_sell_stock(arr []int, n int) int {
	rightMax := arr[n-1]
	res := 0

	for i := n - 2; i >= 0; i-- {
		res = max(res, rightMax-arr[i])
		rightMax = max(rightMax, arr[i])
	}

	return res
}

func main() {
	n := readInt()
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = readInt()
	}

	best_time_to_buy_and_Sell_stock(arr, n)
}
