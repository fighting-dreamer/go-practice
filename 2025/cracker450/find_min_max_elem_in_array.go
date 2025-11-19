package main

import (
	"fmt"
	"log"
)

func readInt() int {
	var n int
	_, err := fmt.Scanf("%d", &n)
	if err != nil {
		log.Fatalf("couldnt read number")
	}
	return n
}

func find_min_max_element(arr []int, n int) (int, int) {
	minV, maxV := arr[0], arr[0]
	for i := 0; i < n; i++ {
		// minV = (minV > arr[i] ? arr[i]:minV);
		if minV > arr[i] {
			minV = arr[i]
		}

		// maxV = (maxV < arr[i] ? arr[i]:maxV)
		if maxV < arr[i] {
			maxV = arr[i]
		}
	}
	return minV, maxV
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
	minVal, maxVal := find_min_max_element(arr, n)
	fmt.Printf("min : %d, max: %d", minVal, maxVal)
}
