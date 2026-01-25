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

func (t *BSTTree[T]) Search(target T) *TreeNode[T] {
	head := t.Head
	for head != nil {
		if head.Val == target {
			return head
		} else {
			if head.Val < target {
				head = head.Right
			} else {
				head = head.Left
			}
		}
	}

	return nil
}

func matchSubTree[T cmp.Ordered](head *TreeNode[T], subtree *TreeNode[T]) bool {
	if head == nil && subtree != nil {
		return false
	}
	if head == nil && subtree == nil {
		return true
	}
	if head != nil && subtree == nil {
		return true
	}

	return head.Val == subtree.Val && matchSubTree(head.Left, subtree.Left) && matchSubTree(head.Right, subtree.Right)
}

func checkIsSubtree[T cmp.Ordered](head *TreeNode[T], subtree *TreeNode[T]) bool {
	for head != nil {
		if head.Val == subtree.Val {
			if matchSubTree(head, subtree) {
				return true
			}
		}

		if head.Val > subtree.Val {
			head = head.Left
		} else {
			head = head.Right
		}
	}

	return false
}

func createBST[T cmp.Ordered](cnt int) (*BSTTree[T], []int) {
	bst := NewBinarySearchTree[T]()
	nums := []int{}
	for i := 0; i < cnt; i++ {
		n := rand.Intn(1000)
		fmt.Printf("%2d,", n)
		nums = append(nums, n)
		bst.AddNode(T(n))
	}
	fmt.Println()

	return bst, nums
}

func main() {
	// bst := NewBinarySearchTree[int]()
	// nums := []int{}
	// for i := 0; i < 20; i++ {
	// 	n := rand.Intn(1000)
	// 	fmt.Printf("%2d,", n)
	// 	nums = append(nums, n)
	// 	bst.AddNode(n)
	// }

	bst, nums := createBST[int](20)
	fmt.Println()
	if rand.Intn(1000)&1 == 0 {
		fmt.Println("Expected it to contain matching subtree")
		target := nums[rand.Intn(20)]
		fmt.Println("Target : ", target)
		subtree := bst.Search(target)
		fmt.Println("expected subtree.")
		fmt.Printf("checkIsSubtree(bst.Head, subtree): %v\n", checkIsSubtree(bst.Head, subtree))
	} else {
		subtree, _ := createBST[int](10)

		fmt.Println("expected not subtree.")
		fmt.Printf("checkIsSubtree(bst.Head, subtree.Head): %v\n", checkIsSubtree(bst.Head, subtree.Head))
	}
}
