package main

import "fmt"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func post_traversal(head *Node, level int, columnNodeMap map[int]*Node) {
	if head == nil {
		return
	}
	// we can go to "right first then left" OR "left first, then right" , both are correct.
	post_traversal(head.Left, level-1, columnNodeMap)
	post_traversal(head.Right, level+1, columnNodeMap)

	if _, ok := columnNodeMap[level]; !ok {
		columnNodeMap[level] = head
	}

	return
}

func bottom_view(head *Node) []int {
	columnNodeMap := map[int]*Node{}
	level := 0

	post_traversal(head, level, columnNodeMap)
	res := []int{}
	for _, v := range columnNodeMap {
		res = append(res, v.Val)
	}

	return res
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
	head.Right.Left = new_node(6)
	head.Right.Right = new_node(7)
	head.Right.Right.Right = new_node(8)
	fmt.Println("----------")
	fmt.Println(bottom_view(head))
}
