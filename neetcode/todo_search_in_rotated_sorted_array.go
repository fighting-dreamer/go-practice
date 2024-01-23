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

func search_in_rotated_sorted_array(nums []int, n int, target int) int {
	// return index or -1
	start := 0
	end := n - 1
	var mid int
	// invariant is the range between start and end is having the pivot element.
	for start <= end {
		mid = (start + end) / 2
		if nums[mid] >= nums[start] && nums[mid] >= nums[end] {
			// mid is in the first sorted section
			start = mid + 1
		} else {
			end = mid - 1
		}
	}

	return nums[start]
}

func main() {
	n := readInt()
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = readInt()
	}
	target := readInt()

	search_in_rotated_sorted_array(nums, n, target)

}
