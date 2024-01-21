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

func product_of_array_except_self(nums []int, n int) []int {
	// O(n) solution
	output := make([]int, n)
	output[0] = 1
	for i := 1; i < n; i++ {
		output[i] = output[i-1] * nums[i-1]
	}
	product := 1
	for i := n - 1; i >= 0; i-- {
		output[i] = output[i] * product
		product = product * nums[i]
	}

	return output
}

func main() {
	n := readInt()
	nums := []int{}
	for i := 0; i < n; i++ {
		nums = append(nums, readInt())
	}
	output := product_of_array_except_self(nums, n)
	fmt.Println(output)
}
