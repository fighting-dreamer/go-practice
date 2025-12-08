package main

import "fmt"

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

func buildYlinklistToTest(flag bool) (*Node, *Node) {
	head1 := NewNode(1)
	head2 := NewNode(11)
	curr1 := head1
	curr2 := head2

	curr1.next = NewNode(2)
	curr1 = curr1.next

	curr1.next = NewNode(3)
	curr1 = curr1.next

	curr1.next = NewNode(4)
	curr1 = curr1.next

	curr1.next = NewNode(5)
	curr1 = curr1.next

	curr2.next = NewNode(22)
	curr2 = curr2.next

	curr2.next = NewNode(33)
	curr2 = curr2.next

	var sharedList *Node = nil
	if flag {
		for i := 0; i < 7; i++ {
			sharedList = addLinklist(sharedList, 101+i)
		}
	}
	curr1.next = sharedList
	curr2.next = sharedList

	return head1, head2
}

func lengthOfLinklist(head *Node) int {
	res := 0

	for head != nil {
		head = head.next
		res++
	}

	return res
}

func findYpointInPartialSharedLinklist(head1 *Node, head2 *Node) *Node {
	length1 := lengthOfLinklist(head1)
	length2 := lengthOfLinklist(head2)

	// just keeping head1 as one with longer linklist
	if length1 < length2 {
		head1, head2 = head2, head1
		length1, length2 = length2, length1
	}

	for i := 0; i < length1-length2; i++ {
		head1 = head1.next
	}

	for head1 != nil && head2 != nil {
		// fmt.Println(head1.val, head2.val)
		if head1 == head2 {
			// we found the common point
			return head1
		}
		head1 = head1.next
		head2 = head2.next
	}
	return nil
}

func main() {
	head1, head2 := buildYlinklistToTest(true)
	printLinklist(head1)
	printLinklist(head2)
	sharedPoint := findYpointInPartialSharedLinklist(head1, head2)
	if sharedPoint == nil {
		fmt.Println("No common point")
	} else {
		printLinklist(sharedPoint)
	}
}
