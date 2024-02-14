package main

import (
	"cmp"
	"fmt"
	"log"
	"slices"
)

type _Schedule struct {
	Arrival   int
	Departure int
}

type _ScheduleCheckpoint struct {
	val  int
	sign int
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

func find_minimum_number_of_platforms(ss []_Schedule, n int) int {
	cs := make([]_ScheduleCheckpoint, 2*n)
	for i := 0; i < n; i++ {
		cs[2*i] = _ScheduleCheckpoint{
			val:  ss[i].Arrival,
			sign: 1,
		}
		cs[2*i+1] = _ScheduleCheckpoint{
			val:  ss[i].Departure,
			sign: -1,
		}
	}

	slices.SortFunc(cs, func(a, b _ScheduleCheckpoint) int {
		if a.val == b.val {
			return cmp.Compare(b.sign, a.sign)
		}
		if a.val < b.val {
			return -1
		}
		return 1
	})

	fmt.Println(cs)

	curr, cnt := 0, 0
	for i := 0; i < 2*n; i++ {
		curr += cs[i].sign
		cnt = max(cnt, curr)
	}

	return cnt
}

func main() {
	n := readInt()
	schedules := make([]_Schedule, n)
	for i := 0; i < n; i++ {
		schedules[i] = _Schedule{
			Arrival:   readInt(),
			Departure: readInt(),
		}
	}

	fmt.Println(find_minimum_number_of_platforms(schedules, n))
}
