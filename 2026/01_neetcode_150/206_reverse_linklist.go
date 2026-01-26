package main

import (
	"fmt"
	"math/rand"
)

type Node struct {
	val  int
	next *Node
}

func NewNode(val int) *Node {
	return &Node{
		val:  val,
		next: nil,
	}
}

func AddNode(head *Node, val int) *Node {
	if head == nil {
		head = NewNode(val)
	} else {
		curr := head
		for curr.next != nil {
			curr = curr.next
		}
		curr.next = NewNode(val)
	}

	return head
}

func printList(head *Node) {
	for head != nil {
		fmt.Printf("%d,", head.val)
		head = head.next
	}
	fmt.Println()
}

func reverseList(head *Node) *Node {
	curr := head
	var prev *Node
	var next *Node

	for curr != nil {
		next = curr.next
		curr.next = prev

		prev = curr
		curr = next
	}

	return prev
}

func main() {
	for i := 0; i < 10; i++ {
		n := rand.Intn(10)
		var head *Node
		for i := 0; i < n; i++ {
			head = AddNode(head, rand.Intn(1000))
		}
		fmt.Println("Before : ")
		printList(head)
		fmt.Println("After : ")
		head = reverseList(head)
		printList(head)
	}
}
