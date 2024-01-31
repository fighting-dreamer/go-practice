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

func rain_water_trapping(nums []int, n int) int {
	leftMax := make([]int, n)
	rightMax := make([]int, n)

	leftMax[0] = 0
	rightMax[n-1] = 0

	for i := 1; i < n; i++ {
		leftMax[i] = max(leftMax[i-1], nums[i-1])
	}
	for i := n - 2; i >= 0; i-- {
		rightMax[i] = max(rightMax[i+1], nums[i+1])
	}

	res := 0
	for i := 0; i < n; i++ {
		m := min(leftMax[i], rightMax[i])
		if m > nums[i] {
			res += m - nums[i]
		}
	}

	return res
}

func main() {
	n := readInt()
	nums := make([]int, n)

	for i := 0; i < n; i++ {
		nums[i] = readInt()
	}

	fmt.Println(rain_water_trapping(nums, n))
}
