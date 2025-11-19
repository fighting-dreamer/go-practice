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
		log.Fatal("couldnt read number")
	}

	return n
}

func print_array(arr []int, n int ) {
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", arr[i])
	}
	fmt.Println()
}

func sort_array_ascending(arr []int , n int ) {
	sort.Ints(arr)
}

func sort_array_descending(arr []int, n int) {
	sort.Sort(sort.Reverse(sort.IntSlice(arr)))
}

func main() {
	n := readInt()
	arr := make([]int, n)

	for i := 0; i < n; i++ {
		arr[i] = readInt()
	}

	print_array(arr, n)
	sort_array_descending(arr, n)
	print_array(arr, n)
}