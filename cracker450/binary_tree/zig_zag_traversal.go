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

func reverse(arr []int) []int {
	n := len(arr)
	for i := 0; i < n/2; i++ {
		temp := arr[i]
		arr[i] = arr[(n-1)-i]
		arr[(n-1)-i] = temp
	}

	return arr
}

func zigzag_traversal(head *Node) []int {
	currLevel := []int{}
	level := 0
	res := []int{}
	q := queue.New[*Node]()

	q.Push(head)
	q.Push(nil)

	for q.Size() > 1 {
		top := q.Pop()

		if top == nil {
			q.Push(nil)
			if level%2 == 0 {
				res = append(res, currLevel...)
			} else {
				res = append(res, reverse(currLevel)...)
			}
			currLevel = []int{}
			level++
			continue
		}

		currLevel = append(currLevel, top.Val)

		if top.Left != nil {
			q.Push(top.Left)
		}

		if top.Right != nil {
			q.Push(top.Right)
		}
	}
	if level%2 == 0 {
		res = append(res, currLevel...)
	} else {
		res = append(res, reverse(currLevel)...)
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
	head.Right.Right.Left = new_node(8)
	head.Right.Right.Right = new_node(9)
	fmt.Println("----------")
	fmt.Println(zigzag_traversal(head))
}
