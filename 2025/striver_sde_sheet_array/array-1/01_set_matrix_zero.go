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
	is_first_col_zero := false
	is_first_row_zero := false

	for i := 0; i < n; i++ {
		if mat[i][0] == 0 {
			is_first_col_zero = true
			break
		}
	}

	for j := 0; j < m; j++ {
		if mat[0][j] == 0 {
			is_first_row_zero = true
			break
		}
	}

	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			if mat[i][j] == 0 {
				mat[i][0] = 0
				mat[0][j] = 0
			}
		}
	}

	// pass 2

	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			if mat[i][0] == 0 || mat[0][j] == 0 {
				mat[i][j] = 0
			}
		}
	}
	if is_first_col_zero {
		for i := 0; i < n; i++ {
			mat[i][0] = 0
		}
	}

	if is_first_row_zero {
		for j := 0; j < m; j++ {
			mat[0][j] = 0
		}
	}
}

func set_matrix_zero_2(mat [][]int, n int, m int) {
	is_first_col_zero := false

	// pass 1 : marking
	for row_i := 0; row_i < n; row_i++ {
		if mat[row_i][0] == 0 {
			is_first_col_zero = true
			// row_i is zero then that row will be zero later.
		}
		for col_j := 1; col_j < m; col_j++ {
			if mat[row_i][col_j] == 0 {
				mat[row_i][0] = 0
				mat[0][col_j] = 0
			}
		}
	}

	// pass 2 : set matrix zero
	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			if mat[i][0] == 0 || mat[0][j] == 0 {
				mat[i][j] = 0
			}
		}
	}

	if mat[0][0] == 0 {
		for j := 0; j < m; j++ {
			mat[0][j] = 0 // marking all cells of row-0 as zero
		}
	}

	if is_first_col_zero {
		for i := 0; i < n; i++ {
			mat[i][0] = 0 // marking all cells of col-0 as zero
		}
	}
}

func copy_matrix(mat [][]int, n int, m int) [][]int {
	cm := make([][]int, n)
	for i := 0; i < n; i++ {
		cm[i] = make([]int, m)
		for j := 0; j < m; j++ {
			cm[i][j] = mat[i][j]
		}
	}

	return cm
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
	fmt.Println("----------------------------------------")
	fmt.Println("----------------------------------------")

	print_matrix(mat, n, m)
	fmt.Println()
	mat1 := copy_matrix(mat, n, m)
	set_matrix_zero(mat1, n, m)
	print_matrix(mat1, n, m)
	fmt.Println("----------------------------------------")
	mat2 := copy_matrix(mat, n, m)
	set_matrix_zero_2(mat2, n, m)
	print_matrix(mat2, n, m)
}
