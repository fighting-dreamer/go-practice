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

func find_duplicate(arr []int, n int) int {
	xor := 0
	for i := 1; i <= n; i++ {
		xor = xor ^ i
	}
	for i := 0; i < n+1; i++ {
		xor = xor ^ arr[i]
	}

	return xor
}

func main() {
	n := readInt()
	arr := make([]int, n+1)
	for i := 0; i < n+1; i++ {
		arr[i] = readInt()
	}

	fmt.Printf("find_duplicate(arr, n): %v\n", find_duplicate(arr, n))
}
