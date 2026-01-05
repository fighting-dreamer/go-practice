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

func swap(arr []int, i, j int) {
	if i != j {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func find_farthest_max_number_index(arr []int, key int, itNext int, n int) int {
	res := itNext - 1
	for i := n - 1; i >= itNext; i-- {
		if arr[res] < arr[i] {
			res = i
		}
	}
	if res == itNext-1 {
		return -1
	}
	return res
}

func largest_number_with_at_most_k_swaps_helper(arr []int, it int, k int, n int) {
	if k == 0 || it == n {
		fmt.Println(arr) // no more swaps can be done
		return
	}

	si := find_farthest_max_number_index(arr, arr[it], it+1, n)
	// we didn't find element with value greater than arr[it]
	// move to next index
	if si == -1 {
		largest_number_with_at_most_k_swaps_helper(arr, it+1, k, n)
		return
	} else {
		swap(arr, it, si)
		fmt.Println(k, " :", arr)
		largest_number_with_at_most_k_swaps_helper(arr, it+1, k-1, n)
		swap(arr, it, si)
	}

}

func largest_number_with_at_most_k_swaps(arr []int, k int, n int) {
	largest_number_with_at_most_k_swaps_helper(arr, 0, k, n)
}

// does not work : 7 5 9 9, k = 2
// res : 9 9 7 5, op:(0,2),(1,3)
func main() {
	n := readInt()
	k := readInt()
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = readInt()
	}

	largest_number_with_at_most_k_swaps(arr, k, n)
}
