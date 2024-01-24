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

func reverse_linklist[T any](head *dsa.Node[T]) *dsa.Node[T] {
	curr := head
	var prev, next *dsa.Node[T]
	for curr != nil {
		next = curr.Next
		curr.Next = prev

		prev = curr
		curr = next
	}

	return prev
}

func reorder_linklist[T any](l *dsa.List[T]) {
	slow, fast := l.Head, l.Head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	reversed := reverse_linklist[T](slow.Next)
	slow.Next = nil // halfing the linklist after slow's node

	curr := l.Head

	for curr != nil && reversed != nil {
		next := curr.Next
		rNext := reversed.Next

		curr.Next = reversed
		reversed.Next = next

		curr = next
		reversed = rNext
	}
}

func main() {
	n := readInt()

	l := dsa.NewLinkList[int]()
	for i := 0; i < n; i++ {
		l.Add(readInt())
	}
	fmt.Println("Before Reorder")
	l.Print()
	reorder_linklist(l)
	fmt.Println("After Reorder")
	l.Print()

}
