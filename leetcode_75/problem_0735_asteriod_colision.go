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

func asteriod_collision(nums []int, n int) []int {
	st := stack.New[int]()
	st.Push(nums[0])
	for i := 1; i < n; i++ {
		if st.Top() > 0 && nums[i] < 0 {
			wt := -nums[i]
			for !st.Empty() && st.Top() < wt {
				st.Pop()
			}
			if st.Top() == wt {
				st.Pop()
				continue
			}
			if st.Empty() {
				st.Push(nums[i])
			}
		} else {
			st.Push(nums[i])
		}
	}
	res := make([]int, st.Size())
	i := st.Size() - 1
	for i >= 0 {
		res[i] = st.Pop()
		i--
	}

	return res
}

func main() {
	n := readInt()
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = readInt()
	}

	fmt.Println(asteriod_collision(nums, n))
}
