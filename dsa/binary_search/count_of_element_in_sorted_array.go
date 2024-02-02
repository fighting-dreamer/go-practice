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

func lowerBound(nums []int, n, t int) int {
	start := 0
	end := n - 1
	for start <= end {
		mid := (start + end) / 2
		if nums[mid] == t {
			if mid == 0 || nums[mid-1] < nums[mid] {
				return mid
			}
			end = mid - 1
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

func upperBound(nums []int, n, t int) int {
	start := 0
	end := n - 1
	for start <= end {
		mid := (start + end) / 2
		if nums[mid] == t {
			if mid == n-1 || nums[mid+1] > nums[mid] {
				return mid
			}
			start = mid + 1
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

func count_elements_in_array(nums []int, n, t int) int {
	l := lowerBound(nums, n, t)
	u := upperBound(nums, n, t)
	fmt.Println(l, u)
	if l != -1 && u != -1 {
		return u - l + 1
	}

	return 0
}

func main() {
	n, t := readInt(), readInt()
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = readInt()
	}

	fmt.Println(count_elements_in_array(nums, n, t))
}
