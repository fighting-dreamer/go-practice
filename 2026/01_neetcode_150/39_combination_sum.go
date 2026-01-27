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

func bracktracking(arr []int, n int, curr int, target int, sumTillNow int, path []int, res *[][]int) {
	if sumTillNow == target {
		*res = append(*res, path)
		return
	}
	for i := curr; i < n; i++ {
		if sumTillNow+arr[i] <= target {
			bracktracking(arr, n, i, target, sumTillNow+arr[i], append(path, arr[i]), res)
		}
	}
}

func combination_sum(arr []int, n int, target int) [][]int {
	// slices.sort(arr)
	// remove the duplicates just to simplyfy creating unique result arrays
	res := [][]int{}
	bracktracking(arr, n, 0, target, 0, []int{}, &res)
	for i := 0; i < len(res); i++ {
		fmt.Println(res[i])
	}

	return res
}

func main() {
	n := readInt()
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = readInt()
	}
	target := readInt()

	combination_sum(arr, n, target)
}
