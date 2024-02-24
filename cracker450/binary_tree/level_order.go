package main

import "fmt"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func level_order(head *Node) map[int][]int {
	if head == nil {
		return map[int][]int{}
	}

	res := map[int][]int{} // per level result.
	q := []*Node{}
	currLevel := 0
	q = append(q, head)
	q = append(q, nil)
	for len(q) > 1 {
		top := q[0]
		q = q[1:]
		if top == nil {
			// we have come to an end for a level
			q = append(q, nil)
			currLevel++
			continue
		}
		// fmt.Println(top.Val)
		res[currLevel] = append(res[currLevel], top.Val)
		if top.Left != nil {
			q = append(q, top.Left)
		}
		if top.Right != nil {
			q = append(q, top.Right)
		}
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
	fmt.Println("----------")
	fmt.Println(level_order(head))
}
