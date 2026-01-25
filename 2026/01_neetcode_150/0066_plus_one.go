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

func printList(head *Node) {
	for head != nil {
		fmt.Printf("%2d,", head.val)
		head = head.next
	}
	fmt.Println()
	return
}

func buildList(n int) *Node {
	var head *Node
	for n/10 != 0 {
		head = AddNode(head, n%10)
		n = n / 10
	}
	head = AddNode(head, n%10)
	return head
}

func plus_one(head *Node) *Node {
	carry := 1
	curr := head
	var old *Node
	for curr != nil {
		old = curr

		temp := curr.val + carry
		curr.val = temp % 10
		carry = temp / 10

		curr = curr.next
	}

	if carry > 0 {
		old.next = NewNode(carry)
	}

	return head
}

func main() {
	n := readInt()
	head := buildList(n)
	printList(head)

	head = plus_one(head)
	printList(head)
}
