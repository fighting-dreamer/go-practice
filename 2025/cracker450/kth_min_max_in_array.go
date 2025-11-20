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
		log.Fatal("couldn;t read number")
	}
	return n
}

func print_array(arr []int, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", arr[i])
	}
	fmt.Println()
}

func minKth(arr []int, n int, k int) int {
	maxHeap := priorityqueue.New[int](comparator.Reverse[int](comparator.IntComparator))
	for i := 0; i < k; i++ {
		maxHeap.Push(arr[i])
	}
	// by end of this operation, all elements in heap are smaller k elements
	// top most element is the topKth smaller => minKth element
	for i := k; i < n; i++ {
		// top most ele is max elem, if arr[i] i smaller than that => keep it in heap
		if arr[i] < maxHeap.Top() {
			maxHeap.Pop()
			maxHeap.Push(arr[i])
		}
	}
	return maxHeap.Top()
}

func maxKth(arr []int, n int, k int) int {
	minHeap := priorityqueue.New[int](comparator.IntComparator)
	for i := 0; i < k; i++ {
		minHeap.Push(arr[i])
	}
	// by end of the operation, heap will hold top k elements
	// its a minHeap, the element will be kth max element
	for i := k; i < n; i++ {
		// if the arr[i] is larger => keep it in heap
		if arr[i] > minHeap.Top() {
			minHeap.Pop()
			minHeap.Push(arr[i])
		}
	}

	return minHeap.Top()
}

func find_kth_min_max_in_array(arr []int, n, k int) (int, int) {
	return minKth(arr, n, k), maxKth(arr, n, k)
}

func main() {
	n := readInt()
	k := readInt()
	fmt.Println("n : ", n, "k : ", k)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = readInt()
	}

	print_array(arr, n)
	minKthVal, maxKthVal := find_kth_min_max_in_array(arr, n, k)
	fmt.Printf("minKth : %d, maxKth : %d", minKthVal, maxKthVal)
}
