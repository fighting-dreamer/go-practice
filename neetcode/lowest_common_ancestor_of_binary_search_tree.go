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

func lowest_common_ancestor[T cmp.Ordered](head *Node[T], val1, val2 T) *Node[T] {
	if head == nil {
		return nil
	}

	if head.Val == val1 || head.Val == val2 {
		return head
	}

	left := lowest_common_ancestor(head.Left, val1, val2)
	right := lowest_common_ancestor(head.Right, val1, val2)

	if left != nil && right != nil {
		return head
	}

	if left != nil {
		return left
	}
	return right
}

// input : 9 6 2 3 0 4 7 9 3 5 0 3
// output : 2
func main() {
	n := readInt()

	var head *Node[int]
	for i := 0; i < n; i++ {
		head = Add(head, readInt())
	}

	fmt.Println(func() int {
		n := lowest_common_ancestor(head, readInt(), readInt())
		return n.Val
	}())
}
