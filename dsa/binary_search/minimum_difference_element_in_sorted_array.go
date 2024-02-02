package main

import (
	"fmt"
	"log"
	"math"
)

// star

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

func find_minimum_difference_element(nums []int, n, t int) int {
	start := 0
	end := n - 1
	for start <= end {
		mid := (start + end) / 2
		if nums[mid] == t {
			fmt.Println("----> ", start, end)
			return mid
		} else {
			if nums[mid] < t {
				start = mid + 1
			} else {
				end = mid - 1
			}
		}
	}
	// STAR : When we get out of the loop, the start and end are the next elements to the key we were trying to find
	// example : 2 4 6 8 10 11 13, w eare finding 12 in this, the start will point to 13 and end to 11(next nearest elements to key)
	fmt.Println(start, end)
	if start < n && start >= 0 && end < n && end >= 0 {
		fmt.Println(nums[start], nums[end])
	}

	var s, e int
	if start < n && start >= 0 {
		s = int(math.Abs(float64(nums[start] - t)))
	} else {
		s = math.MaxInt
	}
	if end < n && end >= 0 {
		e = int(math.Abs(float64(nums[end] - t)))
	} else {
		e = math.MaxInt
	}
	if s < e {
		return start
	} else {
		return end
	}
}

func main() {
	n, t := readInt(), readInt()
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = readInt()
	}

	x := find_minimum_difference_element(nums, n, t)
	fmt.Println(x, " : ", nums[x])
}
