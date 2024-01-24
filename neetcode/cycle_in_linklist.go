package main

import (
	"fmt"
	"log"
	"math/rand"

	"nipun.io/coding/dsa"
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

func tail_node[T any](l *dsa.List[T]) *dsa.Node[T] {
	tail := l.Head
	for tail.Next != nil {
		tail = tail.Next
	}

	return tail
}

func nth_node[T any](l *dsa.List[T], n int) *dsa.Node[T] {
	temp := l.Head

	for n != 0 {
		temp = temp.Next
		n--
	}

	return temp
}

func print_list_till_n[T any](l *dsa.List[T], n int) {
	head := l.Head
	for n != 0 && head != nil {
		fmt.Print(head.Val, " ")
		head = head.Next
		n--
	}

	fmt.Println()
}

func cycle_position[T any](l *dsa.List[T]) int {
	slow, fast := l.Head, l.Head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			// cycle detected.
			break
		}
	}
	if fast == nil || fast.Next == nil {
		return -1 // no cycle
	}

	// cycle detected
	head := l.Head
	pos := 0
	for slow != head {
		slow = slow.Next
		head = head.Next
		pos++
	}
	// the start of cycle is at head /slow node as both are same now.

	return pos
}

func main() {
	n := readInt()

	l := dsa.NewLinkList[int]()
	for i := 0; i < n; i++ {
		l.Add(readInt())
	}
	// Add Cycle :
	pos := rand.Int() % n // the node to which the tail of linklist will connect
	fmt.Println(pos)
	l.Print()
	tail := tail_node[int](l)

	tail.Next = nth_node[int](l, pos)
	print_list_till_n[int](l, n+1)

	fmt.Println(cycle_position(l))
}
