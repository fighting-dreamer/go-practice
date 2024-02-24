package main

import (
	"fmt"

	"github.com/liyue201/gostl/ds/queue"
)

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func right_view(head *Node) []int {
	// if the next element in level order traversal marks the end of that level.
	// OR
	// we can keep track of an old value, once we encounter a new level Or end of iteration over tree => the old value is the end or the right side of each level.

	q := queue.New[*Node]()
	q.Push(head)
	q.Push(nil)

	var old, top *Node
	res := []int{}

	for q.Size() > 1 {
		top = q.Pop()

		if top == nil {
			// the old value is the last of this level.
			res = append(res, old.Val)
			q.Push(nil)
			continue
		}

		if top.Left != nil {
			q.Push(top.Left)
		}

		if top.Right != nil {
			q.Push(top.Right)
		}
		old = top
	}
	res = append(res, old.Val)

	return res
}

func new_node(Val int) *Node {
	return &Node{
		Val:   Val,
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
	head.Left.Left.Right = new_node(8)

	fmt.Println(right_view(head))
}
