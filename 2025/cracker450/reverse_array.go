package main

import (
	"fmt"
	"log"
)

func readInt() int {
	var n int
	_, err := fmt.Scanf("%d", &n)
	if err != nil {
		log.Fatalln("couldn;t read nnumber")
	}
	return n
}

func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func reverse_array(arr []int, n int) {
	for i := 0; i <= n/2; i++ {
		swap(arr, i, n-i-1)
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

	reverse_array(arr, n)
	print_array(arr, n)
}
