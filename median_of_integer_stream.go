package main

import (
	"fmt"
	"log"

	"github.com/liyue201/gostl/ds/priorityqueue"
	"github.com/liyue201/gostl/utils/comparator"
)

func getMedian(minHeap, maxHeap *priorityqueue.PriorityQueue[int]) float64 {
	l := minHeap.Size() + maxHeap.Size()
	if l%2 == 0 {
		return (float64(minHeap.Top()) + float64(maxHeap.Top())) / 2
	} else {
		if minHeap.Size() > maxHeap.Size() {
			return float64(minHeap.Top())
		} else {
			return float64(maxHeap.Top())
		}
	}
}

func process(minHeap, maxHeap *priorityqueue.PriorityQueue[int], val int) {
	if maxHeap.Size() > 0 && maxHeap.Top() < val {
		maxHeap.Push(val)
	} else {
		minHeap.Push(val)
	}

	if maxHeap.Size() > minHeap.Size()+1 {
		minHeap.Push(maxHeap.Pop())
	}

	if minHeap.Size() > maxHeap.Size()+1 {
		maxHeap.Push(minHeap.Pop())
	}
}

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

func main() {
	n := readInt()

	maxHeap := priorityqueue.New[int](comparator.IntComparator) // maxheap
	minHeap := priorityqueue.New[int](comparator.Reverse[int](comparator.IntComparator))

	for i := 0; i < n; i++ {
		t := readInt()
		if t == -1 {
			fmt.Println(getMedian(minHeap, maxHeap))
		} else {
			process(maxHeap, minHeap, t)
		}
	}
}
