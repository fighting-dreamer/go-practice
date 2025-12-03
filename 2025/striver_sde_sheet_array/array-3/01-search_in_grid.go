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

func search_in_sorted_grid(mat [][]int, n int, m int, key int) (int, int) {
	mid := struct {
		i int
		j int
	}{
		i: 0,
		j: m - 1,
	} // right most cell of first row, mid of all matrix

	// kinda like a binary search
	for mid.i <= n && mid.j >= 0 {
		// fmt.Printf("%d %d : %d\n", mid.i, mid.j, mat[mid.i][mid.j])
		if mat[mid.i][mid.j] == key {
			return mid.i, mid.j
		} else {
			if mat[mid.i][mid.j] < key {
				mid.i++
			} else {
				mid.j--
			}
		}
	}
	return -1, -1
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

	key := readInt()
	fmt.Println(search_in_sorted_grid(mat, n, m, key))
}
