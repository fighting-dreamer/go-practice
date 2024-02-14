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

type Interval struct {
	Start int
	End   int
}

func isOverlapping(a, b Interval) bool {
	return (a.Start <= b.End && a.Start >= b.Start) ||
		(b.Start <= a.End && b.Start >= a.Start)
}

func merge_intervals(intervals []Interval, n int) []Interval {
	slices.SortFunc(intervals, func(a, b Interval) int {
		return cmp.Compare(a.Start, b.Start) // sorting by start time, earlier the interval come, the better.
	})

	curr := intervals[0]
	res := []Interval{}
	for i := 1; i < n; i++ {
		if isOverlapping(curr, intervals[i]) {
			curr.Start = min(curr.Start, intervals[i].Start)
			curr.End = max(curr.End, intervals[i].End)
		} else {
			if curr.End < intervals[i].Start {
				res = append(res, curr)
				curr = intervals[i]
			}
			// else {
			// 	res = append(res, intervals[i])
			// }
		}
	}
	res = append(res, curr)
	return res
}

func main() {
	n := readInt()
	intervals := make([]Interval, n)
	for i := 0; i < n; i++ {
		intervals[i] = Interval{
			Start: readInt(),
			End:   readInt(),
		}
	}

	intervals = merge_intervals(intervals, n)
	fmt.Println(intervals)
}
