package main

import (
	"cmp"
	"fmt"
	"log"
	"slices"
)

// https://www.geeksforgeeks.org/maximum-trains-stoppage-can-provided/

type Schedule struct {
	arrival   int
	departure int
	platform  int
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

func find_max_trains_that_can_run_without_overlapping(schs []Schedule, n int) int {
	slices.SortFunc(schs, func(a, b Schedule) int {
		if a.departure == b.departure {
			return cmp.Compare(b.arrival, a.arrival)
		}
		if a.departure < b.departure {
			return -1
		}
		return 1
	}) // ascending order of departure.

	cnt := 1
	last := schs[0].departure
	for i := 1; i < n; i++ {
		if schs[i].arrival > last {
			cnt++
			last = schs[i].departure
		}
	}

	return cnt
}

func find_maximum_trains(schedules []Schedule, trainCnt, platformCnt int) int {
	platform2ScheduleMap := map[int][]Schedule{}
	for i := 0; i < trainCnt; i++ {
		platform2ScheduleMap[schedules[i].platform] = append(platform2ScheduleMap[schedules[i].platform], schedules[i])
	}
	res := 0
	for _, v := range platform2ScheduleMap {
		res += find_max_trains_that_can_run_without_overlapping(v, len(v))
	}

	return res
}

/*

3 6
1000 1030 1
1010 1030 1
1000 1020 2
1030 1230 2
1200 1230 3
900 1005 1

*/

func main() {
	n := readInt() // number of platforms.
	m := readInt() // m train schedules
	schedules := make([]Schedule, m)
	for i := 0; i < m; i++ {
		schedules[i] = Schedule{
			arrival:   readInt(),
			departure: readInt(),
			platform:  readInt(),
		}
	}

	fmt.Println(find_maximum_trains(schedules, m, n))
}
