package main

import (
	"fmt"
	"log"
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

func maxPathSum(head *Node[int]) (int, int) {
	if head == nil {
		return 0, 0
	}

	maxRootPathSumLeft, maxPathSumLeft := maxPathSum(head.Left)
	maxRootPathSumRight, maxPathSumRight := maxPathSum(head.Right)

	maxRootPathSum := max(maxRootPathSumLeft, maxRootPathSumRight) + head.Val
	maxPathSum := max(maxPathSumLeft+maxPathSumRight+head.Val, max(maxPathSumLeft, maxPathSumRight))

	return maxRootPathSum, maxPathSum
}

func MaxPathSum(head *Node[int]) int {
	_, res := maxPathSum(head)
	return res
}

// diameter 5 for 1 2 3 4 5 => 4
// diameter 15 for 1 2 3 4 5 -1 6 -1 -1 7 8 -1 -1 -1 9  => 7
func main() {
	n := readInt()
	// fmt.Println(bitCount(7))
	var head *Node[int]
	var t int
	for i := 0; i < n; i++ {
		t = readInt()
		if t == -1 {
			continue
		}
		head = AddTreeNode(head, t, i)
	}

	// PrintInorder(head)
	// fmt.Println("\n\n")
	// fmt.Println(Diameter(head))
	fmt.Println(MaxPathSum(head))
}
