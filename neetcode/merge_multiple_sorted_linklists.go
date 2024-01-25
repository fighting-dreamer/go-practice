package main

import (
	"cmp"
	"fmt"
	"log"

	"nipun.io/coding/dsa"
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

func mergeLinkList[T cmp.Ordered](list1, list2 *dsa.List[T]) *dsa.List[T] {
	if list1.Size == 0 && list2.Size == 0 {
		return &dsa.List[T]{}
	}
	if list1.Size == 0 {
		return list2
	}
	if list2.Size == 0 {
		return list1
	}

	var head *dsa.Node[T]
	lm := &dsa.List[T]{}
	lm.Size = list1.Size + list2.Size

	head1 := list1.Head
	head2 := list2.Head
	fmt.Println("-------------------------")
	list1.Print()
	list2.Print()
	fmt.Println("-------------------------")

	if head1.Val <= head2.Val {
		head = head1
		head1 = head1.Next
	} else {
		head = head2
		head2 = head2.Next
	}
	lm.Head = head
	for head1 != nil && head2 != nil {
		if head1.Val <= head2.Val {
			head.Next = head1
			head1 = head1.Next
		} else {
			head.Next = head2
			head2 = head2.Next
		}
		head = head.Next
	}

	if head1 != nil {
		head.Next = head1
	} else {
		head.Next = head2
	}

	return lm
}

func mergeMultipleSortedLinklists[T cmp.Ordered](lists []dsa.List[T], lengths []int, start int, end int) *dsa.List[T] {
	if start >= end {
		return &lists[end]
	}
	fmt.Println(start, end)

	mid := (start + end) / 2
	list1 := mergeMultipleSortedLinklists[T](lists, lengths, start, mid)
	list2 := mergeMultipleSortedLinklists[T](lists, lengths, mid+1, end)
	return mergeLinkList[T](list1, list2)
}

func merge_multiple_sorted_linklists[T cmp.Ordered](lists []dsa.List[T], lengths []int, k int) *dsa.List[T] {
	return mergeMultipleSortedLinklists[T](lists, lengths, 0, k-1)
}

func print_lists[T any](lists []dsa.List[T]) {
	for i := 0; i < len(lists); i++ {
		lists[i].Print()
	}
	fmt.Println()
}

func main() {
	k := readInt()
	lists := make([]dsa.List[int], k)
	lengths := make([]int, k)
	for i := 0; i < k; i++ {
		lengths[i] = readInt()
	}
	for i := 0; i < k; i++ {
		for j := 0; j < lengths[i]; j++ {
			lists[i].Add(readInt())
		}
	}
	// print_lists[int](lists)
	lm := merge_multiple_sorted_linklists(lists, lengths, k)
	// print_lists[int](lists)
	fmt.Println()
	fmt.Println()
	lm.Print()
}
