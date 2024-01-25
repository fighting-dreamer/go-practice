package main

import (
	"cmp"
	"fmt"
	"log"
)

type Node[T any] struct {
	Val   T
	Left  *Node[T]
	Right *Node[T]
}

func Add[T cmp.Ordered](head *Node[T], val T) *Node[T] {
	if head == nil {
		head = &Node[T]{
			Val: val,
		}
	} else {
		curr := head
		var old *Node[T]
		for curr != nil {
			old = curr
			if curr.Val >= val {
				curr = curr.Left
			} else {
				curr = curr.Right
			}
		}

		if old.Val >= val {
			old.Left = &Node[T]{
				Val: val,
			}
		} else {
			old.Right = &Node[T]{
				Val: val,
			}
		}
	}
	return head
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

func Inorder[T any](head *Node[T]) []T {
	if head == nil {
		return []T{}
	}

	vals := Inorder(head.Left)
	vals = append(vals, head.Val)
	vals = append(vals, Inorder(head.Right)...)

	return vals
}

func isSameTree[T cmp.Ordered](head1, head2 *Node[T]) bool {
	vals1 := Inorder(head1)
	vals2 := Inorder(head2)

	if len(vals1) != len(vals2) {
		return false
	}
	for i, _ := range vals1 {
		if vals1[i] != vals2[i] {
			return false
		}
	}

	return true
}

func main() {
	n, m := readInt(), readInt()

	var head1 *Node[int]
	for i := 0; i < n; i++ {
		head1 = Add(head1, readInt())
	}

	var head2 *Node[int]
	for i := 0; i < m; i++ {
		head2 = Add(head2, readInt())
	}

	fmt.Println(isSameTree(head1, head2))
}
