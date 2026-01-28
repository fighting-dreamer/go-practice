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

func print_matrix(mat [][]int, m int, n int) {
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			fmt.Printf("%2d ", mat[i][j])
		}
		fmt.Println()
	}
	fmt.Println("--------------------------")
}

func set_matrix_zero(mat [][]int, m int, n int) [][]int {
	first_row_flag := false
	first_col_flag := false
	if mat[0][0] == 0 {
		first_col_flag = true
		first_row_flag = true
	}

	for i := 1; i < m && first_col_flag == false; i++ {
		if mat[i][0] == 0 {
			first_col_flag = true
		}
	}
	for i := 1; i < n && first_row_flag == false; i++ {
		if mat[0][i] == 0 {
			first_row_flag = true
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if mat[i][j] == 0 {
				mat[i][0] = 0
				mat[0][j] = 0
			}
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if mat[i][0] == 0 || mat[0][j] == 0 {
				mat[i][j] = 0
			}
		}
	}
	if first_col_flag {
		for i := 0; i < m; i++ {
			mat[i][0] = 0
		}
	}
	if first_row_flag {
		for i := 0; i < n; i++ {
			mat[0][i] = 0
		}
	}

	return mat
}

func main() {
	m, n := readInt(), readInt()
	mat := make([][]int, m)
	for i := 0; i < m; i++ {
		mat[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			mat[i][j] = readInt()
		}
	}
	print_matrix(mat, m, n)
	mat = set_matrix_zero(mat, m, n)
	print_matrix(mat, m, n)
}
