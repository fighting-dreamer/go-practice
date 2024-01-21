package main

import (
	"fmt"
	"log"
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

func two_sum(nums []int, n int, target int) (int, int) {
	start := 0
	end := n - 1
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	for {
		if start == end {
			break
		}
		sum := nums[start] + nums[end]
		if sum == target {
			return nums[start], nums[end]
		} else {
			if sum < target {
				start++
			} else {
				end--
			}
		}
	}
	return -1, -1
}

func main() {
	n := readInt()
	nums := []int{}
	for i := 0; i < n; i++ {
		nums = append(nums, readInt())
	}

	target := readInt()
	target_pair_a, target_pair_b := two_sum(nums, n, target)
	fmt.Println(target_pair_a, target_pair_b)
}
