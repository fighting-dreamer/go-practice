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

func search_matrix(mat [][]int, n, m, t int) (int, int) {
	i, j := 0, m-1

	for i < n && j < m && i >= 0 && j >= 0 {
		// fmt.Println(mat[i][j], i, j)
		if mat[i][j] == t {
			return i, j
		}
		if mat[i][j] > t {
			j--
		} else {
			i++
		}
	}

	return -1, -1
}

func main() {
	n, m, t := readInt(), readInt(), readInt()
	mat := make([][]int, n)
	for i := 0; i < n; i++ {
		mat[i] = make([]int, m)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			mat[i][j] = readInt()
		}
	}

	fmt.Println(search_matrix(mat, n, m, t))

}
