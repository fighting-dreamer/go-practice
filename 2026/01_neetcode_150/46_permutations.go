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

func permutations_backtracking(arr []int, n int, curr int, res *[][]int) {
	if curr == n {
		*res = append(*res, append([]int{}, arr...))
		return
	}

	for i := curr; i < n; i++ {
		arr[curr], arr[i] = arr[i], arr[curr]
		permutations_backtracking(arr, n, curr+1, res)
		arr[curr], arr[i] = arr[i], arr[curr]
	}
}

func permutations(arr []int, n int) {
	res := [][]int{}
	permutations_backtracking(arr, n, 0, &res)

	for i := 0; i < len(res); i++ {
		fmt.Println(res[i])
	}
}

func main() {
	n := readInt()
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = readInt()
	}

	permutations(arr, n)
}
