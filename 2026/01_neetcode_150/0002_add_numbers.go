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

func buildList(n int) *Node {
	var head *Node
	for n/10 != 0 {
		head = AddNode(head, n%10)
		n = n / 10
	}
	head = AddNode(head, n%10)
	return head
}

func reverseList(head *Node) *Node {
	curr := head
	var prev *Node = nil
	var next *Node = nil
	for curr != nil {
		next = curr.next
		curr.next = prev

		prev = curr
		curr = next
	}

	return prev
}

func printList(head *Node) {
	for head != nil {
		fmt.Printf("%3d", head.val)
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

func add_numbers(headA *Node, headB *Node) *Node {
	var res *Node = nil
	carry := 0
	for headA != nil && headB != nil {
		curr := headA.val + headB.val + carry
		res = AddNode(res, curr%10)
		carry = curr / 10

		headA = headA.next
		headB = headB.next
	}
	for headA != nil {
		curr := headA.val + carry
		res = AddNode(res, curr%10)
		carry = curr / 10

		headA = headA.next
	}

	for headB != nil {
		curr := headB.val + carry
		res = AddNode(res, curr%10)
		carry = curr / 10

		headB = headB.next
	}
	if carry != 0 {
		res = AddNode(res, carry)
	}

	return res
}

func main() {
	a := readInt()
	b := readInt()
	listA := buildList(a)
	printList(listA)
	listB := buildList(b)
	printList(listB)
	// reverseListA := reverseList(listA)
	// printList(reverseListA)
	// reverseListB := reverseList(listB)
	// printList(reverseListB)

	res := add_numbers(listA, listB)
	printList(res)
}
