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

/*

    	   1
       /       \
      2         3
   /     \    /    \
  4       5  6       7
   \
     8

output : 1 2 4 8
*/

func Left_view(head *Node) []int {
	// do level order traversal
	// get the first node in each level

	q := queue.New[*Node]()
	q.Push(head)
	q.Push(nil)
	flag := true
	res := []int{}

	for q.Size() > 1 {
		top := q.Pop()
		if top == nil {
			q.Push(nil)
			flag = true
			continue
		}
		if flag == true {
			// curr node is the first node for that level.

			res = append(res, top.Val)
			flag = false
		}

		if top.Left != nil {
			q.Push(top.Left)
		}

		if top.Right != nil {
			q.Push(top.Right)
		}
	}

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

	fmt.Println(Left_view(head))
}
