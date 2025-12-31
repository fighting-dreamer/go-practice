package main

import (
	"fmt"
	"log"
	"sort"
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

func print_all_sums_from_subsets_helper(arr []int, n int, currIndex int, arrTillNow []int, sumTillNow int, res *[]int) {
	if currIndex == n {
		fmt.Println(arrTillNow, ": ", sumTillNow)
		*res = append(*res, sumTillNow)
	} else {
		print_all_sums_from_subsets_helper(arr, n, currIndex+1, append(arrTillNow, arr[currIndex]), sumTillNow+arr[currIndex], res)
		print_all_sums_from_subsets_helper(arr, n, currIndex+1, arrTillNow, sumTillNow, res)
	}
}

func print_all_sums_from_subsets(arr []int, n int) []int {
	res := []int{}
	print_all_sums_from_subsets_helper(arr, n, 0, []int{}, 0, &res)
	sort.Ints(res)
	return res
}

func main() {
	n := readInt()
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = readInt()
	}

	fmt.Printf("print_all_sums_from_subsets(arr, n): %v\n", print_all_sums_from_subsets(arr, n))
}
