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

type Coord struct {
	x int
	y int
}

func isValidCoord(x, y, m, n int) bool {
	return x >= 0 && x < m && y >= 0 && y < n
}

func rotten_oranges(grid [][]int, m int, n int) int {
	dirs := []Coord{
		{x: 1, y: 0},
		{x: 0, y: -1},
		{x: -1, y: 0},
		{x: 0, y: 1},
	}

	vis := make([][]int, m)
	for i := 0; i < m; i++ {
		vis[i] = make([]int, n)
	}

	q := []Coord{}
	freshOranges := 0

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 2 {
				vis[i][j] = 1
				q = append(q, Coord{
					x: i,
					y: j,
				})
			} else if grid[i][j] == 1 {
				freshOranges++
			}
		}
	}

	if freshOranges == 0 {
		return 0
	}

	minutes := 0

	for len(q) > 0 {
		l := len(q)
		anyRottenThisturn := false
		for t := 0; t < l; t++ {
			top := q[0]
			q = q[1:]

			for i := 0; i < 4; i++ {
				newX := top.x + dirs[i].x
				newY := top.y + dirs[i].y

				if isValidCoord(newX, newY, m, n) && vis[newX][newY] == 0 && grid[newX][newY] == 1 {
					grid[newX][newY] = 2 // make it rotten

					q = append(q, Coord{
						x: newX,
						y: newY,
					})
					freshOranges--
					anyRottenThisturn = true
					vis[newX][newY] = 1 // mark this newX, newY as visited
				}
			}
		}
		if anyRottenThisturn {
			minutes++
		}
	}

	if freshOranges > 0 {
		return -1
	}

	return minutes
}

func main() {
	m := readInt()
	n := readInt()

	grid := make([][]int, m)
	for i := 0; i < m; i++ {
		grid[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			grid[i][j] = readInt()
		}
	}

	rotten_oranges(grid, m, n)
}
