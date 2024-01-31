package main

import (
	"fmt"
	"log"
	"math"
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

func min_max_subarray_sum_of_size_k(nums []int, n, k int) (int, int) {
	if k > n {
		return 0, 0
	}

	ssum := 0
	for i := 0; i < k; i++ {
		ssum += nums[i]
	}
	maxSum, minSum := ssum, ssum
	for i := k; i < n; i++ {
		ssum = ssum + nums[i] - nums[i-k]
		maxSum = max(maxSum, ssum)
		minSum = min(minSum, ssum)
	}

	return maxSum, minSum
}

func min_max_subarray_sum_of_size_k_sliding_window(nums []int, n, k int) (int, int) {
	if k > n {
		return 0, 0
	}

	ssum := 0
	maxSum := 0
	minSum := math.MaxInt
	for i, j := 0, 0; i < n && j < n; {
		ssum = ssum + nums[j]
		if j-i+1 < k {
			j++
		} else if j-i+1 == k {
			maxSum = max(maxSum, ssum)
			minSum = min(minSum, ssum)
			ssum = ssum - nums[i]
			// maintain window size :

			i++
			j++
		}
	}

	return maxSum, minSum
}

func main() {
	n, k := readInt(), readInt()
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = readInt()
	}
	fmt.Println(min_max_subarray_sum_of_size_k(nums, n, k))
	fmt.Println(min_max_subarray_sum_of_size_k_sliding_window(nums, n, k))
}
