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

func sum_max_non_adjacent(arr []int, n int) int {
	dp := make([]int, n)
	dp[0] = arr[0]
	dp[1] = arr[1]
	for i := 2; i < n; i++ {
		dp[i] = max(dp[i-1], dp[i-2]+arr[i])
	}

	return dp[n-1]
}

func main() {
	n := readInt()
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = readInt()
	}
	fmt.Println("sum_max_non_adjacent(arr, n) : ", sum_max_non_adjacent(arr, n))
}
