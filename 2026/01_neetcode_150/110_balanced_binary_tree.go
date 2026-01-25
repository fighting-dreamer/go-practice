package main

import (
	"cmp"
	"fmt"
	"math"
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

func isBalanced_helper(head *TreeNode[int]) (int, bool) {
	if head == nil {
		return 0, true
	}

	lh, isLeftSideBalanced := isBalanced_helper(head.Left)
	if !isLeftSideBalanced {
		return -1, false
	}
	rh, isRightSideBalanced := isBalanced_helper(head.Right)
	if !isRightSideBalanced {
		return -1, false
	}

	h := max(lh, rh) + 1
	isCurrBalanced := math.Abs(float64(lh-rh)) <= 1

	return h, isLeftSideBalanced && isRightSideBalanced && isCurrBalanced
}

func isBalanced(head *TreeNode[int]) bool {
	_, res := isBalanced_helper(head)
	return res
}

func BalancedSampleBST() *BSTTree[int] {
	bst := NewBinarySearchTree[int]()
	bst.AddNode(30)
	bst.AddNode(10)
	bst.AddNode(50)
	bst.AddNode(40)
	bst.AddNode(70)

	return bst
}

func main() {
	bst := NewBinarySearchTree[int]()
	for i := 0; i < 20; i++ {
		n := rand.Intn(1000)
		fmt.Printf("%d,", n)
		bst.AddNode(n)
	}
	fmt.Println()

	fmt.Printf("isBalanced(bst.Head): %v\n", isBalanced(bst.Head))

	// -------------
	fmt.Println("Blanaced one testing")
	bst = BalancedSampleBST()
	// PrintInorder(bst)
	fmt.Printf("isBalanced(bst.Head): %v\n", isBalanced(bst.Head))
}
