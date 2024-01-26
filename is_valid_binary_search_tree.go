package main

// STAR : https://go.dev/play/p/Z9jtjo6Rdbq

import (
	"cmp"
	"fmt"
	"log"
	"math"

	"golang.org/x/exp/constraints"
)

var (
	negInf = math.Inf(-1)
	posInf = math.Inf(1)
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

func PrintInorder[T any](head *Node[T]) {
	if head == nil {
		return
	}
	PrintInorder(head.Left)
	fmt.Println(head.Val)
	PrintInorder(head.Right)

	return
}

// the leftLimit and rightLimit are of type float, but need to be convertable to type T, for us it could be int.
func isValidBinarySearchTree[T constraints.Integer | constraints.Float](head *Node[T], leftLimit, rightLimit float64) bool {
	if head == nil {
		return true
	}
	if (leftLimit == negInf) || (cmp.Compare(float64(head.Val), leftLimit) >= 0) &&
		rightLimit == posInf || cmp.Compare(float64(head.Val), rightLimit) <= 0 {
		return isValidBinarySearchTree(head.Left, leftLimit, head.Val) && isValidBinarySearchTree(head.Right, head.Val, rightLimit)
	}
	return false
}

func isValidBST[T cmp.Ordered](head *Node[T]) bool {
	return isValidBinarySearchTree(head, negInf, posInf)
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
	var head *Node[int]
	for i := 0; i < n; i++ {
		head = Add(head, readInt())
	}
	isValidBST(head)
}
