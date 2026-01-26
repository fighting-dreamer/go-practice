package main

import (
	"cmp"
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

type TreeNode[T any] struct {
	Val   T
	Left  *TreeNode[T]
	Right *TreeNode[T]
}

type BinaryTree[T cmp.Ordered] struct {
	Size    int
	Head    *TreeNode[T]
	nullVal T
}

func NewBinaryTree[T cmp.Ordered](nullVal T) *BinaryTree[T] {
	return &BinaryTree[T]{
		Size:    0,
		nullVal: nullVal,
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
	if val == t.nullVal {
		return
	}
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

func count_good_nodes_helper(head *TreeNode[int], maxOnPath int) int {
	if head == nil {
		return 0
	}
	res := 0
	if head.Val >= maxOnPath {
		res++
		maxOnPath = head.Val
	}

	res += count_good_nodes_helper(head.Left, maxOnPath)
	res += count_good_nodes_helper(head.Right, maxOnPath)

	return res
}

func count_good_nodes(head *TreeNode[int]) int {
	return count_good_nodes_helper(head, 0)
}

func ExampleBinaryTree() *BinaryTree[int] {
	bt := NewBinaryTree[int](-1)
	bt.AddNode(3)
	bt.AddNode(1)
	bt.AddNode(4)
	bt.AddNode(3)
	bt.AddNode(-1)
	bt.AddNode(1)
	bt.AddNode(5)

	return bt
}

func ExampleBinaryTree2() *BinaryTree[int] {
	bt := NewBinaryTree[int](-1)
	bt.AddNode(3)
	bt.AddNode(3)
	bt.AddNode(-1)
	bt.AddNode(4)
	bt.AddNode(2)

	return bt
}

func main() {
	bt := ExampleBinaryTree2()
	fmt.Printf("count_good_nodes(bt.Head): %v\n", count_good_nodes(bt.Head))
}
