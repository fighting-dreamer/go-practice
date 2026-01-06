package main

import (
	"cmp"
	"fmt"
	"log"
	"math"
	"math/rand"
)

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

// ------------------------------

type TreeNode[T cmp.Ordered | any] struct {
	Val   T
	Left  *TreeNode[T]
	Right *TreeNode[T]
}

type BSTTree[T cmp.Ordered] struct {
	Size int
	Head *TreeNode[T]
}

func NewBinarySearchTree[T cmp.Ordered]() *BSTTree[T] {
	return &BSTTree[T]{
		Size: 0,
	}
}

func (t *BSTTree[T]) AddNode(val T) {
	defer func() { t.Size++ }()

	if t.Head == nil {
		t.Head = &TreeNode[T]{
			Val: val,
		}
	} else {
		curr := t.Head
		var old *TreeNode[T]
		for curr != nil {
			old = curr
			if curr.Val >= val {
				curr = curr.Left
			} else {
				curr = curr.Right
			}
		}
		if old.Val >= val {
			old.Left = &TreeNode[T]{
				Val: val,
			}
		} else {
			old.Right = &TreeNode[T]{
				Val: val,
			}
		}
	}

	return
}

type BinaryTree[T any] struct {
	Size int
	Head *TreeNode[T]
}

func NewBinaryTree[T any]() *BinaryTree[T] {
	return &BinaryTree[T]{
		Size: 0,
	}
}

func kthBit(n, k int) int {
	return (n >> k) & 1
}

func bitSize(n int) int {
	k := 0
	for n >= (1 << k) {
		k++
	}

	return k
}

func (t *BinaryTree[T]) AddNode(val T) {
	t.Size++
	if t.Size == 1 {
		t.Head = &TreeNode[T]{
			Val: val,
		}
	} else {
		// size = 2 => 10 => just to left
		// size = 3 => 11 => just to right
		// size = 4 => 100 => just to left of left
		// 5 => 101 => left -> right
		// 6 => 110 => right -> left
		// 7 => 111 => right -> right
		// 8 => 1000 => left -> left -> left
		// ... 11 => 1011 => left -> right -> right
		// size & 1 tells the node is going to be a left node or right node
		// (size >> 1) 's reverse if 1 => right, else left
		bitCount := bitSize(t.Size) - 1 // for 6 => 3 -1 = 2
		head := t.Head
		for bitCount > 1 {
			kthBit := kthBit(t.Size, bitCount-1)
			// fmt.Print(head.Val, " ")
			if kthBit == 0 {
				head = head.Left
			} else {
				head = head.Right
			}
			bitCount--
		}
		if kthBit(t.Size, 0) == 0 {
			head.Left = &TreeNode[T]{
				Val: val,
			}
		} else {
			head.Right = &TreeNode[T]{
				Val: val,
			}
		}

		// fmt.Println("Last Node : ", head.Val, " ==> ", kthBit(t.Size, 0), " ==> ", val)
	}
}

// ----------------------------

func BstInorderTraversal[T cmp.Ordered](head *TreeNode[T]) {
	if head == nil {
		return
	}
	BstInorderTraversal(head.Left)
	fmt.Printf("%d ", head.Val)
	BstInorderTraversal(head.Right)
}

func PrintInorder[T cmp.Ordered](t *BSTTree[T]) {
	if t.Size == 0 {
		return
	}

	BstInorderTraversal(t.Head)
}

// ----------------------------

func find_range_sum(head *TreeNode[int], res int, l, r int) int {
	if head == nil {
		return res
	}
	if head.Val <= r && head.Val >= l {
		res += head.Val
	}

	if head.Left != nil {
		res += find_range_sum(head.Left, res, l, r)
	}
	if head.Right != nil {
		res += find_range_sum(head.Right, res, l, r)
	}

	return res
}

func range_sum(bst *BSTTree[int]) {
	l := readInt()
	r := readInt()

	res := find_range_sum(bst.Head, 0, l, r)
	fmt.Println(res)
}

// ----------------------------

