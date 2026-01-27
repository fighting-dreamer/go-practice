package main

import (
	"fmt"
	"log"
	"slices"
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

func rotate_image(grid [][]int, n int) [][]int {
	// transpose and reverse
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			grid[i][j], grid[j][i] = grid[j][i], grid[i][j]
		}
	}

	for i := 0; i < n; i++ {
		slices.Reverse(grid[i])
	}

	return grid
}

func main() {
	n := readInt()

	grid := make([][]int, n)
	for i := 0; i < n; i++ {
		grid[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			grid[i][j] = readInt()
		}
	}

	fmt.Printf("rotate_image(grid): %v\n", rotate_image(grid, n))
}
