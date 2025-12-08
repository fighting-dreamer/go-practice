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

func reverse_k_nodes_in_linklist(head *Node, k int) (*Node, *Node) {
	var prev *Node = nil
	var curr *Node = head
	var next *Node = nil
	for k > 0 && curr != nil {
		next = curr.next
		curr.next = prev

		prev = curr
		curr = next

		k--
	}

	return prev, head
}

func reverse_in_group_of_k(head *Node, k int) *Node {
	head, last := reverse_k_nodes_in_linklist(head, k)

	curr := head
	lastt := last
	fmt.Println(head.val, last.val)
	printLinklist(head)
	fmt.Println()
	for lastt != nil && lastt.next != nil {
		curr, lastt = reverse_k_nodes_in_linklist(lastt.next, k)
		fmt.Println(curr.val, lastt.val)
		printLinklist(lastt)

		last.next = curr
		last = lastt
	}

	return head
}

func main() {
	s := readString()
	k := readInt()
	var head *Node = nil
	for i := len(s) - 1; i >= 0; i-- {
		head = addLinklist(head, int(s[i])-int('0'))
	}
	printLinklist(head)
	head = reverse_in_group_of_k(head, k)
	printLinklist(head)
}
