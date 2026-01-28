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

func subset_2_helper(nums []int, n int, curr int, path []int, res *[][]int) {
	if curr == n {
		*res = append(*res, append([]int{}, path...))
		return
	}

	subset_2_helper(nums, n, curr+1, append(path, nums[curr]), res)
	for curr+1 < n && nums[curr] == nums[curr+1] {
		curr++
	}
	subset_2_helper(nums, n, curr+1, path, res)
}

func subsets_2(nums []int, n int) [][]int {
	res := [][]int{}
	slices.Sort(nums)
	subset_2_helper(nums, n, 0, []int{}, &res)

	return res
}

func main() {
	n := readInt()
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = readInt()
	}

	fmt.Printf("subsets_2(arr, n): %v\n", subsets_2(arr, n))
}
