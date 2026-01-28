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

func insert_interval(intervals []Interval, n int, newInterval Interval) []Interval {
	res := []Interval{}
	slices.SortFunc(intervals, func(a, b Interval) int {
		return cmp.Compare(a.start, b.start)
	})

	i := 0

	for ; i < n && intervals[i].end < newInterval.start; i++ {
		res = append(res, intervals[i])
	}

	for i < n && intervals[i].start <= newInterval.end {
		newInterval.start = min(newInterval.start, intervals[i].start)
		newInterval.end = max(newInterval.end, intervals[i].end)
		i++
	}
	res = append(res, newInterval)

	res = append(res, intervals[i:]...)

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
	newInterval := Interval{
		start: readInt(),
		end:   readInt(),
	}

	fmt.Printf("insert_interval(intervals, n, newInterval): %v\n", insert_interval(intervals, n, newInterval))
}
