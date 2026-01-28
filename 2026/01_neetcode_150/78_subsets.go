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

func power_set_iterative(nums []int, n int) [][]int {
	res := [][]int{}
	for i := 0; i < (1 << n); i++ {
		curr := []int{}
		for j := 0; j < n; j++ {
			if i&(1<<j) > 0 {
				curr = append(curr, nums[j])
			}
		}
		res = append(res, append([]int{}, curr...))
	}

	return res
}

func power_set_helper(nums []int, n int, curr int, path []int, res *[][]int) {
	if curr == n {
		*res = append(*res, path)
		return
	}

	power_set_helper(nums, n, curr+1, path, res)
	power_set_helper(nums, n, curr+1, append(path, nums[curr]), res)
}

func power_set_recursive(nums []int, n int) [][]int {
	res := [][]int{}
	power_set_helper(nums, n, 0, []int{}, &res)
	return res
}

func main() {
	n := readInt()
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = readInt()
	}

	fmt.Printf("power_set_iterative(nums, n): %v\n", power_set_iterative(nums, n))
	fmt.Printf("power_set_recursive(nums, n): %v\n", power_set_recursive(nums, n))
}
