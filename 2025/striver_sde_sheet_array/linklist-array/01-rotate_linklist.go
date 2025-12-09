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

func rotateLinklist(head *Node, k int) *Node {
	start := head
	var old *Node = nil
	for i := 0; i < k; i++ {
		old = head
		head = head.next
	}
	old.next = nil

	//---------//

	curr := head
	for curr != nil {
		old = curr
		curr = curr.next
	}
	old.next = start

	return head
}

func main() {
	s := readString()
	k := readInt()
	var head *Node = nil
	for i := 0; i < len(s); i++ {
		head = addLinklist(head, int(s[i])-int('0'))
	}

	printLinklist(head)
	head = rotateLinklist(head, k)
	printLinklist(head)
}