func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func find_minimum_absolute_difference(head *TreeNode[int], last **TreeNode[int], res *int) {
	if head == nil {
		return
	}
	find_minimum_absolute_difference(head.Left, last, res)
	if *last != nil {
		fmt.Println(head.Val, (*last).Val)
		*res = min(*res, Abs(head.Val-(*last).Val))
	}
	*last = head
	find_minimum_absolute_difference(head.Right, last, res)
}

func minimum_absolute_difference(head *TreeNode[int]) int {
	var last *TreeNode[int] = nil
	var res int = math.MaxInt
	find_minimum_absolute_difference(head, &last, &res)
	return res
}

// -------------------

func traverse_tree(head *TreeNode[int], fm map[int]int) {
	if head == nil {
		return
	}
	fm[head.Val]++
	traverse_tree(head.Left, fm)
	traverse_tree(head.Right, fm)
}

func find_mode(head *TreeNode[int]) int {
	fm := map[int]int{}

	traverse_tree(head, fm)
	maxV := math.MinInt
	res := -1
	for k, v := range fm {
		if v >= maxV {
			maxV = v
			res = k
		}
	}

	return res
}

// ----------------

func greater_sum_tree_code(head *TreeNode[int], res *int) {
	if head == nil {
		return
	}

	greater_sum_tree_code(head.Right, res)
	*res += head.Val
	head.Val = *res
	greater_sum_tree_code(head.Left, res)
}

func greater_sum_tree(head *TreeNode[int]) {
	// the left sub-tree will have the nodes in right subtree that would be greater than them
	// we need to update them to
	// to accumulate values from right subtree first then update head and the left subtree as w ego forward
	var res int = 0
	greater_sum_tree_code(head, &res)
	return
}

// -----------------

func preorder_Traversal(head *TreeNode[int]) {
	if head == nil {
		return
	}
	fmt.Printf("%d ", head.Val)
	preorder_Traversal(head.Left)
	preorder_Traversal(head.Right)
}

func traverse_and_store(head *TreeNode[int], arr *[]int) {
	if head == nil {
		return
	}

	traverse_and_store(head.Left, arr)
	*arr = append(*arr, head.Val)
	traverse_and_store(head.Right, arr)
}

func build_balanced_bst(bst *BSTTree[int], arr []int, start, end int) {
	if start > end {
		return
	}
	mid := (start + end) / 2
	bst.AddNode(arr[mid])
	build_balanced_bst(bst, arr, start, mid-1)
	build_balanced_bst(bst, arr, mid+1, end)
}

func build_balanced_binary_search_tree(head *TreeNode[int]) *TreeNode[int] {
	arr := []int{}
	traverse_and_store(head, &arr)
	fmt.Println("Build balanced bst : ")
	fmt.Println(arr)
	bst := NewBinarySearchTree[int]()
	build_balanced_bst(bst, arr, 0, len(arr)-1)
	BstInorderTraversal(bst.Head)
	fmt.Println()

	fmt.Println("Tree 1 : ")
	preorder_Traversal(head)
	fmt.Println()
	fmt.Println("Tree 2 : ")
	preorder_Traversal(bst.Head)
	return bst.Head
}

func lowest_common_ancestor_code(head *TreeNode[int], l int, r int) *TreeNode[int] {
	if head == nil {
		return nil
	}

	if head.Val == l || head.Val == r {
		return head
	}

	ln := lowest_common_ancestor_code(head.Left, l, r)
	rn := lowest_common_ancestor_code(head.Right, l, r)

	if ln != nil && rn != nil {
		return head // data is present in left and right
	}

	if ln != nil {
		return ln
	}
	return rn
}

func lowest_common_ancestor(head *TreeNode[int]) {
	l := readInt()
	r := readInt()

	res := lowest_common_ancestor_code(head, l, r)
	fmt.Println(res.Val)
}

func complete_binary_tree(head *TreeNode[int]) bool {
	q := []*TreeNode[int]{}
	q = append(q, head)
	nullEncountered := false
	for len(q) > 0 {
		top := q[0]
		q = q[1:]

		if top == nil {
			nullEncountered = true
		} else {
			if nullEncountered {
				return false
			}
			q = append(q, top.Left)
			q = append(q, top.Right)
		}
	}

	return true
}

