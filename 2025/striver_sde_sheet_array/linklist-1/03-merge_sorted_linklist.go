package main

import (
	"fmt"
	"log"
)

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

func addLinklist(head *Node, val int) *Node {
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

func printLinklist(head *Node) {
	for head != nil {
		fmt.Printf("%d ", head.val)
		head = head.next
	}
	fmt.Println()
}

func merge_sorted_linklists(head1 *Node, head2 *Node) *Node {
	head := NewNode(0)
	headCopy := head
	for head1 != nil && head2 != nil {
		if head1.val < head2.val {
			head.next = head1
			head1 = head1.next
		} else {
			head.next = head2
			head2 = head2.next
		}
		head = head.next
	}

	for head1 != nil {
		head.next = head1
		head1 = head1.next
		head = head.next

	}

	for head2 != nil {
		head.next = head2
		head2 = head2.next
		head = head.next

	}

	return headCopy.next // removing the dummy first head node.
}

func main() {
	n := readInt()
	m := readInt()
	var head1 *Node = nil
	var head2 *Node = nil

	for i := 0; i < n; i++ {
		head1 = addLinklist(head1, readInt())
	}
	for i := 0; i < m; i++ {
		head2 = addLinklist(head2, readInt())
	}

	printLinklist(head1)
	printLinklist(head2)
	head := merge_sorted_linklists(head1, head2)
	printLinklist(head)
}
