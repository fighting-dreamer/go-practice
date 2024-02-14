package main

import (
	"fmt"
	"log"
	"slices"
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

func isOdd(i int) bool {
	return i&1 == 1
}

func maximum_sum_after_k_negations(nums []int, n, k int) int {
	slices.Sort(nums)
	var i int
	for i = 0; i < k && nums[i] < 0; i++ {
		nums[i] = -nums[i]
	} // we can turn same element -ve multiple times.
	if isOdd(k - i) {
		nums[i-1] = -nums[i-1]
	}
	s := 0
	for i := 0; i < n; i++ {
		s += nums[i]
	}

	return s
}

func main() {
	n, k := readInt(), readInt()
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = readInt()
	}

	fmt.Println(maximum_sum_after_k_negations(arr, n, k))

	return
}