func pruning_tree_code(head *TreeNode[int]) *TreeNode[int] {
	if head == nil {
		return nil
	}

	head.Left = pruning_tree_code(head.Left)
	head.Right = pruning_tree_code(head.Right)

	if head.Val == 0 && head.Left == nil && head.Right == nil {
		return nil
	}
	return head
}

func pruning_tree(head *TreeNode[int]) {
	// if the subtree on left or right contain 1 then keep it, else remove it
	fmt.Println("Before")
	BstInorderTraversal(head)
	head = pruning_tree_code(head)
	fmt.Println("After : ")
	BstInorderTraversal(head)
}

func copyArr(arr []int) []int {
	if len(arr) == 0 {
		return []int{}
	}
	return append([]int(nil), arr...)
}

func isLeaf(head *TreeNode[int]) bool {
	return head.Left == nil && head.Right == nil
}

func path_sum_2_code(head *TreeNode[int], res *[][]int, target int, sumTillnow int, curr []int) {
	if head == nil {
		return
	}

	if sumTillnow+head.Val == target {
		if !isLeaf(head) {
			return
		}
		curr = append(curr, head.Val)
		(*res) = append(*res, copyArr(curr))
		return
	}

	path_sum_2_code(head.Left, res, target, sumTillnow+head.Val, append(curr, head.Val))
	path_sum_2_code(head.Right, res, target, sumTillnow+head.Val, append(curr, head.Val))
}

func path_sum_2(head *TreeNode[int]) [][]int {
	res := [][]int{}
	target := readInt()
	path_sum_2_code(head, &res, target, 0, []int{})
	for i := 0; i < len(res); i++ {
		fmt.Println(res[i])
	}
	return res
}

func max_diff_node_and_ancestor_code(head *TreeNode[int], currMin int, currMaxx int) int {
	if head == nil {
		return currMaxx - currMin
	}
	currMaxx = max(currMaxx, head.Val)
	currMin = min(currMin, head.Val)

	leftDiff := max_diff_node_and_ancestor_code(head.Left, currMin, currMaxx)
	rightDiff := max_diff_node_and_ancestor_code(head.Right, currMin, currMaxx)
	return max(leftDiff, rightDiff)
}

func max_diff_node_and_ancestor(head *TreeNode[int]) {
	res := max_diff_node_and_ancestor_code(head, head.Val, math.MinInt)
	fmt.Println(res)
}

func diameter_code(head *TreeNode[int]) (int, int) {
	if head == nil {
		return 0, 0
	}

	ldiameter, lheight := diameter_code(head.Left)
	rdiameter, rheight := diameter_code(head.Right)
	diameter := max(max(ldiameter, rdiameter), lheight+rheight+1)
	height := max(lheight, rheight) + 1

	return diameter, height
}

func diameter(head *TreeNode[int]) int {
	d, _ := diameter_code(head)
	fmt.Println(d)
	return d
}

func main() {
	// bst := NewBinarySearchTree[int]()
	bst := NewBinaryTree[int]()
	nn := 20 + rand.Intn(10)
	for i := 0; i < nn; i++ {
		n := rand.Intn(1000)
		fmt.Printf("%d,", n)
		bst.AddNode(n)
	}

	fmt.Println("\nInorder : ")
	BstInorderTraversal(bst.Head)
	fmt.Println()
	// range_sum(bst)
	// fmt.Printf("minimum_absolute_difference(bst.Head): %v\n", minimum_absolute_difference(bst.Head)) // best
	// fmt.Printf("find_mode(bst.Head): %v\n", find_mode(bst.Head))
	// fmt.Println("Greater Sum : start")
	// greater_sum_tree(bst.Head)
	// BstInorderTraversal(bst.Head)
	// fmt.Println("\nGreater Sum : end")
	// build_balanced_binary_search_tree(bst.Head)

	// lowest_common_ancestor(bst.Head) // best
	// complete_binary_tree() // best
	// pruning_tree(bst.Head) // best

	// path_sum_2(bst.Head)
	// max_diff_node_and_ancestor()
	diameter(bst.Head)
}
