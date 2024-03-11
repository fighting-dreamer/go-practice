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

func find_consecutive_ones_iii(nums []int, k int) int {
	n := len(nums)
	i, j := 0, 0
	maxRes := 0
	currAvaliableZerosCount := k
	currOnes := 0

	for i < n && j < n {
		if nums[j] == 1 {
			currOnes++
			j++
		} else {
			if currAvaliableZerosCount > 0 {
				currAvaliableZerosCount--
				currOnes++
				j++
			} else {
				for nums[i] != 0 {
					currOnes--
					i++
				}

				currAvaliableZerosCount++
				currOnes--
				i++
			}
		}
		maxRes = max(maxRes, currOnes)
	}

	return maxRes
}

func main() {
	n, k := readInt(), readInt()
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = readInt()
	}

	fmt.Println(find_consecutive_ones_iii(nums, k))
}
