package main

import (
	"fmt"
	"log"
)

// insert interval in a sorted interval list (sorted by the start time of the interval)

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

func insert_interval_IMPROVED(intervals []Interval, newInterval Interval) []Interval {
	res := make([]Interval, 0)
	i := 0
	n := len(intervals)

	// Phase 1: Add all intervals that come before the new interval
	for i < n && intervals[i].End < newInterval.Start {
		res = append(res, intervals[i])
		i++
	}

	// Phase 2: Merge overlapping intervals
	// We check if the current interval starts before the newInterval ends
	for i < n && intervals[i].Start <= newInterval.End {
		newInterval.Start = min(newInterval.Start, intervals[i].Start)
		newInterval.End = max(newInterval.End, intervals[i].End)
		i++
	}
	// Append the merged result
	res = append(res, newInterval)

	// Phase 3: Add the remaining intervals
	if i < n {
		res = append(res, intervals[i:]...)
	}

	return res
}

func insert_interval(intervals []Interval, n int, newInterval Interval) []Interval {
	// we know the intervals are sorted by the start time
	// else we sort them :
	// slices.SortFunc(intervals, func(i, j Interval) int {
	// 	return cmp.Compare(i.Start, j.Start) // sorting by start time
	// })

	res := []Interval{}
	isOverlapping := func(a, b Interval) bool {
		/*       Case 1             Case 2
		A       ------------ |  ----------
		B  --------          |         ---------
		*/
		// we consider overlapping if the start and end of a and b are same respectively in any order.
		// a : (1,5) and b : (5,6) is overlapping
		// a : (6,9) and b : (3,6) is overlapping

		return (a.Start <= b.End && b.End <= a.End) ||
			(b.Start >= a.Start && b.Start <= a.End)
	}

	for i := 0; i < n; i++ {
		if isOverlapping(newInterval, intervals[i]) {
			newInterval.Start = min(newInterval.Start, intervals[i].Start)
			newInterval.End = max(newInterval.End, intervals[i].End)
		} else {
			if newInterval.End < intervals[i].Start {
				res = append(res, newInterval)
				res = append(res, intervals[i:]...)
				return res // we have inserted the interval.
			} else {
				//Scenrio : intervals[i].End < newInterval.Start
				// basically, the new interval would be coming after few initial intervals.
				res = append(res, intervals[i])
			}
		}
	}
	// we were not able to insert the interval till end and the newInterval is the updated merged interval thta is yet to be inserted. so we insert it.
	res = append(res, newInterval)
	return res
}

func main() {
	n := readInt()
	newInterval := Interval{
		Start: readInt(),
		End:   readInt(),
	}
	intervals := make([]Interval, n)
	for i := 0; i < n; i++ {
		intervals[i] = Interval{
			Start: readInt(),
			End:   readInt(),
		}
	}

	//check : insert_interval_IMPROVED()
	intervals = insert_interval(intervals, n, newInterval)
	fmt.Println(intervals)
}
