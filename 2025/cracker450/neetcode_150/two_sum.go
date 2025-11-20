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
		log.Fatal("Couldn;t read number")
	}
	return n
}

func print_array(arr []int, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", arr[i])
	}
	fmt.Println()
}

func sort_array(arr []int, n int) {
	sort.Ints(arr)
}

func two_sum(arr []int, n int, k int) (int, int) {
	i, j := 0, n - 1
	for i < j {
		curr := arr[i] + arr[j]
		if curr == k {
			return i, j
		}else {
			if curr > k {
				j--
			}else {
				i++
			}
		}
	}

	return -1, -1
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
	i, j := two_sum(arr, n, k)
	fmt.Println(i, j)
	fmt.Println(arr[i], arr[j])
}