package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/liyue201/gostl/ds/pair"
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

func print_matrix(mat [][]int, n int, m int) {
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Printf("%d ", mat[i][j])
		}
		fmt.Println()
	}
	fmt.Println()
}

type LoopRange struct {
	start *pair.Pair
	end   *pair.Pair
	delta *pair.Pair
}

func (l LoopRange) String() string {
	var sb strings.Builder

	fmt.Fprintf(&sb, "start : %d, %d \n", l.start.Front.(int), l.start.Back.(int))
	fmt.Fprintf(&sb, "end : %d, %d \n", l.end.Front.(int), l.end.Back.(int))
	fmt.Fprintf(&sb, "delta : %d, %d \n", l.delta.Front.(int), l.delta.Back.(int))

	return sb.String()
}

func swapLoopRange(mat [][]int, rangeA LoopRange, rangeB LoopRange, n int) {
	for i := 0; i < n; i++ {
		a_row := rangeA.start.Front.(int) + rangeA.delta.Front.(int)*i
		a_col := rangeA.start.Back.(int) + rangeA.delta.Back.(int)*i

		b_row := rangeB.start.Front.(int) + rangeB.delta.Front.(int)*i
		b_col := rangeB.start.Back.(int) + rangeB.delta.Back.(int)*i
		fmt.Println("swap : ", a_row, a_col, b_row, b_col)
		mat[a_row][a_col], mat[b_row][b_col] = mat[b_row][b_col], mat[a_row][a_col]
	}
}

func rotate_matrix_helper(mat [][]int, n int, start_row int, start_col int, curr_size int) {
	if curr_size <= 1 {
		return
	}

	top_range := LoopRange{
		start: pair.MakePair(start_row, start_col),
		end:   pair.MakePair(start_row, start_col+curr_size-1),
		delta: pair.MakePair(0, 1),
	}
	right_range := LoopRange{
		start: pair.MakePair(start_row, start_col+curr_size-1),
		end:   pair.MakePair(start_row+curr_size-1, start_col+curr_size-1),
		delta: pair.MakePair(1, 0),
	}
	bottom_range := LoopRange{
		start: pair.MakePair(start_row+curr_size-1, start_col+curr_size-1),
		end:   pair.MakePair(start_row+curr_size-1, start_col),
		delta: pair.MakePair(0, -1),
	}
	left_range := LoopRange{
		start: pair.MakePair(start_row+curr_size-1, start_col),
		end:   pair.MakePair(start_row, start_col),
		delta: pair.MakePair(-1, 0),
	}

	fmt.Println("n : ", curr_size)
	fmt.Println("-----------------------------")

	fmt.Printf("top_range: %v\n", top_range)
	fmt.Printf("right_range: %v\n", right_range)
	fmt.Printf("bottom_range: %v\n", bottom_range)
	fmt.Printf("left_range: %v\n", left_range)

	swapLoopRange(mat, top_range, right_range, curr_size-1)
	swapLoopRange(mat, top_range, bottom_range, curr_size-1)
	swapLoopRange(mat, top_range, left_range, curr_size-1)

	print_matrix(mat, n, n)
	fmt.Println("-----------------------------")

	rotate_matrix_helper(mat, n, start_row+1, start_col+1, curr_size-2)
}

// square matrix
func rotate_matrix(mat [][]int, n int) {
	rotate_matrix_helper(mat, n, 0, 0, n)
}

func main() {
	n := readInt()
	mat := make([][]int, n) // square matrix

	for i := 0; i < n; i++ {
		mat[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			mat[i][j] = readInt()
		}
	}

	rotate_matrix(mat, n)
}
