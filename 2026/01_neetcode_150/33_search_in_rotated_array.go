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

func find_pivot(arr []int, n int) int {
	start := 0
	end := n - 1
	for start < end {
		mid := start + (end-start)/2

		if arr[mid] > arr[end] {
			start = mid + 1
		} else {
			end = mid
		}
	}

	return start
}
func search_in_rotated_array(arr []int, n int, target int) int {
	pivot := find_pivot(arr, n)
	fmt.Println("pivot : ", pivot, arr[pivot])
	l, ok := slices.BinarySearch(arr[:pivot], target)
	if ok {
		return l
	}

	r, ok := slices.BinarySearch(arr[pivot:], target)
	if ok {
		return pivot + r
	}

	return -1
}

func main() {
	n := readInt()
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = readInt()
	}

	target := readInt()
	fmt.Printf("search_in_rotated_array(arr, n, target): %v\n", search_in_rotated_array(arr, n, target))
}
