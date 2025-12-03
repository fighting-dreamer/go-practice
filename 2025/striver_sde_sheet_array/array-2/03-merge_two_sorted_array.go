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

func print_array(arr []int, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", arr[i])
	}
	fmt.Println()
}

func merge_sorted_arrays(arr []int, n int, brr []int, m int) ([]int, []int) {
	i := n - 1
	j := 0

	// swapping elements greater than starting elements of brr
	// bigger number move to brr
	// either brr will be completely swapped or partially
	// partially swapped =>
	// 1. some brr element is same or greater than arr element as we move
	//  - and further brr elements will be greater as it is sorted
	// 2. arr was small array but all elemnts in arr will now be smaller ones
	// completly swapped => all brr elements are greater than that of arr
	for i >= 0 && j < m && arr[i] > brr[j] {
		arr[i], brr[j] = brr[j], arr[i]
		i--
		j++
	}
	sort.Ints(arr)
	sort.Ints(brr)

	return arr, brr
}

func main() {
	n, m := readInt(), readInt()
	arr := make([]int, n)
	brr := make([]int, m)

	for i := 0; i < n; i++ {
		arr[i] = readInt()
	}

	for i := 0; i < m; i++ {
		brr[i] = readInt()
	}

	sort.Ints(arr)
	sort.Ints(brr)
	merge_sorted_arrays(arr, n, brr, m)
	print_array(arr, n)
	print_array(brr, m)
}
