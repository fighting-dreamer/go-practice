package main

import (
	"fmt"
	"log"
	"math"
)

type Node[T any] struct {
	Val   T
	Left  *Node[T]
	Right *Node[T]
}

func bitCount(c int) int {
	k := 0
	// 8(1000) => 4
	// 7(111) => 3
	// 11(1011) => 4
	for c >= (1 << k) {
		k++
	}

	return k
}

func isNthBitSet(n, bits int) int {
	return (n >> bits) & 1
}

func AddTreeNode[T any](head *Node[T], val T, count int) *Node[T] {
	if count == 0 {
		return &Node[T]{
			Val: val,
		}
	}
	// count should start with 1
	count++ // coz input count start with 0
	// 11 = 1011 => left -> right -> right(new-node)
	bits := bitCount(count) - 1
	curr := head
	// fmt.Println(bits)
	for bits > 1 {
		// fmt.Println(count, " : ", curr.Val)
		if isNthBitSet(count, bits-1) == 1 {
			curr = curr.Right
		} else {
			curr = curr.Left
		}
		bits--
	}

	if count&1 == 1 {
		curr.Right = &Node[T]{
			Val: val,
		}
	} else {
		curr.Left = &Node[T]{
			Val: val,
		}
	}

	return head
}

func PrintInorder[T any](head *Node[T]) {
	if head == nil {
		return
	}

	PrintInorder(head.Left)
	fmt.Println(head.Val)
	PrintInorder(head.Right)
}

func diameter[T any](head *Node[T]) (int, int) {
	if head == nil {
		// no diameter or depth for nil node.
		return 0, 0
	}
	// fmt.Println(head.Val)
	depthLeft, diameterLeft := diameter(head.Left)
	depthRight, diameterRight := diameter(head.Right)
	depth := max(depthLeft, depthRight) + 1
	diameter := max(depthLeft+depthRight+1, max(diameterLeft, diameterRight))
	fmt.Println(head.Val, " : ", depth, " : ", diameter)
	return depth, diameter
}

func Diameter[T any](head *Node[T]) int {
	_, diameter := diameter[T](head)
	return diameter
}

func readInt() int {
	var n int
	_, err := fmt.Scanf("%d", &n)
	if err != nil {
		log.Fatal("could not read input number", err)
	}
	return n
}

func readString() string {
	var s string
	_, err := fmt.Scan(&s)
	if err != nil {
		log.Fatal("could not read input string", err)
	}
	return s
}

func readDouble() float64 {
	var f float64
	_, err := fmt.Scanf("%f", &f)
	if err != nil {
		log.Fatal("could not read input float", err)
	}
	return f
}

func readChar() rune {
	var r rune
	_, err := fmt.Scanf("%c", &r)
	if err != nil {
		log.Fatal("Could not read input char", err)
	}
	return r
}

func serialize_binary_tree(head *Node[int]) []*int {
	res := []*int{}

	q := []*Node[int]{}
	q = append(q, head)
	q = append(q, nil)
	for len(q) > 1 {
		front := q[0]
		q = q[1:]
		if front == nil {
			q = append(q, nil)
			continue
		}

		res = append(res, &front.Val)
		if front.Val == math.MinInt {
			continue
		}
		if front.Left != nil {
			q = append(q, front.Left)
		} else {
			q = append(q, &Node[int]{
				Val: math.MinInt,
			})
		}
		if front.Right != nil {
			q = append(q, front.Right)
		} else {
			q = append(q, &Node[int]{
				Val: math.MinInt,
			})
		}
	}
	l := len(res) - 1
	for i := l; i >= 0; i-- {
		if *res[i] != math.MinInt {
			res = res[:i+1]
			break
		}
	}

	return res
}

func main() {
	n := readInt()
	var head *Node[int]
	var t int
	for i := 0; i < n; i++ {
		t = readInt()
		if t == -1 {
			continue
		}
		head = AddTreeNode(head, t, i)
	}

	nums := serialize_binary_tree(head)
	for i := 0; i < len(nums); i++ {
		if *nums[i] == math.MinInt {
			fmt.Println("null")
			continue
		}
		fmt.Println(*nums[i])
	}
	var root *Node[int]
	for i := 0; i < len(nums); i++ {
		if *nums[i] == math.MinInt {
			continue
		}
		root = AddTreeNode(root, *nums[i], i)
	}
	nums2 := serialize_binary_tree(head)

	for i := 0; i < len(nums); i++ {
		if *nums2[i] != *nums[i] {
			fmt.Println("Some issue with serialization")
			return
		}
	}
	fmt.Println("All good")
	return
}
