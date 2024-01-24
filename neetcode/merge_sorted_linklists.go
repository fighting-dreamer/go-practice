package main

import (
	"cmp"
	"fmt"
	"log"

	"nipun.io/coding/dsa"
)

// STAR : cmp.Ordered to ensure comparable data of Val

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

func merge_sorted_linklists[T cmp.Ordered](l1 *dsa.List[T], l2 *dsa.List[T]) *dsa.List[T] {
	lm := dsa.NewLinkList[T]()
	lm.Size = l1.Size + l2.Size

	head1 := l1.Head
	head2 := l2.Head

	if head1.Val <= head2.Val {
		lm.Head = head1
		head1 = head1.Next
	} else {
		lm.Head = head2
		head2 = head2.Next
	}
	head := lm.Head
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

	for head1 != nil {
		head.Next = head1
		head1 = head1.Next
		head = head.Next
	}

	for head2 != nil {
		head.Next = head2
		head2 = head2.Next
		head = head.Next
	}

	return lm
}

func main() {
	n, m := readInt(), readInt()

	l1 := dsa.NewLinkList[int]()
	l2 := dsa.NewLinkList[int]()

	for i := 0; i < n; i++ {
		l1.Add(readInt())
	}
	for i := 0; i < m; i++ {
		l2.Add(readInt())
	}

	lm := merge_sorted_linklists[int](l1, l2)

	lm.Print()
}
