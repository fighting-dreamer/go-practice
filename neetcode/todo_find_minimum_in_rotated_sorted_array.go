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

func find_minimum_in_rotated_sorted_array(nums []int, n int) int {
	// the array will start with a value that is going to be equal or more than the last index value.
	// ex : 4 5 6 7 0 1 2 3 ==> 4 > 3
	// all unique numbers.

	if len(nums) == 1 || nums[0] < nums[len(nums)-1] {
		return nums[0]
	}

	start := 0
	end := n - 1
	var mid int
	for start <= end {
		mid = start + (end-start)/2
		if nums[mid] >= nums[0] {
			// mid to end is sorted
			// our target lies in start to mid - 1
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

	fmt.Println(find_minimum_in_rotated_sorted_array(nums, n))
}
