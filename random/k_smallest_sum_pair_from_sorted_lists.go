package main

import (
	"cmp"
	"fmt"
	"log"

	"github.com/liyue201/gostl/ds/priorityqueue"
)

type Val struct {
	i        int
	j        int
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

func k_smallest_sum_pair_from_sorted_list(nums1 []int, nums2 []int, k int) [][]int {
	res := [][]int{}
	minHeap := priorityqueue.New[Val](func(a, b Val) int {
		return cmp.Compare(a.priority, b.priority)
	})
	minHeap.Push(Val{
		0, 0, nums1[0] + nums2[0],
	})
	vis := map[[2]int]bool{}
	vis[[2]int{0, 0}] = true
	for k > 0 && !minHeap.Empty() {
		top := minHeap.Pop()
		res = append(res, []int{nums1[top.i], nums2[top.j]})
		k--

		if top.i+1 < len(nums1) {
			if _, isPresent := vis[[2]int{top.i + 1, top.j}]; !isPresent {
				minHeap.Push(Val{
					top.i + 1,
					top.j,
					nums1[top.i+1] + nums2[top.j],
				})
				vis[[2]int{top.i + 1, top.j}] = true
			}
		}
		if top.j+1 < len(nums2) {
			if _, isPresent := vis[[2]int{top.i, top.j + 1}]; !isPresent { // check for duplicate pair
				minHeap.Push(Val{
					i:        top.i,
					j:        top.j + 1,
					priority: nums1[top.i] + nums2[top.j+1],
				})
				vis[[2]int{top.i, top.j + 1}] = true
			}
		}
	}

	return res
}

func main() {
	n1, n2, k := readInt(), readInt(), readInt()
	nums1 := make([]int, n1)
	nums2 := make([]int, n2)
	for i := 0; i < n1; i++ {
		nums1[i] = readInt()
	}
	for i := 0; i < n2; i++ {
		nums2[i] = readInt()
	}

	fmt.Println(k_smallest_sum_pair_from_sorted_list(nums1, nums2, k))
}
