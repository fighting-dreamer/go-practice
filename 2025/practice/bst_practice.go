package main

import (
	"fmt"
	"math/rand"
	"slices"
)

type BSTNode struct {
	val   int
	left  *BSTNode
	right *BSTNode
}

func NewBSTNode(val int) *BSTNode {
	return &BSTNode{
		val:   val,
		left:  nil,
		right: nil,
	}
}

func AddNode(head *BSTNode, val int) *BSTNode {
	if head == nil {
		head = NewBSTNode(val)
	} else {
		var old *BSTNode
		curr := head
		for curr != nil {
			old = curr
			if curr.val < val {
				curr = curr.right
			} else {
				curr = curr.left
			}
		}
		if old.val < val {
			old.right = NewBSTNode(val)
		} else {
			old.left = NewBSTNode(val)
		}
	}
	return head
}

func print_inorder(head *BSTNode) {
	if head == nil {
		return
	}
	print_inorder(head.left)
	fmt.Printf("%d ", head.val)
	print_inorder(head.right)
}

func level_order_traveral(head *BSTNode) {
	if head == nil {
		return
	}
	q := []*BSTNode{}
	q = append(q, head)
	q = append(q, nil)
	currLevel := 0
	fmt.Println("\nCurrnet LeveL : ", currLevel)
	for len(q) > 1 {
		top := q[0]
		q = q[1:]
		if top == nil {
			currLevel++
			fmt.Println("\nCurrnet LeveL : ", currLevel)
			q = append(q, nil)
			fmt.Println()
			continue
		}
		fmt.Printf("%d ", top.val)
		if top.left != nil {
			q = append(q, top.left)
		}

		if top.right != nil {
			q = append(q, top.right)
		}
	}
}

func zigzagTraversal(head *BSTNode) {
	res := [][]int{}
	currLevel := []int{}
	q := []*BSTNode{}
	q = append(q, head)
	q = append(q, nil)
	level := 0

	for len(q) > 1 {
		top := q[0]
		q = q[1:]

		if top == nil {
			dst := make([]int, len(currLevel))
			copy(dst, currLevel)
			if level&1 == 1 {
				slices.Reverse(dst)
			}
			res = append(res, dst)
			level++
			currLevel = []int{}
			q = append(q, nil)
			continue
		}
		currLevel = append(currLevel, top.val)
		if top.left != nil {
			q = append(q, top.left)
		}

		if top.right != nil {
			q = append(q, top.right)
		}
	}
	for i := 0; i < len(res); i++ {
		fmt.Println("currLevel : ", i)
		for j := 0; j < len(res[i]); j++ {
			fmt.Printf("%d ", res[i][j])
		}
		fmt.Println()
	}
}

func vertical_order_helper(head *BSTNode, m map[int][]int, level int) {
	if head == nil {
		return
	}
	m[level] = append(m[level], head.val)
	vertical_order_helper(head.left, m, level-1)
	vertical_order_helper(head.right, m, level+1)
}

func vertical_traversal(head *BSTNode) {
	m := map[int][]int{}

	vertical_order_helper(head, m, 0)
	for k, v := range m {
		fmt.Println(k, v)
	}
}

func isLeaf(head *BSTNode) bool {
	return head.left == nil && head.right == nil
}

func leftView(head *BSTNode, res []int) []int {
	if head == nil || isLeaf(head) {
		return res
	}
	res = append(res, head.val)
	if head.left != nil {
		res = leftView(head.left, res)
	} else {
		res = leftView(head.right, res)
	}
	return res
}

func rightView(head *BSTNode, res []int) []int {
	if head == nil || isLeaf(head) {
		return res
	}
	res = append(res, head.val)
	if head.right != nil {
		res = rightView(head.right, res)
	} else {
		res = rightView(head.left, res)
	}
	return res
}

func leafNodes(head *BSTNode, res []int) []int {
	if head == nil {
		return res
	}
	if isLeaf(head) {
		res = append(res, head.val)
		return res
	}
	if head.left != nil {
		res = leafNodes(head.left, res)
	}
	if head.right != nil {
		res = leafNodes(head.right, res)
	}
	return res
}

func boundry_traversal(head *BSTNode) {
	res := []int{}
	// left view
	fmt.Println("Left view : ")
	lV := leftView(head, []int{})
	res = append(res, lV...)
	fmt.Println(lV)
	// leaf nodes
	fmt.Println("Leaf Nodes : ")
	leafN := leafNodes(head, []int{})
	res = append(res, leafN...)
	fmt.Println(leafN)
	// right view
	fmt.Println("Right Nodes : ")
	rV := rightView(head, []int{})
	res = append(res, rV...)
	fmt.Println(rV)

	fmt.Println("Complete : ")
	fmt.Println(res)
}

func main() {
	var head *BSTNode
	for i := 0; i < 20; i++ {
		n := -99 + rand.Intn(198)
		fmt.Printf("%d,", n)
		head = AddNode(head, n)
	}
	fmt.Println("\nInorder : ")
	print_inorder(head)
	fmt.Println("\nlevel order : ")
	level_order_traveral(head)

	fmt.Println("\n Zig zag traversal \n")
	zigzagTraversal(head)

	fmt.Println("\n Vertial order : ")
	vertical_traversal(head)

	fmt.Println("\nBoundry Traversal : ")
	boundry_traversal(head)
}
