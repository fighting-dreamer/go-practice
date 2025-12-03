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

func print_array(arr []int, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", arr[i])
	}
	fmt.Println()
}

func merge(arr []int, start int, mid int, end int) {
	fmt.Println("merge : ", start, mid, end)
	brr := make([]int, end-start+1)
	curr := 0
	i := start
	j := mid + 1
	for i <= mid && j <= end {
		if arr[i] < arr[j] {
			brr[curr] = arr[i]
			i++
		} else {
			brr[curr] = arr[j]
			j++
		}
		curr++
	}
	for i <= mid {
		brr[curr] = arr[i]
		curr++
		i++
	}
	for j <= end {
		brr[curr] = arr[j]
		j++
		curr++
	}

	for i := start; i <= end; i++ {
		arr[i] = brr[i-start]
	}
}

func merge_sort_helper(arr []int, start int, end int) {
	if start >= end {
		return
	}

	p := (start + end) / 2
	fmt.Println("merge helper : ", start, p, end)
	merge_sort_helper(arr, start, p)
	merge_sort_helper(arr, p+1, end)
	merge(arr, start, p, end)
}

func merge_sort(arr []int, n int) {
	merge_sort_helper(arr, 0, n-1)
}

func main() {
	n := readInt()
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = readInt()
	}
	print_array(arr, n)
	merge_sort(arr, n)
	print_array(arr, n)
}
