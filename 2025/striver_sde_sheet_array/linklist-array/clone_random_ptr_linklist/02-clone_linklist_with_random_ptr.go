package main

import (
	"fmt"
	"log"
	"math/rand"
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
	val    int
	next   *Node
	random *Node
}

func NewNode(val int) *Node {
	return &Node{
		val:    val,
		next:   nil,
		random: nil,
	}
}

func addLinklist(head *Node, valNode *Node) *Node {
	if head == nil {
		head = valNode
	} else {
		curr := head
		for curr.next != nil {
			curr = curr.next
		}
		curr.next = valNode
	}
	return head
}

func printLinklist(head *Node) {
	for head != nil {
		fmt.Printf("%d ", head.val)
		if head.random != nil {
			fmt.Printf("%d \n", head.random.val)
		} else {
			fmt.Printf("nil \n")
		}
		fmt.Println("----------------")
		head = head.next

	}
	fmt.Println()
}

func buildRandomPtrLinklist(s string) *Node {
	var head *Node = nil
	for i := 0; i < len(s); i++ {
		valNode := NewNode(int(s[i]) - int('0'))
		head = addLinklist(head, valNode)
	}
	printLinklist(head)

	mp := map[int]*Node{}
	curr := head
	it := 1
	for curr != nil {
		mp[it] = curr
		curr = curr.next
		it++
	}

	curr = head
	for curr != nil {
		curr.random = mp[rand.Intn(int(float64(it)*1.5))]
		curr = curr.next
	}

	return head
}

func clonLinkListWithRandomPtr(head *Node) *Node {
	var cloneHead *Node = nil
	mp := map[*Node]*Node{}
	curr := head
	for curr != nil {
		valNode := NewNode(curr.val)
		cloneHead = addLinklist(cloneHead, valNode)
		mp[curr] = valNode
		curr = curr.next
	}

	currClone := cloneHead
	curr = head
	for curr != nil {
		currClone.random = mp[curr.random]
		curr = curr.next
		currClone = currClone.next
	}

	return cloneHead
}

func main() {
	s := readString()
	head := buildRandomPtrLinklist(s)
	printLinklist(head)
	fmt.Println("-----------cloned------------")
	cloneHead := clonLinkListWithRandomPtr(head)
	printLinklist(cloneHead)
}
