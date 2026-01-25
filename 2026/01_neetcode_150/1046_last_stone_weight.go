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

func last_stone_weight(arr []int, n int) int {
	maxHeap := priorityqueue.New[int](comparator.Reverse(comparator.IntComparator))
	for i := 0; i < n; i++ {
		maxHeap.Push(arr[i])
	}
	for !maxHeap.Empty() && maxHeap.Size() != 1 {
		a := maxHeap.Pop()
		b := maxHeap.Pop()
		if a == b {
			continue
		} else {
			maxHeap.Push(a - b) // a i greater than b as it was popped first
		}
	}
	if maxHeap.Empty() {
		return 0
	}
	return maxHeap.Pop()
}

func main() {
	n := readInt()
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = readInt()
	}

	fmt.Printf("last_stone_weight(arr, n): %v\n", last_stone_weight(arr, n))
}
