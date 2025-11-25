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

func addNode(head *Node, val int) *Node {
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
		fmt.Printf("%d ", head.val)
		head = head.next
	}
	fmt.Println()
}

func getListValueAtIndex(head *Node, index int) int {
	if index == 0 {
		return head.val
	}
	for index != 0 && head.next != nil {
		head = head.next
		index--
	}
	if index != 0 {
		return -1
	}
	return head.val
}

func deleteList(head *Node, val int) *Node {
	if head == nil {
		return nil
	}

	if head.val == val {

		return head.next

	}
	curr := head
	for curr.next != nil && curr.next.val != val {
		curr = curr.next
	}

	if curr.next != nil {
		curr.next = curr.next.next
	}
	return head
}

func reverseLinkList(head *Node) *Node {
	var curr, prev, nnext *Node = head, nil, nil
	for curr != nil {
		nnext = curr.next
		curr.next = prev
		prev = curr
		curr = nnext
	}

	return prev
}

func main() {
	var head *Node = nil
	n := rand.Intn(20)
	for i := 0; i < n; i++ {
		head = addNode(head, rand.Intn(100))
	}
	printList(head)

	toDeleteVal := getListValueAtIndex(head, rand.Intn(20))
	if toDeleteVal == -1 {
		toDeleteVal = getListValueAtIndex(head, 0)
	}
	fmt.Println("to delete val : ", toDeleteVal)
	head = deleteList(head, toDeleteVal)
	printList(head)
	head.next.next = nil
	head = reverseLinkList(head)
	printList(head)
}
