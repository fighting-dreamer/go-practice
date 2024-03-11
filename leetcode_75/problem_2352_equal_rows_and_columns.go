package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
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

func count_rows_cols_that_are_equal(grid [][]int) int {
	res := 0
	n, m := len(grid), len(grid[0])
	rowCountMap := map[string]int{}
	for i := 0; i < n; i++ {
		s := strings.Builder{}
		for j := 0; j < m; j++ {
			s.WriteString(strconv.Itoa(grid[i][j]))
			s.WriteString("$")
		}
		rowCountMap[s.String()]++
	}

	for i := 0; i < m; i++ {
		s := strings.Builder{}
		for j := 0; j < n; j++ {
			s.WriteString(strconv.Itoa(grid[j][i]))
			s.WriteString("$")
		}
		res += rowCountMap[s.String()]
	}

	return res
}

func main() {
	n, m := readInt(), readInt()
	grid := make([][]int, n)
	for i := 0; i < n; i++ {
		grid[i] = make([]int, m)
		for j := 0; j < m; j++ {
			grid[i][j] = readInt()
		}
	}

	fmt.Println(count_rows_cols_that_are_equal(grid))
}
