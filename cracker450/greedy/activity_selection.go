package main

import (
	"fmt"
	"log"
	"slices"

	"github.com/liyue201/gostl/ds/pair"
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

func find_maximum_non_overlapping_activities(activities []pair.Pair, n int) int {
	slices.SortFunc(activities, func(a, b pair.Pair) int {
		if a.Back.(int) == b.Back.(int) {
			return 0
		}
		if a.Back.(int) < b.Back.(int) {
			return -1
		}
		return 1
	})

	cnt := 1
	last := activities[0].Back.(int)
	for i := 1; i < n; i++ {
		if last < activities[i].Front.(int) {
			cnt++
			last = activities[i].Back.(int)
		}
	}

	return cnt
}

func main() {
	n := readInt()
	activities := make([]pair.Pair, n)
	for i := 0; i < n; i++ {
		activities[i] = *pair.MakePair(readInt(), readInt())
	}

	fmt.Println(find_maximum_non_overlapping_activities(activities, n))
}
