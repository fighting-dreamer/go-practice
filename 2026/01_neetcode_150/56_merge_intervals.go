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

func merge_intervals(intervals []Interval, n int) []Interval {
	slices.SortFunc(intervals, func(a, b Interval) int {
		return cmp.Compare(a.start, b.start)
	})
	res := []Interval{}

	curr := intervals[0]
	for i := 1; i < n; i++ {
		if intervals[i].start > curr.end {
			res = append(res, curr)
			curr = intervals[i]
		} else {
			curr.start = min(curr.start, intervals[i].start)
			curr.end = max(curr.end, intervals[i].end)
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
			start: readInt(),
			end:   readInt(),
		}
	}
	fmt.Println(intervals)
	fmt.Printf("merge_intervals(intervals, n): %v\n", merge_intervals(intervals, n))
}
