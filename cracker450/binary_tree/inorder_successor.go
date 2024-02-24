package main

import (
	"fmt"
	"math/rand"
)

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

func AddNode(head *Node, val int) *Node {
	// fmt.Println(val)
	if head == nil {
		return new_node(val)
	} else {
		temp, old := head, (*Node)(nil)

		for temp != nil {
			// fmt.Println("curr : ", temp.Val)
			old = temp
			if temp.Val < val {
				temp = temp.Right
			} else {
				temp = temp.Left
			}
		}

		if old.Val > val {
			old.Left = new_node(val)
		} else {
			old.Right = new_node(val)
		}
		return head
	}
}

func inorder(head *Node) {
	if head == nil {
		return
	}

	inorder(head.Left)
	fmt.Print(head.Val, " ")
	inorder(head.Right)

	return
}

func inorder_helper(head *Node, target int, flag *bool) *Node {
	if head == nil {
		return nil
	}

	l := inorder_helper(head.Left, target, flag)
	if l != nil {
		return l
	}
	if *flag == true {
		*flag = false
		return head
	}

	if head.Val == target {
		*flag = true
	}
	return inorder_helper(head.Right, target, flag)
}

func inorder_successor(head *Node, target int) *Node {
	flag := false
	return inorder_helper(head, target, &flag)

}

func inorder_successor_in_bst(head *Node, target int) *Node {
	succeor := (*Node)(nil)
	curr := head
	for curr != nil {

		if curr.Val == target {
			return succeor
		}
		succeor = curr
		if curr.Val > target {
			// successor it curr or in its left
			curr = curr.Left
		} else {
			curr = curr.Right
		}
	}

	return succeor
}

func main() {
	n := 10
	var head *Node
	temp := 0
	for i := 0; i < n; i++ {
		val := rand.Intn(10000)
		fmt.Println("Val : ", val)
		head = AddNode(head, val)
		// fmt.Println()
		// fmt.Println("head  ", head.Val)
		if rand.Intn(100) > 50 {
			temp = val
		}
	}
	fmt.Println()
	fmt.Println("--------------------")
	// fmt.Print(head.Val)
	// fmt.Println(head.Left.Val)
	// fmt.Println(head.Right.Val)
	// fmt.Println(head.Left.Left.Val)
	// fmt.Println(head.Left.Right.Val)
	// fmt.Println(head.Right.Left.Val)
	// fmt.Println(head.Right.Right.Val)
	inorder(head)
	fmt.Println("\n--------------------")
	fmt.Println(temp)
	fmt.Println(inorder_successor(head, temp))
	fmt.Println(inorder_successor(head, temp))

}
