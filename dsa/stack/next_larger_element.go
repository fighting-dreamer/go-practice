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

func next_larger_element_to_right_brute_force(nums []int, n int) []int {
	res := make([]int, n)
	for i := 0; i < n; i++ {
		res[i] = -1
		for j := i + 1; j < n; j++ {
			if nums[j] > nums[i] {
				res[i] = j
				break
			}
		}
	}

	return res
}

func next_larger_element_to_right_with_stack(nums []int, n int) []int {
	st := stack.New[int]()
	st.Push(nums[n-1])
	res := make([]int, n)
	res[n-1] = -1
	for i := n - 2; i >= 0; i-- {
		for !st.Empty() && st.Top() <= nums[i] {
			st.Pop()
		}
		if st.Empty() {
			res[i] = -1
		} else {
			res[i] = st.Top()
		}
		st.Push(nums[i])
	}

	return res
}

func next_larger_element_to_right(nums []int, n int) []int {
	resBruteForce := next_larger_element_to_right_brute_force(nums, n)
	resWithStack := next_larger_element_to_right_with_stack(nums, n)
	fmt.Println(resBruteForce)
	fmt.Println(resWithStack)
	return resWithStack
}

func next_larger_element_to_left_with_brute_force(nums []int, n int) []int {
	res := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		res[i] = -1
		for j := i - 1; j >= 0; j-- {
			if nums[j] > nums[i] {
				res[i] = nums[j]
				break
			}
		}
	}

	return res
}

func next_larger_element_to_left_with_stack(nums []int, n int) []int {
	st := stack.New[int]()
	res := make([]int, n)

	st.Push(nums[0])
	res[0] = -1
	for i := 1; i < n; i++ {
		for !st.Empty() && st.Top() <= nums[i] {
			st.Pop()
		}
		if st.Empty() {
			res[i] = -1
		} else {
			res[i] = st.Top()
		}
		st.Push(nums[i])
	}

	return res
}

func next_larger_element_to_left(nums []int, n int) []int {
	resWithBruteForce := next_larger_element_to_left_with_brute_force(nums, n)
	resWithStack := next_larger_element_to_left_with_stack(nums, n)
	fmt.Println(resWithBruteForce)
	fmt.Println(resWithStack)
	return resWithStack
}

func main() {
	n := readInt()
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = readInt()
	}

	fmt.Println(next_larger_element_to_right(nums, n))

	fmt.Println("------------------")
	fmt.Println(next_larger_element_to_left(nums, n))
}
