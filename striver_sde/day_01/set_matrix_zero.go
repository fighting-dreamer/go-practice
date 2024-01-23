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

func emptyMatrix[T any](size int) [][]T {
	mat := make([][]T, size)
	for i := 0; i < size; i++ {
		mat[i] = make([]T, size)
	}

	return mat
}

func emptyMatrixMN[T any](m, n int) [][]T {
	mat := make([][]T, m)
	for i := 0; i < m; i++ {
		mat[i] = make([]T, n)
	}

	return mat
}

func set_matrix_zero(mat [][]int, m, n int) [][]int {
	firstRow, firstCol := 1, 1
	for i := 0; i < n; i++ {
		if mat[0][i] == 0 {
			firstRow = 0
			break
		}
	}

	for i := 0; i < m; i++ {
		if mat[i][0] == 0 {
			firstCol = 0
			break
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if mat[i][j] == 0 {
				mat[0][j] = 0
				mat[i][0] = 0
			}
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if mat[i][0] == 0 {
				mat[i][j] = 0
			}
			if mat[0][j] == 0 {
				mat[i][j] = 0
			}
		}
	}

	if firstRow == 0 {
		for i := 0; i < n; i++ {
			mat[0][i] = 0
		}
	}

	if firstCol == 0 {
		for i := 0; i < m; i++ {
			mat[i][0] = 0
		}
	}

	return mat
}

func main() {
	m, n := readInt(), readInt()
	mat := emptyMatrixMN[int](m, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			mat[i][j] = readInt()
		}
	}

	mat = set_matrix_zero(mat, m, n)

	fmt.Println(mat)
}
