package main

import (
	"fmt"
	"log"
	"strconv"
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

func valid_box(grid [][]int, i int) bool {
	start := i / 3
	end := i % 3
	numMap := map[int]int{}
	cnt := 0
	for i := start * 3; i < (start+1)*3; i++ {
		for j := end * 3; j < (end+1)*3; j++ {
			if grid[i][j] != 0 {
				cnt++
				numMap[grid[i][j]]++
			}
		}
	}
	if cnt == len(numMap) {
		return true
	}
	return false
}

func valid_row(grid [][]int, i int) bool {
	cnt := 0
	numMap := map[int]int{}
	for j := 0; j < 9; j++ {
		if grid[i][j] != 0 {
			cnt++
			numMap[grid[i][j]]++
		}
	}
	if cnt == len(numMap) {
		return true
	}
	return false
}

func valid_column(grid [][]int, j int) bool {
	cnt := 0
	numMap := map[int]int{}

	for i := 0; i < 9; i++ {
		if grid[i][j] != 0 {
			cnt++
			numMap[grid[i][j]]++
		}
	}
	if cnt == len(numMap) {
		return true
	}
	return false
}

func valid_sudoku(grid [][]int) bool {
	for i := 0; i < 9; i++ {
		if !valid_box(grid, i) {
			return false
		}
	}

	for i := 0; i < 9; i++ {
		if !valid_row(grid, i) {
			return false
		}
	}

	for i := 0; i < 9; i++ {
		if !valid_column(grid, i) {
			return false
		}
	}
	return true
}

func main() {
	n := 9
	grid := make([][]int, n)
	for i := 0; i < n; i++ {
		grid[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			temp := readString()
			if temp == "." {
				grid[i][j] = 0
				continue
			}
			grid[i][j], _ = strconv.Atoi(temp)
		}
	}
	// fmt.Println(grid)

	fmt.Printf("valid_sudoku(grid): %v\n", valid_sudoku(grid))
}
