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

func two_sum(arr []int, n int, target int) (int, int) {
	l := 0
	r := n - 1
	for l < r {
		sum := arr[l] + arr[r]
		if sum == target {
			return arr[l], arr[r]
		} else {
			if sum < target {
				// increment by small change
				l++
			} else {
				// decrement by small change
				r--
			}
		}
	}

	return -1, -1
}

func main() {
	n := readInt()
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = readInt()
	}

	target := readInt()
	// array is sorted
	fmt.Println(two_sum(arr, n, target))
}
