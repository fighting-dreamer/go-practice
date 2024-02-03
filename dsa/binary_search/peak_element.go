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

func is_peak(nums []int, mid, n int) bool {
	if n == 1 {
		return true
	}

	if mid == 0 {
		return nums[0] > nums[1]
	}
	if mid == n-1 {
		return nums[n-2] < nums[n-1]
	}

	if nums[mid-1] < nums[mid] && nums[mid+1] < nums[mid] {
		return true
	}

	return false
}

func peak_element(nums []int, n int) int {
	start := 0
	end := n - 1
	for start <= end {
		fmt.Println(start, end)
		mid := (start + end) / 2
		if is_peak(nums, mid, n) {
			return nums[mid]
		} else {
			if mid == 0 {
				start = mid + 1
				continue
			}
			if mid == n-1 {
				end = mid - 1
				continue
			}
			if nums[mid] < nums[mid+1] {
				start = mid + 1
			} else {
				end = mid - 1
			}
		}
	}

	return -1
}

func main() {
	n := readInt()
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = readInt()
	}

	fmt.Println(peak_element(nums, n))
}
