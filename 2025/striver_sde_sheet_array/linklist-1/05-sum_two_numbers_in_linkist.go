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

func addNumbersUsingLinkList(head1 *Node, head2 *Node) *Node {
	if head1 == nil {
		return head2
	}
	if head2 == nil {
		return head1
	}

	var res *Node = NewNode(-1) // dummy node
	curr := res
	carry := 0
	for head1 != nil && head2 != nil {
		s := head1.val + head2.val + carry
		if s > 10 {
			carry = s / 10
			s = s % 10
		}
		curr.next = NewNode(s)
		curr = curr.next
		head1 = head1.next
		head2 = head2.next
	}
	for head1 != nil {
		s := head1.val + carry
		if s > 10 {
			carry = s / 10
			s = s % 10
		}
		curr.next = NewNode(s)
		curr = curr.next

		head1 = head1.next

	}
	for head2 != nil {
		s := head2.val + carry
		if s > 10 {
			carry = s / 10
			s = s % 10
		}
		curr.next = NewNode(s)
		curr = curr.next

		head2 = head2.next
	}

	return res.next
}

func main() {
	s1, s2 := readString(), readString()
	var head1 *Node = nil
	var head2 *Node = nil
	for i := len(s1) - 1; i >= 0; i-- {
		head1 = addLinklist(head1, int(s1[i])-int('0'))
	}
	for i := len(s2) - 1; i >= 0; i-- {
		head2 = addLinklist(head2, int(s2[i])-int('0'))
	}
	printLinklist(head1)
	printLinklist(head2)

	resHead := addNumbersUsingLinkList(head1, head2)
	printLinklist(resHead)
}
