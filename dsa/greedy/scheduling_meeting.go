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

func scheduling_meeting(nums [][]int, n int) int {
	// max meeting getting scheduled.
	fmt.Println(nums)
	slices.SortFunc(nums, func(a, b []int) int {
		if a[1] == b[1] {
			return 0
		}
		if a[1] < b[1] {
			return -1
		}
		return 1
	})
	fmt.Println(nums)
	res := 1
	last := nums[0][1]
	for i := 1; i < n; i++ {
		if last < nums[i][0] {
			res++
			last = nums[i][1]
		}
	}

	return res
}

/*
input :
6
1 2
0 6
3 4
8 9
5 7
8 9

output : 4
*/

func main() {
	n := readInt()
	nums := make([][]int, n)
	for i := 0; i < n; i++ {
		nums[i] = []int{readInt(), readInt()}
	}

	fmt.Println(scheduling_meeting(nums, n))
}
