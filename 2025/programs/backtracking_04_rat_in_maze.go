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

var dir = [][]int{
	// down
	{
		1, 0,
	},
	// up
	{
		-1, 0,
	},
	// left
	{
		0, -1,
	},
	// right
	{
		0, 1,
	},
}
var cnt = 0

type Pos struct {
	x int
	y int
}

func isValid(pos Pos, n, m int) bool {
	return pos.x >= 0 && pos.x < n && pos.y >= 0 && pos.y < m
}

func rat_in_maze_helper(arr [][]int, n int, m int, pos Pos, vis [][]int, path []int) bool {
	if pos.x == n-1 && pos.y == m-1 {
		fmt.Println(path)
		cnt++
		return true
	}
	var nextPos Pos
	for i := 0; i < 4; i++ {
		nextPos.x = pos.x + dir[i][0]
		nextPos.y = pos.y + dir[i][1]
		if !isValid(nextPos, n, m) {
			continue
		}
		if arr[nextPos.x][nextPos.y] == 1 && vis[nextPos.x][nextPos.y] == 0 {

			vis[nextPos.x][nextPos.y] = 1
			path = append(path, i)
			rat_in_maze_helper(arr, n, m, nextPos, vis, path)
			path = path[:len(path)-1]
			vis[nextPos.x][nextPos.y] = 0
		}
	}

	return true
}

func rat_in_maze(arr [][]int, n, m int) int {
	vis := make([][]int, n)
	for i := 0; i < n; i++ {
		vis[i] = make([]int, m)
	}
	pos := Pos{
		x: 0,
		y: 0,
	}
	vis[0][0] = 1
	rat_in_maze_helper(arr, n, m, pos, vis, []int{})

	return cnt
}

func main() {
	n := readInt()
	m := readInt()
	arr := make([][]int, n)
	for i := 0; i < n; i++ {
		arr[i] = make([]int, m)
		for j := 0; j < m; j++ {
			arr[i][j] = readInt()
		}
	}
	fmt.Println(arr)

	// count ways for exit
	rat_in_maze(arr, n, m)
}
