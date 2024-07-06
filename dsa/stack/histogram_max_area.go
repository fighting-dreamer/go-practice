package main

import (
	"fmt"
	"log"

	"github.com/liyue201/gostl/ds/stack"
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

func nearest_smaller_to_left(nums []int, n int) []int {
	st := stack.New[int]()
	res := make([]int, n)

	res[0] = -1
	st.Push(0)
	for i := 1; i < n; i++ {
		for !st.Empty() && nums[st.Top()] >= nums[i] {
			st.Pop()
		}
		if st.Empty() {
			res[i] = -1
		} else {
			res[i] = st.Top()
		}
		st.Push(i)
	}

	return res
}

func nearest_smaller_to_right(nums []int, n int) []int {
	st := stack.New[int]()
	res := make([]int, n)

	res[n-1] = -1
	st.Push(n - 1)
	for i := n - 2; i >= 0; i-- {
		for !st.Empty() && nums[st.Top()] >= nums[i] {
			st.Pop()
		}
		if st.Empty() {
			res[i] = -1
		} else {
			res[i] = st.Top()
		}
		st.Push(i)
	}

	return res
}

func computeWidth(index, n, r, l int) int {
	if r == -1 {
		// no smaller histogram found in right
		r = n
	}
	// This if-else is same as l = l + 1 => result is r - (l + 1)
	// if l == -1 {
	// 	// no smaller histogram found in left
	// 	l = 0
	// }else {
	// 	l = l + 1
	// }

	// Look at the image: https://media.geeksforgeeks.org/wp-content/cdn-uploads/histogram1.png
	// l = 1, r = 5, width = 5 - 1 - 1 = 3
	return r - l - 1
}

func histogram_max_area(nums []int, n int) int {
	leftRes := nearest_smaller_to_left(nums, n)
	rightRes := nearest_smaller_to_right(nums, n)

	fmt.Println(leftRes)
	fmt.Println(rightRes)

	maxArea := 0
	for i := 0; i < n; i++ {
		width := computeWidth(i, n, rightRes[i], leftRes[i])
		area := width * nums[i]
		maxArea = max(maxArea, area)
	}

	return maxArea
}

func main() {
	n := readInt()
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = readInt()
	}

	fmt.Println(histogram_max_area(nums, n))
}
