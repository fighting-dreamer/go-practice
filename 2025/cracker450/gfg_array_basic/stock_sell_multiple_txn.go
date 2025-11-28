package main

import (
	"fmt"
	"log"
)

func stock_sell_multiple_txn(prices []int, n int) int {
	res := 0
	for i := 1; i < n; i++ {
		temp := prices[i] - prices[i-1]
		if temp > 0 {
			res += temp
		}
	}
	return res
}

func print_array(arr []int, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", arr[i])
	}
	fmt.Println()
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
	n := readInt()
	prices := make([]int, n)
	for i := 0; i < n; i++ {
		prices[i] = readInt()
	}

	print_array(prices, n)
	fmt.Println("Max profit : ", stock_sell_multiple_txn(prices, n))

}
