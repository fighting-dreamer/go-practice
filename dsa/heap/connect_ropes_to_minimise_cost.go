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

func connect_ropes_to_minimise_cost(nums []int, n int) int {
	minHeap := priorityqueue.New[int](comparator.IntComparator)
	for i := 0; i < n; i++ {
		minHeap.Push(nums[i])
	}
	res := 0
	var first, second int
	for minHeap.Size() != 1 {
		first, second = minHeap.Pop(), minHeap.Pop()
		res = res + first + second
		minHeap.Push(first + second)
	}

	return res
}

func main() {
	n := readInt()
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = readInt()
	}
	fmt.Println(connect_ropes_to_minimise_cost(nums, n))

	return
}
