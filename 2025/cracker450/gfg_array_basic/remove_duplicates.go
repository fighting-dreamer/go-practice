package main

import (
	"fmt"
	"log"
	"sort"
)

func sort_array(arr []int, n int) {
	sort.Ints(arr)
}

func remove_duplicates(arr []int, n int) []int {
	brr := []int{arr[0]}
	for i := 1; i < n; i++ {
		if arr[i-1] != arr[i] {
			brr = append(brr, arr[i])
		}
	}
	fmt.Println("brr : ", brr)
	fmt.Println(len(brr))
	return brr
}

func readInt() int {
	var n int
	_, err := fmt.Scanf("%d", &n)
	if err != nil {
		log.Fatalln("coult not read int")
	}
	return n
}

func print_array(arr []int, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", arr[i])
	}
	fmt.Println()
}

func main() {
	n := readInt()
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = readInt()
	}

	sort_array(arr, n)
	arr = remove_duplicates(arr, n)
	print_array(arr, len(arr))
}
