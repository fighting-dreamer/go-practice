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

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func check_balanced(head *Node) (int, bool) {
	if head == nil {
		return 0, true
	}

	l_h, l_ok := check_balanced(head.Left)
	if l_ok == false {
		return -1, false
	}
	r_h, r_ok := check_balanced(head.Right)
	if r_ok == false {
		return -1, false
	}

	if abs(l_h-r_h) > 1 {
		return -1, false
	}
	return max(l_h, r_h) + 1, true
}

func main() {
	head := new_node(1)
	head.Left = new_node(2)
	head.Right = new_node(3)
	head.Left.Left = new_node(4)
	head.Left.Right = new_node(5)
	// head.Right.Left = new_node(6)
	head.Right.Right = new_node(7)
	head.Left.Left.Right = new_node(8)
	head.Left.Left.Right.Right = new_node(8)

	fmt.Println(check_balanced(head))
}
