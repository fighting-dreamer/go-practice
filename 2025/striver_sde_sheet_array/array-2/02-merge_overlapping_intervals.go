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
	start int
	end   int
}

func isOverlapping(a, b Interval) bool {
	return a.end >= b.start
}

func merge_overlapping_intervals(arr []Interval, n int) []Interval {
	res := []Interval{}
	slices.SortFunc(arr, func(i, j Interval) int {
		return cmp.Compare(i.start, j.start)
	})

	curr := arr[0]
	for i := 1; i < n; i++ {
		if isOverlapping(curr, arr[i]) { // b.start <= a.end
			curr = Interval{
				start: min(arr[i].start, curr.start),
				end:   max(arr[i].end, curr.end),
			}
		} else {
			res = append(res, curr)
			curr = arr[i]
		}
	}
	res = append(res, curr)
	return res
}

func main() {
	n := readInt()
	arr := make([]Interval, n)
	for i := 0; i < n; i++ {
		arr[i] = Interval{
			start: readInt(),
			end:   readInt(),
		}
	}

	res := merge_overlapping_intervals(arr, n)
	for i := 0; i < len(res); i++ {
		fmt.Println(res[i].start, res[i].end)
	}
}
