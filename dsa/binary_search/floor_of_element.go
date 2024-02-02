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

func floor_of_element(nums []int, n, t int) int {
	// floor is largest element available that is element smaller than t
	// find the element smaller than t(say res) and if next element is also smaller but greater than res => keep looking to right
	fmt.Println("searching for floor of ", t)
	start := 0
	end := n - 1
	var res int
	for start <= end {
		mid := (start + end) / 2
		if nums[mid] < t {
			res = nums[mid]
			if mid+1 != n && nums[mid+1] < t {
				start = mid + 1
				continue
			}
			return res
		} else {
			end = mid - 1
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

	fmt.Println(floor_of_element(nums, n, t))
}
