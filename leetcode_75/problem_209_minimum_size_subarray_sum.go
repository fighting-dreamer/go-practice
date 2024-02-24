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

func minimum_size_subarray_sum(nums []int, n, target int) []int {
	i, j := 0, 0
	temp := 0
	res_i, res_j, resSize := 0, -1, n+1
	for i < n && j < n {
		temp += nums[j]
		fmt.Println(i, j, temp)
		if temp < target {
			j++
		} else if temp == target {
			currSize := j - i + 1
			if resSize > currSize {
				resSize = currSize
				res_i, res_j = i, j
			}
			temp -= nums[i]
			i++
			j++
		} else if temp > target {
			for temp > target {
				temp -= nums[i]
				i++
			}

			if temp == target {
				currSize := j - i + 1
				if resSize > currSize {
					resSize = currSize
					res_i, res_j = i, j
				}
			}
			j++
		}
	}

	fmt.Println(res_i, res_j, resSize)

	return nums[res_i : res_j+1]
}

func main() {
	n := readInt()
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = readInt()
	}

	target := readInt()

	fmt.Println(minimum_size_subarray_sum(nums, n, target))
}
