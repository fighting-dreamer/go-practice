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

// only one buy and sell allowed: maximise profit
func stock_buy_sell(arr []int, n int) int {
	max_to_right := make([]int, n)
	max_to_right[n-1] = -1
	for i := n - 2; i >= 0; i-- {
		max_to_right[i] = max(max_to_right[i+1], arr[i+1])
	}
	res := 0
	for i := 0; i < n; i++ {
		res = max(res, max_to_right[i]-arr[i])
	}

	return res
}

func print_array(arr []int, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", arr[i])
	}
	fmt.Println()
}

func main() {
	n := readInt()
	arr := make([]int, n)

	for i := 0; i < n; i++ {
		arr[i] = readInt()
	}
	print_array(arr, n)
	fmt.Printf("stock_buy_sell(arr, n): %v\n", stock_buy_sell(arr, n))
}
