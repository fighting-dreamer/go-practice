package dsa

import "fmt"

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
