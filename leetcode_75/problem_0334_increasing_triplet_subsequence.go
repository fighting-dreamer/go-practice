package main

import (
	"fmt"
	"log"
	"math"
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

func check_three_increasing_triplets_exists(nums []int) bool {
	n := len(nums)
	first, second := math.MaxInt, math.MaxInt
	for i := 0; i < n; i++ {
		if first > nums[i] {
			first = nums[i]
		}
		if nums[i] > first {
			second = min(second, nums[i])
		}

		if nums[i] > second {
			fmt.Println(first, second, nums[i])
			return true
		}
	}

	return false
}

func main() {
	n := readInt()
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = readInt()
	}

	fmt.Println(check_three_increasing_triplets_exists(nums))
}
