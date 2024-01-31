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

func stock_span(nums []int, n int) []int {
	res := make([]int, n)
	res[0] = 1

	st := stack.New[int]()
	st.Push(0)
	for i := 1; i < n; i++ {
		for !st.Empty() && nums[st.Top()] <= nums[i] {
			st.Pop()
		}

		if st.Empty() {
			res[i] = 1
		} else {
			res[i] = i - st.Top()
		}
		st.Push(i)
	}

	return res
}

func main() {
	n := readInt()

	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = readInt()
	}

	fmt.Println(stock_span(nums, n))
}
