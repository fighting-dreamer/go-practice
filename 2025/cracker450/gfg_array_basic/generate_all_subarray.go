package main

import (
	"fmt"
	"log"
)

func readInt() int {
	var n int
	_, err := fmt.Scanf("%d", &n)
	if err != nil {
		log.Fatal("couln't read int")
	}
	return n
}

func generate_all_subarray(arr []int, n int) {
	// sub array : contigous segments of diff lengths
	for l := 1; l <= n; l++ {
		fmt.Println("---------------", l, "--------------------")
		for i := 0; i <= n-l; i++ {
			fmt.Println("Starting from index : ", i, " for l : ", l)
			for j := i; j < i+l; j++ {
				fmt.Printf("%d ", arr[j])
			}
			fmt.Println()
		}
		fmt.Println("-----------------------------------")
	}
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

	print_array(arr, n)
	generate_all_subarray(arr, n)

}
