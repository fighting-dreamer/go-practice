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

func kth_smallest_element(nums []int, n, k int) int {
	maxHeap := priorityqueue.New[int](comparator.Reverse[int](comparator.IntComparator))
	for i := 0; i < k; i++ {
		maxHeap.Push(nums[i])
	}
	for i := k; i < n; i++ {
		if nums[i] < maxHeap.Top() {
			maxHeap.Pop()
			maxHeap.Push(nums[i])
		}
	}

	return maxHeap.Top()
}

func main() {
	n, k := readInt(), readInt()

	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = readInt()
	}
	fmt.Println(kth_smallest_element(nums, n, k))

}
