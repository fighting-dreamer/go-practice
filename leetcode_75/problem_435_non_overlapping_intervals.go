package main

import (
	"cmp"
	"fmt"
	"log"
	"slices"
)

type Interval struct {
	Start int
	End   int
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

func non_overlapping_intervals(intervals []Interval, n int) int {
	// res := []Interval{}
	slices.SortFunc(intervals, func(a, b Interval) int {
		if a.End == b.End {
			return cmp.Compare(a.Start, b.Start)
		}
		if a.End < b.End {
			return -1
		}
		return 1
	})
	// result is (total intervals in input) - (maximum non-overlapping intervals)
	// we are sorting by End as for an interval, if an interval is coming before and overlaps, the one with ening one will be more prone to overlapping with upcoming intervals.

	last := intervals[0].End
	cnt := 1
	for i := 1; i < n; i++ {
		if last <= intervals[i].Start {
			cnt++
			last = intervals[i].End
		}
	}
	fmt.Println(cnt)
	return n - cnt
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
	fmt.Println(non_overlapping_intervals(intervals, n))
}
