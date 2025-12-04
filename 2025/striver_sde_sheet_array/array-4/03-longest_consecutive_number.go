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

func consecutive_number_longest_subarray(arr []int, n int) int {
	sort.Ints(arr)
	fmt.Println(arr)
	res := 0
	var cnt int
	for i := 1; i < n; i++ {
		fmt.Println(i, cnt)
		if arr[i] == arr[i-1]+1 {
			cnt++
		} else {
			cnt = 0
		}
		res = max(res, cnt+1) // including the first element of consecutive numbers
	}
	return res
}

func main() {
	n := readInt()
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = readInt()
	}

	fmt.Printf("consecutive_number_longest_subarray(arr, n): %v\n", consecutive_number_longest_subarray(arr, n))
}
