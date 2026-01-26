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

type TreeNode[T any] struct {
	Val   T
	Left  *TreeNode[T]
	Right *TreeNode[T]
}

type BinaryTree[T any] struct {
	Size int
	Head *TreeNode[T]
}

func NewBinaryTree[T any]() *BinaryTree[T] {
	return &BinaryTree[T]{
		Size: 0,
	}
}

func kthBit(n, k int) int {
	return (n >> k) & 1
}

func bitSize(n int) int {
	k := 0
	for n >= (1 << k) {
		k++
	}

	return k
}

func (t *BinaryTree[T]) AddNode(val T) {
	t.Size++
	if t.Size == 1 {
		t.Head = &TreeNode[T]{
			Val: val,
		}
	} else {
		// size = 2 => 10 => just to left
		// size = 3 => 11 => just to right
		// size = 4 => 100 => just to left of left
		// 5 => 101 => left -> right
		// 6 => 110 => right -> left
		// 7 => 111 => right -> right
		// 8 => 1000 => left -> left -> left
		// ... 11 => 1011 => left -> right -> right
		// size & 1 tells the node is going to be a left node or right node
		// (size >> 1) 's reverse if 1 => right, else left
		bitCount := bitSize(t.Size) - 1 // for 6 => 3 -1 = 2
		head := t.Head
		for bitCount > 1 {
			kthBit := kthBit(t.Size, bitCount-1)
			// fmt.Print(head.Val, " ")
			if kthBit == 0 {
				head = head.Left
			} else {
				head = head.Right
			}
			bitCount--
		}
		if kthBit(t.Size, 0) == 0 {
			head.Left = &TreeNode[T]{
				Val: val,
			}
		} else {
			head.Right = &TreeNode[T]{
				Val: val,
			}
		}

		fmt.Println("Last Node : ", head.Val, " ==> ", kthBit(t.Size, 0), " ==> ", val)
	}
}

func PrintInorder(head *TreeNode[int]) {
	if head == nil {
		return
	}
	PrintInorder(head.Left)
	fmt.Printf("%d, ", head.Val)
	PrintInorder(head.Right)
}

func invert_binary_tree(head *TreeNode[int]) {
	if head == nil {
		return
	}

	invert_binary_tree(head.Left)
	invert_binary_tree(head.Right)
	head.Left, head.Right = head.Right, head.Left
}

func main() {
	bt := NewBinaryTree[int]()
	n := readInt()
	for i := 0; i < n; i++ {
		bt.AddNode(rand.Intn(1000))
	}

	PrintInorder(bt.Head)
	fmt.Println()
	invert_binary_tree(bt.Head)
	PrintInorder(bt.Head)
	fmt.Println()
}
