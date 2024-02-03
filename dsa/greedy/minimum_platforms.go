package main

import (
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

func minimum_platforms(timings [][]int, n int) int {
	tnums := make([][]int, n*2)
	for i := 0; i < 2*n; i += 2 {
		tnums[i] = []int{timings[i/2][0], 1}
		tnums[i+1] = []int{timings[i/2][1], -1}
	}

	slices.SortFunc(tnums, func(a, b []int) int {
		if a[0] == b[0] {
			if a[1] > b[1] { // if some train have departure as same as some other train arrival, we want to consider arrival first in the sorted list.
				return -1
			} else if a[1] < b[1] {
				return 1
			} else {
				return 0
			}
		} else {
			if a[0] < b[0] {
				return -1
			}
			return 1
		}
	})

	fmt.Println(tnums)
	res := 0
	temp := 0
	for i := 0; i < 2*n; i++ {
		temp += tnums[i][1]
		res = max(res, temp)
	}

	return res
}

func main() {
	n := readInt()
	timings := make([][]int, n)
	for i := 0; i < n; i++ {
		timings[i] = []int{readInt(), readInt()}
	}

	fmt.Println(minimum_platforms(timings, n))
}
