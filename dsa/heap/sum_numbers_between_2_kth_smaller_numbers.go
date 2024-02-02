package main

import (
	"fmt"
	"log"

	"github.com/liyue201/gostl/ds/priorityqueue"
	"github.com/liyue201/gostl/utils/comparator"
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

func findKth_smaller(nums []int, n, k int) int {
	maxHeap := priorityqueue.New[int](comparator.Reverse[int](comparator.IntComparator))
	for i := 0; i < n; i++ {
		maxHeap.Push(nums[i])
		if maxHeap.Size() > k {
			maxHeap.Pop()
		}
	}

	return maxHeap.Pop()
}

func find_sum_between_k1_k2th_smaller_numbers(nums []int, n, k1, k2 int) int {
	k1th := findKth_smaller(nums, n, k1)
	k2th := findKth_smaller(nums, n, k2)
	if k2 < k1 {
		// swap
		k2th, k1th = k1th, k2th
	}
	fmt.Println(k1th, k2th)
	res := 0
	for i := 0; i < n; i++ {
		if nums[i] > k1th && nums[i] < k2th {
			res += nums[i]
			fmt.Println(nums[i], " ")
		}
	}

	return res
}

func main() {
	n, k1, k2 := readInt(), readInt(), readInt()

	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = readInt()
	}

	fmt.Println(find_sum_between_k1_k2th_smaller_numbers(nums, n, k1, k2))
}
