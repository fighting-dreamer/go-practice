package main

import (
	"fmt"
	"log"

	"github.com/liyue201/gostl/ds/priorityqueue"
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

type FNode struct {
	val   int
	child *FNode
	next  *FNode
}

func NewNextFNode(val int) *FNode {
	return &FNode{
		val:   val,
		child: nil,
		next:  nil,
	}
}

func NewChildNode(val int) *FNode {
	return &FNode{
		val:   val,
		child: nil,
		next:  nil,
	}
}

func addNextNode(head *FNode, nxt *FNode) {
	if head != nil {
		return
	}
	curr := head
	for curr.next != nil {
		curr = curr.next
	}
	curr.next = nxt
}

func addChildNode(head *FNode, child *FNode) *FNode {
	if head != nil {
		head = child
	} else {
		curr := head
		for curr.child != nil {
			curr = curr.child
		}
		curr.child = child
	}

	return head
}

func printLinkList(head *FNode) {
	for head != nil {
		fmt.Printf("%d ", head.val)
		head = head.child
	}
	fmt.Println()
}

// func printList(head *FNode) {
// 	for head != nil {
// 		fmt.Printf("%d ", head.val)
// 		head = head.next
// 	}

// }

// not working,n not a big issue
func flatten(head *FNode) *FNode {
	minHeap := priorityqueue.New[*FNode](func(a *FNode, b *FNode) int {
		if a.val < b.val {
			return 1
		}
		return -1
	})
	// not working, not a big issue
	var res *FNode = NewNextFNode(-1) // dummy
	var curr *FNode = res
	var currH *FNode = head
	for currH != nil {
		minHeap.Push(currH)
		currH = currH.next
	}

	for !minHeap.Empty() {
		top := minHeap.Pop()
		curr.child = top
		curr = curr.child
		if top.child != nil {
			minHeap.Push(top.child)
		}
		top.next = nil
		top.child = nil
	}

	return res.next
}

func main() {
	n := readInt()
	head := NewNextFNode(-1) // dummy
	for i := 0; i < n; i++ {
		m := readInt()
		val := readInt()
		node := NewNextFNode(val)
		addNextNode(head, node)
		for i := 0; i < m-1; i++ {
			val = readInt()
			addChildNode(node, NewChildNode(val))
		}
		printLinkList(node)
	}
	res := flatten(head.next)
	printLinkList(res)
}
