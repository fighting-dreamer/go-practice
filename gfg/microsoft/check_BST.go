package main

import (
	"fmt"
	"math"
)

type Node struct {
	val   int
	left  *Node
	right *Node
}

func new_node(val int) *Node {
	return &Node{
		val:   val,
		left:  nil,
		right: nil,
	}
}

func bst_input() *Node {
	head := new_node(2)
	head.left = new_node(1)
	head.right = new_node(3)

	return head
}

func non_bst_input() *Node {
	head := new_node(2)
	head.left = new_node(1)
	head.right = new_node(3)
	head.right.right = new_node(1)

	return head
}

func input() *Node {

	return bst_input() // non_bst_input()
}

func check_bst_helper(head *Node, leftLimit, rightLimit int) bool {
	if head == nil {
		return true
	}

	if head.val >= leftLimit && head.val <= rightLimit {
		return check_bst_helper(head.left, leftLimit, head.val) && check_bst_helper(head.right, head.val, rightLimit)
	}
	return false
}

func check_bst(head *Node) bool {
	if head == nil {
		return true
	}

	leftLimit := math.MinInt
	rightLimit := math.MaxInt

	return check_bst_helper(head, leftLimit, rightLimit)
}

func main() {
	head := input()

	fmt.Println(check_bst(head))
}
