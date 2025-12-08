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

func buildSimpleLinklist(s string) *Node {
	var head *Node = nil
	for i := len(s) - 1; i >= 0; i-- {
		head = addLinklist(head, int(s[i])-int('0'))
	}

	return head
}

func getLastNode(head *Node) *Node {
	for head.next != nil {
		head = head.next
	}
	return head
}

func getNthNode(head *Node, n int) *Node {
	for i := 0; i < n-1; i++ {
		head = head.next
	}
	return head
}

func buildCycleLinkList(s string, cycleNodeCount int) *Node {
	head := buildSimpleLinklist(s)

	curr := getNthNode(head, cycleNodeCount)
	fmt.Printf("%d node : %d \n", cycleNodeCount, curr.val)
	end := getLastNode(head)
	fmt.Printf("node : %d \n", end.val)

	end.next = curr

	return head
}

func checkIfCycleExist(head *Node) (*Node, bool) {
	slow, fast := head, head
	isCyclic := false

	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
		if slow == fast {
			//cycle exist
			isCyclic = true

			slow = head
			for slow != fast {
				slow = slow.next
				fast = fast.next
			}

			return slow, isCyclic
		}
	}
	return nil, isCyclic
}

func nocycle() {
	s := readString()
	head := buildSimpleLinklist(s)
	cyclePoint, iscyclic := checkIfCycleExist(head)
	fmt.Println(cyclePoint, iscyclic)
}

func withCycle() {
	s := readString()
	k := readInt()
	head := buildCycleLinkList(s, k%len(s))
	cyclePoint, iscyclic := checkIfCycleExist(head)
	fmt.Println(cyclePoint, iscyclic)
}

func main() {
	// no cycle
	nocycle()
	// with cycle
	withCycle()
}
