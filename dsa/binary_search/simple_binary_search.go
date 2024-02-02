package main

import (
	"fmt"
	"log"
	"slices"
	"sort"
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

func reverse_sorted_binary_search(nums []int, n, t int) int {
	start := 0
	end := n - 1
	for start <= end {
		mid := (start + end) / 2
		if nums[mid] == t {
			return mid
		} else {
			if nums[mid] > t {
				start = mid + 1
			} else {
				end = mid - 1
			}
		}
	}

	return -1
}

func main() {
	n, t := readInt(), readInt()
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = readInt()
	}
	fmt.Println(slices.IsSorted(nums)) // is not descending => next element is always greater or equal : example 1 2 2 3 4 is always greater or equal.

	fmt.Println("main Before sorting : ", nums)
	sort_using_go(nums) // works and inplace work.
	fmt.Println("main After sorting : ", nums)

	fmt.Println(binary_search(nums, n, t))

	fmt.Println("--------------REVERSE-----------------")
	fmt.Println(nums)
	sort.Slice(nums, func(a, b int) bool {
		return nums[a] > nums[b]
	})
	fmt.Println(nums)

	fmt.Println(reverse_sorted_binary_search(nums, n, t))
}
