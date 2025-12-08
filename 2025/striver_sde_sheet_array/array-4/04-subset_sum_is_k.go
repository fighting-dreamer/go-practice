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

func longest_subarray_with_sum_k(arr []int, n int, k int) (int, int) {
	prefixSumIndex := map[int]int{}
	sumTillNow := 0
	res := 0
	start := 0
	end := 0

	prefixSumIndex[0] = -1
	for i := 0; i < n; i++ {
		sumTillNow += arr[i]
		if it, ok := prefixSumIndex[sumTillNow-k]; ok {
			if i-it > res {
				res = i - it
				start = it + 1 // sum till it is found to be same
				end = i
			}
		} else {
			prefixSumIndex[sumTillNow] = i
		}
	}
	fmt.Println(res)
	return start, end
}

func main() {
	n := readInt()
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = readInt()
	}

	targetSum := readInt()

	fmt.Println(longest_subarray_with_sum_k(arr, n, targetSum))

}
