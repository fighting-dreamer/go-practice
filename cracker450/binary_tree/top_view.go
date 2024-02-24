package main

import "fmt"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func pre_traversal(head *Node, level int, columnNodeMap map[int]*Node) {
	if head == nil {
		return
	}
	if _, ok := columnNodeMap[level]; !ok {
		fmt.Println(level, head.Val)
		columnNodeMap[level] = head
	}
	pre_traversal(head.Left, level-1, columnNodeMap)

	pre_traversal(head.Right, level+1, columnNodeMap)
}

func top_view(head *Node) []int {
	columnNodeMap := map[int]*Node{}
	level := 0
	pre_traversal(head, level, columnNodeMap) // Preorder traversal will work
	res := []int{}
	fmt.Println(len(columnNodeMap))
	for _, v := range columnNodeMap {
		res = append(res, v.Val)
	}

	fmt.Println(res)

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
	fmt.Println(top_view(head))
}
