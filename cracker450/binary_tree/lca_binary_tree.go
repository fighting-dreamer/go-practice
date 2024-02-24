package main

import "fmt"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func new_node(val int) *Node {
	return &Node{
		Val:   val,
		Left:  nil,
		Right: nil,
	}
}

func lca_helper(head *Node, val1, val2 int, val1Head, val2Head *Node) {
	if head == nil {
		return
	}

	if head.Val == val1 && val1Head == nil {
		val1Head = head
	}
	if head.Val == val2 && val2Head == nil {
		val2Head = head
	}
	if val1Head != nil && val2Head != nil {
		val1Head = head
		val2Head = head
	}
	lca_helper(head.Left, val1, val2, val1Head, val2Head)
	lca_helper(head.Right, val1, val2, val1Head, val2Head)

	return
}

func lca(head *Node, val1, val2 int) *Node {
	if head == nil {
		return nil
	}
	var h1, h2 *Node
	lca_helper(head, val1, val2, h1, h2)
	if h1 == h2 {
		return h1 // the LCA
	}
}

func main() {
	head := new_node(1)
	head.Left = new_node(2)
	head.Right = new_node(3)
	head.Left.Left = new_node(4)
	head.Left.Right = new_node(5)
	head.Right.Left = new_node(6)
	head.Right.Right = new_node(7)
	head.Left.Left.Right = new_node(8)

	fmt.Println(lca(head, 4, 7).Val)
	fmt.Println(lca(head, 7, 122))
}
