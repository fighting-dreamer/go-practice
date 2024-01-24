package main

import (
	"fmt"
	"log"

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

func remove_nth_node_from_end[T any](l *dsa.List[T], n int) {
	slow, fast := l.Head, l.Head
	i := 0
	for i = 0; i < n && fast != nil; i++ {
		fast = fast.Next
	}
	// fmt.Println(i, fast == nil)

	// case : n = len(l) => remove the first node and shift head pointer.
	if i == n && fast == nil {
		l.Head = l.Head.Next
		return
	}
	// case : n > len(l) => no node to be removed
	if i != n && fast == nil {
		return
	}

	// case : n < len(l) => node to prev of slow to connect to next of slow and remove slow.
	var prev *dsa.Node[T]
	for fast != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next
	}
	prev.Next = slow.Next

	return
}

func main() {
	count := readInt()
	n := readInt()
	l := dsa.NewLinkList[int]()

	for i := 0; i < count; i++ {
		l.Add(readInt())
	}
	l.Print()
	remove_nth_node_from_end(l, n)
	l.Print()
}
