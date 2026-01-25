package main

import (
	"fmt"
	"log"
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
		fmt.Printf("%2d", head.val)
		head = head.next
	}
	fmt.Println()
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

func merge_sorted_lists(listA *Node, listB *Node) *Node {
	res := NewNode(-1)
	resTail := res
	var temp *Node
	for listA != nil && listB != nil {
		if listA.val < listB.val {
			temp = listA
			listA = listA.next
		} else {
			temp = listB
			listB = listB.next
		}
		resTail.next = temp
		resTail = temp
	}
	for listA != nil {
		resTail.next = listA
		resTail = listA

		listA = listA.next
	}

	for listB != nil {
		resTail.next = listB
		resTail = listB

		listB = listB.next
	}

	return res.next // first node was just a placeholder.
}

func main() {
	var headA *Node
	var headB *Node
	n := readInt()
	for i := 0; i < n; i++ {
		headA = AddNode(headA, readInt())
	}
	printList(headA)

	m := readInt()
	for i := 0; i < m; i++ {
		headB = AddNode(headB, readInt())
	}
	printList(headB)

	res := merge_sorted_lists(headA, headB)
	printList(res)
}
