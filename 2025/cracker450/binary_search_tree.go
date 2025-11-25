package main

import (
	"fmt"

	"golang.org/x/exp/rand"
)

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func NewNode(val int) *Node {
	return &Node{
		Val:   val,
		Left:  nil,
		Right: nil,
	}
}

func addBSTNode(head *Node, val int) *Node {
	if head == nil {
		head = NewNode(val)
	} else {
		var curr *Node = head
		var old *Node = nil
		for curr != nil {
			old = curr
			if curr.Val <= val {
				curr = curr.Right
			} else {
				curr = curr.Left
			}
		}

		if old.Val <= val {
			old.Right = NewNode(val)
		} else {
			old.Left = NewNode(val)
		}
	}
	return head
}

func printInorder(head *Node) {
	if head == nil {
		return
	}
	printInorder(head.Left)
	fmt.Printf("%d ", head.Val)
	printInorder(head.Right)
}

func printPreOrder(head *Node) {
	if head == nil {
		fmt.Printf("nil ")
		return
	}
	fmt.Printf("%d ", head.Val)
	fmt.Printf("Left : ")
	printPreOrder(head.Left)
	fmt.Printf("Right : ")
	printPreOrder(head.Right)
}

func main() {
	var head *Node = nil
	n := rand.Intn(20)
	for i := 0; i < n; i++ {
		head = addBSTNode(head, rand.Intn(500))
	}
	printPreOrder(head)
	fmt.Println()
	printInorder(head)
}
