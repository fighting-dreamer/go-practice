package main

import (
	"fmt"
	"log"
)

func readInt() int {
	var n int
	_, err := fmt.Scanf("%d", &n)
	if err != nil {
		log.Fatalln("error")
	}
	return n
}

func print_array(arr []int, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", arr[i])
	}
	fmt.Println()
}

func reverse_array(arr []int, n int) {
	i := 0
	j := n - 1
	for i < j {
		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}
}

func rotate_array(arr []int, n int, k int) {
	reverse_array(arr[:k], k)
	reverse_array(arr[k:], n-k)
	reverse_array(arr, n)
}

func main() {
	n := readInt()
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = readInt()
	}
	k := readInt()
	print_array(arr, n)
	rotate_array(arr, n, k)
	print_array(arr, n)
}
