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

func unique_paths(n int, m int) int {
	grid := make([][]int, n)
	for i := 0; i < n; i++ {
		grid[i] = make([]int, m)
	}

	for i := 0; i < n; i++ {
		grid[i][0] = 1
	}

	for i := 0; i < m; i++ {
		grid[0][i] = 1
	}

	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			grid[i][j] = grid[i-1][j] + grid[i][j-1]
		}
	}

	fmt.Println(grid)

	return grid[n-1][m-1]
}

func main() {
	n, m := readInt(), readInt()

	fmt.Printf("unique_paths(n, m): %v\n", unique_paths(n, m))
}
