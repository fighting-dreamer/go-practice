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

func sort_k_sorted(nums []int, n, k int) []int {
	minHeap := priorityqueue.New[int](comparator.IntComparator)
	for i := 0; i < k; i++ {
		minHeap.Push(nums[i])
	}
	for i := k; i < n; i++ {
		minHeap.Push(nums[i])
		if minHeap.Size() > k {
			nums[i-k] = minHeap.Top()
			minHeap.Pop()
		}
	}
	i := n - k
	for !minHeap.Empty() {
		nums[i] = minHeap.Pop()
		i++
	}

	return nums
}

func main() {
	n, k := readInt(), readInt()

	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = readInt()
	}

	nums = sort_k_sorted(nums, n, k)
	fmt.Println(nums)
}
