package main

// STAR : https://go.dev/play/p/Z9jtjo6Rdbq

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

func PrintInorder[T any](head *Node[T]) {
	if head == nil {
		return
	}
	PrintInorder(head.Left)
	fmt.Println(head.Val)
	PrintInorder(head.Right)

	return
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

func ModifiedInorder[T any](head *Node[T], counter *int) *Node[T] {
	if head == nil {
		return nil
	}
	res := ModifiedInorder(head.Left, counter)
	if res != nil {
		return res
	}
	(*counter)--
	if *counter == 0 {
		return head
	}

	return ModifiedInorder(head.Right, counter)
}

func kthSmallestElementInBST[T any](head *Node[T], k int) T {
	counter := &k
	res := ModifiedInorder(head, counter)
	fmt.Println("check : ", *counter, res)

	return res.Val
}

func main() {
	n, k := readInt(), readInt()
	var head *Node[int]
	for i := 0; i < n; i++ {
		head = Add(head, readInt())
	}
	PrintInorder(head)
	fmt.Println()
	fmt.Println(kthSmallestElementInBST(head, k))
}
