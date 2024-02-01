package main

import (
	"fmt"
	"log"

	"github.com/liyue201/gostl/ds/deque"
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

// input : 1 3 2 1 5 -1 -3 4 3 6 7, k = 3
// dq : 1 -> 3 -> 3 2 -> 3 2 1 -> 5 -> 5 -1 -> 5 -1 -3 -> 4 -> 4 3 -> 6 -> 7
// output : 3, 3, 5, 5, 5, 4, 4, 6, 7

func max_elemet_in_all_subarray(nums []int, n, k int) []int {
	dq := deque.New[int]()
	res := []int{}
	for i, j := 0, 0; i < n && j < n; {
		fmt.Println(i, j, dq)
		if dq.Empty() || nums[j] > dq.Front() {
			for !dq.Empty() && dq.Front() < nums[j] {
				dq.PopFront()
			}
			dq.PushFront(nums[j])
		} else {
			dq.PushBack(nums[j])
		}

		if j-i+1 < k {
			j++
		} else if j-i+1 == k {
			// fmt.Println("window : ", i, j, dq, nums[i])
			res = append(res, dq.Front())

			if nums[i] == dq.Front() {
				dq.PopFront()
			}

			i++
			j++
		}
	}

	return res
}

// input : 9 3 1 2 3 1 4 5 2 3 6
// output : 3 3 4 5 5 6

// input : 11 3 1 3 2 1 5 -1 -3 4 3 6 7
// output : 3 3 5 5 5 4 4 6 7

// input : 10 4 8 5 10 7 9 4 15 12 90 13
// output : 10 10 10 15 15 90 90
func main() {
	n, k := readInt(), readInt()
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = readInt()
	}

	fmt.Println(max_elemet_in_all_subarray(nums, n, k))
}
