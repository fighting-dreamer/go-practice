package main

import (
	"fmt"
	"log"
)

func print_array(arr []int, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", arr[i])
	}
	fmt.Println()
}

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

func move_zeros_to_last(arr []int, n int) {
	i, j := 0, n-1
	for i < j {
		// move till you encounter zero
		for ; i < j; i++ {
			if arr[i] == 0 {
				break
			}
		}
		// move till you encounter non-zero
		for ; j > i; j-- {
			if arr[j] != 0 {
				break
			}
		}

		// make swap
		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}
}

func main() {
	n := readInt()
	arr := make([]int, n)

	for i := 0; i < n; i++ {
		arr[i] = readInt()
	}
	print_array(arr, n)
	move_zeros_to_last(arr, n)
	print_array(arr, n)
}
