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

/*
	1 2 3 4
	1 2 4 3
	1 3 2 4
	1 3 4 2
	1 4 2 3
	1 4 3 2
	2 1 3 4
	2 1 4 3
	...

	say you are at 1 3 4 2
	next one is 1 4 2 3

	moving from right to left
	you see 2 < 3, we can swap them to as there can be an element greater, you get 1 4 3 2

	1 4 3 2 => you see `1` < `4`, now this 1 can be swapped, but we need element just greater than 1
	that element will be in first element from right toward left => 2
	2 {1 4 3} => sort them => 2 1 3 4

	any random sequence with duplicates : 2 1 5 4 3 0 0
	`1` < `5` coming from right
	2 3 {5 4 1 0 0} => sort them => 2 3 0 0 1 4 5

	you can actually just reverse them too.

	2 4 3 1

	`2` < `4`
	next greater than 2 is 3
	3 {4 2 1} => reverse them => 3 1 2 4

	4 3 2 1

*/

func next_greater_element(arr []int, curr int, n int) int {
	for i := n - 1; i > curr; i-- {
		if arr[i] > arr[curr] {
			return i
		}
	}
	return -1
}

func swap(arr []int, i int, j int) {
	if i == -1 || j == -1 {
		return
	}
	arr[i], arr[j] = arr[j], arr[i]
}

func reverse(arr []int, start, end int) {
	for start < end {
		swap(arr, start, end)
		start++
		end--
	}
}

func next_permutation(arr []int, n int) {
	for i := n - 2; i >= 0; i-- {
		if arr[i] < arr[i+1] {
			swap(arr, i, next_greater_element(arr, i, n))
			reverse(arr, i+1, n-1)
		}
	}
}

func main() {
	n := readInt()
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = readInt()
	}
	fmt.Println(arr)
	next_permutation(arr, n)
	fmt.Println(arr)
}
