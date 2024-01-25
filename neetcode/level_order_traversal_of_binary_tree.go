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

func PrintInorder[T any](head *Node[T]) {
	if head == nil {
		return
	}

	PrintInorder(head.Left)
	fmt.Println(head.Val)
	PrintInorder(head.Right)
}

func LevelOrderTraversal[T any](head *Node[T]) {
	if head == nil {
		return
	}

	q := []*Node[T]{}
	q = append(q, head)
	q = append(q, nil)

	for len(q) > 1 {
		first := q[0]
		q = q[1:] // STAR assignment(q = q[1:]), not initliazation( q := q[1:]) : the use of simple '=' and not using ':='.
		if first == nil {
			q = append(q, nil)

			continue
		}

		fmt.Println(first.Val)
		if first.Left != nil {
			q = append(q, first.Left)
		}
		if first.Right != nil {
			q = append(q, first.Right)
		}
	}

	return
}

func main() {
	var head *Node[int]
	n := readInt()

	for i := 0; i < n; i++ {
		head = Add(head, readInt())
	}
	// PrintInorder(head)
	LevelOrderTraversal(head)
}
