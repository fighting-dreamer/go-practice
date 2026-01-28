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

func search_in_sorted_matrix(mat [][]int, m int, n int, target int) bool {
	i, j := 0, n-1
	for i >= 0 && i < m && j >= 0 && j < n {
		if mat[i][j] == target {
			fmt.Println(i, j, mat[i][j])
			return true
		} else {
			if mat[i][j] < target {
				i++
			} else {
				j--
			}
		}
	}

	return false
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

	target := readInt()

	fmt.Printf("search_in_sorted_matrix(mat, m, n, target): %v\n", search_in_sorted_matrix(mat, m, n, target))
}
