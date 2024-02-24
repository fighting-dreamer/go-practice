package main

import "fmt"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func diameter_helper(head *Node) (int, int) {
	if head == nil {
		return 0, 0
	}

	ld, lh := 0, 0
	if head.Left != nil {
		ld, lh = diameter_helper(head.Left)
	}
	rd, rh := 0, 0
	if head.Right != nil {
		rd, rh = diameter_helper(head.Right)
	}
	d, h := max(max(ld, rd), lh+rh+1), max(lh, rh)+1
	return d, h
}

func diameter(head *Node) int {
	d, _ := diameter_helper(head)
	return d
}

func new_node(val int) *Node {
	return &Node{
		Val:   val,
		Left:  nil,
		Right: nil,
	}
}

func main() {
	head := new_node(1)
	head.Left = new_node(2)
	head.Right = new_node(3)
	head.Left.Left = new_node(4)
	head.Left.Right = new_node(5)
	head.Left.Right.Right = new_node(6)
	head.Left.Right.Right.Right = new_node(7)
	head.Left.Right.Right.Right.Right = new_node(8)

	fmt.Println(diameter(head))

}
