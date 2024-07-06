package main

import (
	"cmp"
	"fmt"
	"log"
	"slices"
)

type Interval struct {
	start int
	end   int
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

func minimum_arrows_to_burst_ballons(nums [][]int, n int) int {
	intervals := make([]Interval, n)
	for i := 0; i < n; i++ {
		intervals[i] = Interval{nums[i][0], nums[i][1]}
	}
	slices.SortFunc(intervals, func(a, b Interval) int {
		return cmp.Compare[int](a.end, b.end)
	})
	fmt.Println(intervals)
	res := 1
	currIntervalEnd := intervals[0].end
	for i := 1; i < n; i++ {
		if currIntervalEnd < intervals[i].start {
			res++
			currIntervalEnd = intervals[i].end
		}
	}

	return res
}

func main() {
	n := readInt()
	nums := make([][]int, n)
	for i := 0; i < n; i++ {
		nums[i] = []int{readInt(), readInt()}
	}

	fmt.Println(minimum_arrows_to_burst_ballons(nums, n))
}
