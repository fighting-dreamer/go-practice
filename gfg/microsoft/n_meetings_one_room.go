package main

import (
	"cmp"
	"fmt"
	"log"
	"slices"
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

func n_meetings_one_room(nums [][]int, n int) int {
	slices.SortFunc(nums, func(a, b []int) int {
		return cmp.Compare(a[1], b[1])
	})

	current_end := nums[0][1]
	cnt := 1

	for i := 1; i < n; i++ {
		if current_end < nums[i][0] {
			cnt++
			current_end = nums[i][1]
		}
	}

	return cnt
}

func main() {
	n := readInt()
	nums := make([][]int, n)
	for i := 0; i < n; i++ {
		nums[i] = []int{readInt(), readInt()}
	}

	fmt.Println(n_meetings_one_room(nums, n))
}
