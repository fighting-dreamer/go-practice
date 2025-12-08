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

func getNthNode(head *Node, n int) *Node {
	for i := 0; i < n-1 && head != nil; i++ {
		head = head.next
	}
	return head
}

func delete_Nth_node_using_nth_node(nthNode *Node) {
	var prev *Node = nil
	curr := nthNode
	// moving/copying values from next node to curr node till we reach last node.
	for curr.next != nil {
		curr.val = curr.next.val

		prev = curr
		curr = curr.next
	}
	// curr next is nil => curr is last node
	// we delete it
	prev.next = nil
}

func main() {
	s := readString()
	k := readInt()
	var head *Node = nil

	for i := len(s) - 1; i >= 0; i-- {
		head = addLinklist(head, int(s[i])-int('0'))
	}
	kthNode := getNthNode(head, k)
	printLinklist(head)
	delete_Nth_node_using_nth_node(kthNode)
	printLinklist(head)
}
