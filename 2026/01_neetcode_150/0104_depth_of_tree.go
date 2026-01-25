package main

import (
	"cmp"
	"fmt"
	"math/rand"
)

type TreeNode[T cmp.Ordered | any] struct {
	Val   T
	Left  *TreeNode[T]
	Right *TreeNode[T]
}

type BSTTree[T cmp.Ordered] struct {
	Size int
	Head *TreeNode[T]
}

func NewBinarySearchTree[T cmp.Ordered]() *BSTTree[T] {
	return &BSTTree[T]{
		Size: 0,
	}
}

func (t *BSTTree[T]) AddNode(val T) {
	defer func() { t.Size++ }()

	if t.Head == nil {
		t.Head = &TreeNode[T]{
			Val: val,
		}
	} else {
		curr := t.Head
		var old *TreeNode[T]
		for curr != nil {
			old = curr
			if curr.Val >= val {
				curr = curr.Left
			} else {
				curr = curr.Right
			}
		}
		if old.Val >= val {
			old.Left = &TreeNode[T]{
				Val: val,
			}
		} else {
			old.Right = &TreeNode[T]{
				Val: val,
			}
		}
	}

	return
}

func BstInorderTraversal[T cmp.Ordered](head *TreeNode[T]) {
	if head == nil {
		return
	}
	BstInorderTraversal(head.Left)
	fmt.Println(head.Val, " ")
	BstInorderTraversal(head.Right)
}

func PrintInorder[T cmp.Ordered](t *BSTTree[T]) {
	if t.Size == 0 {
		return
	}

	BstInorderTraversal(t.Head)
}

func height(head *TreeNode[int]) int {
	if head == nil {
		return 0
	}
	ld := height(head.Left)
	rd := height(head.Right)
	return 1 + max(ld, rd)
}

func main() {
	bst := NewBinarySearchTree[int]()
	for i := 0; i < 20; i++ {
		n := rand.Intn(1000)
		fmt.Printf("%d,", n)
		bst.AddNode(n)
	}
	fmt.Println()
	fmt.Printf("height(bst.Head): %v\n", height(bst.Head))
}
