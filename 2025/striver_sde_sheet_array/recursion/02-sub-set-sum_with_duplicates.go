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

func print_all_unique_subsets_if_Array_have_duplicates_helper(arr []int, n int, currIndex int, sumTillNow int, arrTillNow []int, res *[][]int) {
	if currIndex == n {
		return 
	}
}

func print_all_unique_subsets_if_Array_have_duplicates(arr []int, n int) [][]int {
	res := [][]int{}
	print_all_unique_subsets_if_Array_have_duplicates_helper(arr, n, 0, 0, &res)
	return res
}

func main() {
	n := readInt()
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = readInt()
	}

	print_all_unique_subsets_if_Array_have_duplicates(arr, n)
}
