package main

import (
	"fmt"
	"log"
)

type Range struct {
	start_row int
	start_col int
	row_incr  int
	col_incr  int
	size      int
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

func swap(nums [][]int, r1, c1, r2, c2 int) {
	nums[r1][c1], nums[r2][c2] = nums[r2][c2], nums[r1][c1]
}

func swapRange(nums [][]int, range1, range2 Range) {
	row1, col1, row2, col2 := range1.start_row, range1.start_col, range2.start_row, range2.start_col
	for i := 0; i < range1.size; i++ {
		swap(nums, row1, col1, row2, col2)
		row1 += range1.row_incr
		col1 += range1.col_incr
		row2 += range2.row_incr
		col2 += range2.col_incr
	}

	return
}

func rotate_image_by_90helper(nums [][]int, n int, startRow, startCol, endRow, endCol int) {
	if n < 2 {
		return
	}
	// have to be inplace

	// strategy : example of 4X4 matrix :
	// first 3 cells of first row swap with first 3 cells of right most column
	// the first 3 cells of right most swap with 3 right most cells of last row
	// the right most 3 cells of last row swap with last 3 cells of first column.
	// call same kinda operation on a 2X2 matrix inside of this outer most row/columns.
	ranges := []Range{
		{startRow, startCol, 0, 1, n - 1},
		{startRow, endCol, 1, 0, n - 1},
		{endRow, endCol, 0, -1, n - 1},
		{endRow, startCol, -1, 0, n - 1},
	}
	for i := 0; i < n-1; i++ {
		swapRange(nums, ranges[0], ranges[i+1]) // STAR imporant to only swap with first range.
	}

	rotate_image_by_90helper(nums, n-2, startRow+1, startCol+1, endRow-1, endCol-1)
}

func rotate_image_by_90(nums [][]int, n int) {
	rotate_image_by_90helper(nums, n, 0, 0, n-1, n-1)
	return
}

func main() {
	n := readInt()
	nums := make([][]int, n)
	for i := 0; i < n; i++ {
		nums[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			nums[i][j] = readInt()
		}
	}

	rotate_image_by_90(nums, n)
	fmt.Println(nums)
}
