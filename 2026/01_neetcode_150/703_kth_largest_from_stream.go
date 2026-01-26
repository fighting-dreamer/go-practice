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

func kth_Largest(arr []int, n int, k int) {
	// after initial setup, for each new number, give he kth largest
	minheap := priorityqueue.New[int](comparator.IntComparator)
	for i := 0; i < k; i++ {
		minheap.Push(arr[i])
	}
	for i := k; i < n; i++ {
		if minheap.Top() < arr[i] {
			minheap.Pop()
			minheap.Push(arr[i])
		}
	}
	newAdditonCount := readInt()
	for i := 0; i < newAdditonCount; i++ {
		t := readInt()
		if minheap.Top() < t {
			minheap.Pop()
			minheap.Push(t)
		}
		fmt.Println(t, " => ", minheap.Top())
	}
}

func main() {
	k := readInt()
	n := readInt()
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = readInt()
	}

	kth_Largest(arr, n, k)
}
