package main

import (
	"fmt"
	"log"
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

func print_matrix(mat [][]int, n, m int) {
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Printf("%d ", mat[i][j])
		}
		fmt.Println()
	}
}

func set_matrix_zero(mat [][]int, n, m int) {
	// for any i, j, if we encounter zero, we can mark mat[i][0] and mat[0][j] as zero.
	// special case is for first row or first column : it will mark mat[0][0] as zero,
	// but it might be only row or only column was zero, so
	// to solve for that : we keep a single variable for say just column
	// pass 1.1 : go through matrix, for first column, check if you saw any zero, set is_first_col_zero as true
	// pass 1.2 : go through rest of matrix, any i, j is zero => mat[i][0] and mat[0][j] is zero.
	// pass 2 : for all rows or colum with zero in first frow or col are marked zero in matrix,
	// then we check if is_first_col_zero, and accordingly mark first col zero if its true.

	// if mat[0][0] was actually zero => is_first_col_zero => true and mat[0][0] is zero ==> row zero is also zero
	// if first row is having zero(except mat[0][0]) => column will have zero but is_first_col_zero will be false

	// pass 1 :
	is_first_col_zero := false
	for i := 0; i < n; i++ {
		if mat[i][0] == 0 {
			is_first_col_zero = true
		}
	}

	for i := 0; i < n; i++ {
		// rest of columns
		for j := 0; j < m; j++ {
			if mat[i][j] == 0 {
				mat[i][0] = 0
				mat[0][j] = 0
			}
		}
	}
	// pass 2 :
	for i := 1; i < n; i++ {
		if mat[i][0] == 0 {
			for j := 1; j < m; j++ {
				mat[i][j] = 0
			}
		}
	}

	for j := 1; j < m; j++ {
		if mat[0][j] == 0 {
			for i := 1; i < n; i++ {
				mat[i][j] = 0
			}
		}
	}
	// is first row suppose to be zero => if mat[0][0] is zero
	if mat[0][0] == 0 {
		for j := 0; j < m; j++ {
			mat[0][j] = 0
		}
	}
	// is first colum zero
	if is_first_col_zero {
		for i := 0; i < n; i++ {
			mat[i][0] = 0
		}
	}
}

func main() {
	n := readInt()
	m := readInt()

	mat := make([][]int, n)
	for i := 0; i < n; i++ {
		mat[i] = make([]int, m)
		for j := 0; j < m; j++ {
			mat[i][j] = readInt()
		}
	}

	print_matrix(mat, n, m)
	set_matrix_zero(mat, n, m)
	print_matrix(mat, n, m)
}
