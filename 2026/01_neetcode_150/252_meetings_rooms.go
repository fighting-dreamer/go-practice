package main

import (
	"cmp"
	"fmt"
	"log"
	"slices"
)

// link : https://algo.monster/liteproblems/252

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
	start int
	end   int
}

func NewInterval(start, end int) Interval {
	return Interval{
		start: start,
		end:   end,
	}
}

func do_all_meeting_dont_overlap(arr []Interval, n int) bool {
	slices.SortFunc(arr, func(a Interval, b Interval) int {
		return cmp.Compare(a.start, b.start)
	})

	for i := 1; i < n; i++ {
		if arr[i-1].end >= arr[i].start {
			return false
		}
	}

	return true
}

func main() {
	n := readInt()
	arr := make([]Interval, n)
	for i := 0; i < n; i++ {
		arr[i] = NewInterval(readInt(), readInt())
	}

	fmt.Printf("do_all_meeting_dont_overlap(arr, n): %v\n", do_all_meeting_dont_overlap(arr, n))
}
