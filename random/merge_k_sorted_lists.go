package main

import (
	"cmp"
	"fmt"
	"log"

	"github.com/liyue201/gostl/ds/priorityqueue"
)

type Val struct {
	row      int
	col      int
	priority int
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

func merge_k_sorted_lists(lists [][]int, k int) []int {
	l := 0
	for i := 0; i < k; i++ {
		l += len(lists[i])
	}
	result := make([]int, l)
	minHeap := priorityqueue.New[Val](func(a, b Val) int {
		return cmp.Compare(a.priority, b.priority)
	})
	i := 0
	for j := 0; j < k; j++ {
		minHeap.Push(Val{
			row:      j,
			col:      0,
			priority: lists[j][0],
		})
	}

	for i < l {
		top := minHeap.Pop()
		result[i] = top.priority
		i++
		if top.col+1 < len(lists[top.row]) {
			minHeap.Push(Val{
				row:      top.row,
				col:      top.col + 1,
				priority: lists[top.row][top.col+1],
			})
		}
	}

	return result
}

func main() {
	k := readInt()
	lists := make([][]int, k)
	for i := 0; i < k; i++ {
		n := readInt()
		lists[i] = make([]int, n)
		for j := 0; j < n; j++ {
			lists[i][j] = readInt()
		}
	}

	fmt.Println(merge_k_sorted_lists(lists, k))
}
