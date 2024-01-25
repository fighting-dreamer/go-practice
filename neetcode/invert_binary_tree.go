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
		return &Node[T]{
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

func PrintInorder[T any](head *Node[T]) {
	if head == nil {
		return
	}

	PrintInorder(head.Left)
	fmt.Println(head.Val)
	PrintInorder(head.Right)
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

func InvertBinaryTree[T any](head *Node[T]) {
	if head == nil {
		return
	}
	head.Left, head.Right = head.Right, head.Left

	InvertBinaryTree(head.Left)
	InvertBinaryTree(head.Right)
}

func main() {
	var head *Node[int]
	n := readInt()

	for i := 0; i < n; i++ {
		head = Add(head, readInt())
	}
	PrintInorder(head)
	fmt.Println("---------------------")
	InvertBinaryTree(head)
	fmt.Println("---------------------")
	PrintInorder(head)
	fmt.Println("---------------------")
	InvertBinaryTree(head)
	fmt.Println("---------------------")
	PrintInorder(head) // this is suppose to be same as first inorder print value.

}
