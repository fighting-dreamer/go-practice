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
		log.Fatal("Couldn't read number")
	}

	return n
}

func binary_serach(arr []int, n int, k int) int {
	start, end := 0, n-1
	var mid int
	for start < end {
		mid = start + (end-start)/2

		if arr[mid] == k {
			return mid
		} else {
			if arr[mid] < k {
				start = mid + 1
			} else {
				end = mid - 1
			}
		}
	}

	return -1
}

func sort_array(arr []int, n int) {
	sort.Ints(arr)
}

func print_array(arr []int, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", arr[i])
	}
	fmt.Println()
}

func main() {
	n := readInt()
	k := readInt()
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = readInt()
	}
	print_array(arr, n)
	sort_array(arr, n)
	print_array(arr, n)
	fmt.Println("k : ", k, " index : ", binary_serach(arr, n, k))
}
