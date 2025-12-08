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

func remove_kth_node(head *Node, k int) *Node {
	if k == 1 {
		head = head.next
	} else {
		var prevKthNode *Node = nil
		kthNode := head
		for i := 0; i < k-1 && kthNode != nil; i++ {
			prevKthNode = kthNode
			kthNode = kthNode.next
		}
		// kthNode points to Kth node
		if kthNode != nil {
			prevKthNode.next = kthNode.next
		}
	}
	return head
}

func main() {
	n := readInt()

	var head *Node = nil
	for i := 0; i < n; i++ {
		head = addLinklist(head, readInt())
	}
	printLinklist(head)

	fmt.Println("\nnode to be removed : ")
	k := readInt()
	head = remove_kth_node(head, k)
	printLinklist(head)
}
