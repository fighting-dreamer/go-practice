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

// diameter => edge count in the path, not the node count, otherwise d = max(ld, rd, lh + rh + 1)
func diameter_helper[T cmp.Ordered](head *TreeNode[T]) (int, int) {
	if head == nil {
		return 0, 0
	}
	ld, lh := diameter_helper(head.Left)
	rd, rh := diameter_helper(head.Right)
	d, h := max(ld, rd, lh+rh-1), max(lh, rh)+1
	return d, h
}

func diameter[T cmp.Ordered](head *TreeNode[T]) int {
	d, _ := diameter_helper(head)
	return d
}

func main() {
	bst := NewBinarySearchTree[int]()
	for i := 0; i < 20; i++ {
		n := rand.Intn(1000)
		fmt.Printf("%2d,", n)
		bst.AddNode(n)
	}
	fmt.Println()

	fmt.Printf("diameter(bst.Head): %v\n", diameter(bst.Head))
}
