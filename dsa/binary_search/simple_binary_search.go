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

func sort_using_go(nums []int) {
	fmt.Println("Before sorting : ", nums)
	slices.Sort(nums)
	fmt.Println("After sorting : ", nums)
}

func implement_binary_search(nums []int, n, t int) int {
	start := 0
	end := n - 1
	var mid int
	for start <= end {
		mid = (start + end) / 2
		if nums[mid] == t {
			return mid
		} else {
			if nums[mid] > t {
				end = mid - 1
			} else {
				start = mid + 1
			}
		}
	}

	return -1
}

func binary_search(nums []int, n, t int) int {
	// index where t is present
	// using golang : sort
	// Works : //position := sort.SearchInts(nums, t) // return index if found or length of array if not found.

	// works :
	// position, ok := slices.BinarySearch(nums, t)
	// if !ok {
	// 	return -1
	// }
	position := implement_binary_search(nums, n, t)
	return position
}

func main() {
	n, t := readInt(), readInt()
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = readInt()
	}

	fmt.Println("main Before sorting : ", nums)
	sort_using_go(nums) // works and inplace work.
	fmt.Println("main After sorting : ", nums)

	fmt.Println(binary_search(nums, n, t))
}
