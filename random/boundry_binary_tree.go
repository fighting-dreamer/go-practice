package main

import "fmt"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func leftBoundry(head *Node) map[*Node]int {
	res := map[*Node]int{}
	// level := 0
	for head.Left != nil {
		res[head] = 1
		head = head.Left
	}

	return res
}

func level_order_leaf_and_boundry(head *Node) map[*Node]int {
	if head == nil {
		return map[*Node]int{}
	}
	resStart := map[int]*Node{}
	resEnd := map[int]*Node{}
	resLeaf := map[int][]*Node{}
	resCandidates := map[*Node]int{}

	level := 0
	q := []*Node{}
	q = append(q, head)
	q = append(q, nil)

	for len(q) > 1 {
		top := q[0]
		q = q[1:]

		if top == nil {
			level++
			q = append(q, nil)
			continue
		}
		if _, ok := resStart[level]; !ok {
			resStart[level] = top
			// fmt.Println(level, resStart)
		}
		if _, lok := resStart[level]; lok {
			if _, ok := resEnd[level]; !ok {
				resEnd[level] = top
			}
		}

		if top.Left == nil && top.Right == nil {
			resLeaf[level] = append(resLeaf[level], top)
		}

		if top.Left != nil {
			q = append(q, top.Left)
		}
		if top.Right != nil {
			q = append(q, top.Right)
		}
	}
	fmt.Println("-------------")
	fmt.Println(resStart)
	fmt.Println(resCandidates)
	fmt.Println(resEnd)
	fmt.Println("------------")
	for _, val := range resStart {
		resCandidates[val] = 1
	}
	for _, v := range resLeaf {
		for _, val := range v {
			resCandidates[val] = 1
		}
	}
	for _, v := range resEnd {
		resCandidates[v] = 1
	}

	return resCandidates
}

func rightBoundry(head *Node) map[*Node]int {
	res := map[*Node]int{}
	for head.Right != nil {
		res[head] = 1
		head = head.Right
	}

	return res
}

func preorder_helper(head *Node, candidates map[*Node]int, res []int, k *int) {
	if head == nil {
		return
	}

	if _, ok := candidates[head]; ok {
		res[*k] = head.Val
		*k = *k + 1
	}
	if head.Left != nil {
		preorder_helper(head.Left, candidates, res, k)
	}

	if head.Right != nil {
		preorder_helper(head.Right, candidates, res, k)
	}
}

func preorder(head *Node, candidates map[*Node]int) []int {
	res := make([]int, len(candidates))
	k := 0
	preorder_helper(head, candidates, res, &k)
	return res
}

func boundry_of_tree(head *Node) []int {
	if head == nil {
		return []int{}
	}
	leftB := leftBoundry(head)
	fmt.Println(leftB)
	candidates := level_order_leaf_and_boundry(head)
	fmt.Println(candidates)
	rightB := rightBoundry(head)
	fmt.Println(rightB)

	for k, _ := range leftB {
		candidates[k] = 1
	}

	for k, _ := range rightB {
		candidates[k] = 1
	}

	res := preorder(head, candidates)
	return res
}

func new_node(val int) *Node {
	return &Node{
		Val:   val,
		Left:  nil,
		Right: nil,
	}
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

func main() {
	head := new_node(1)
	head.Left = new_node(2)
	head.Right = new_node(3)
	head.Left.Left = new_node(4)
	head.Left.Right = new_node(5)
	head.Right.Left = new_node(6)
	head.Right.Right = new_node(7)
	head.Left.Right.Left = new_node(8)
	head.Left.Right.Right = new_node(9)
	fmt.Println(level_order(head))
	fmt.Println(boundry_of_tree(head))
}
