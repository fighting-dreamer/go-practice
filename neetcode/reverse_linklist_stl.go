package main

import (
	"fmt"
	"log"

	"github.com/liyue201/gostl/ds/list/simplelist"
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

func print_list[T any](l *simplelist.List[T]) {
	l.Traversal(func(value T) bool {
		fmt.Println(value)
		return true
	})
}

func length[T any](l *simplelist.List[T]) int {
	return l.Len()
}

func reverse_linklist[T any](l *simplelist.List[T]) {
	fmt.Println("could not Reverse : simplelist does not support setting Next value of Node directly or creating a Node directly.")
}

func main() {
	n := readInt()
	l := simplelist.New[int]()
	for i := 0; i < n; i++ {
		l.PushBack(readInt())
	}
	fmt.Printf("l: %v\n", l)
	print_list[int](l)
	fmt.Println(length[int](l))
	reverse_linklist(l)
}
