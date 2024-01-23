package main

import (
	"fmt"
	"log"
)

type Node[T any] struct {
	Val  T
	Next *Node[T]
}

type List[T any] struct {
	Size int
	Head *Node[T]
}

func NewLinkList[T any]() *List[T] {
	return &List[T]{
		Size: 0,
		Head: nil,
	}
}

func (l *List[T]) Empty() bool {
	return l.Size == 0
}

func (l *List[T]) Add(val T) {
	defer func() { l.Size++ }()
	newNode := &Node[T]{
		Val:  val,
		Next: nil,
	}
	if l.Empty() {
		l.Head = newNode
		return
	}

	temp := l.Head
	var old *Node[T]
	for temp != nil {
		old = temp
		temp = temp.Next
	}
	old.Next = newNode

	return
}

func (l *List[T]) Print() {
	temp := l.Head
	for temp != nil {
		fmt.Print(temp.Val, " ")
		temp = temp.Next
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

func (l *List[T]) length() int {
	return l.Size
}

func (l *List[T]) reverse() {
	curr := l.Head
	var prev, next *Node[T]
	for curr != nil {
		next = curr.Next
		curr.Next = prev

		prev = curr
		curr = next
	}
	l.Head = prev
}

func main() {
	l := NewLinkList[int]()
	n := readInt()
	for i := 0; i < n; i++ {
		l.Add(readInt())
		// l.Print()
		// fmt.Println(l.Size)
	}
	l.reverse()
	l.Print()
}
