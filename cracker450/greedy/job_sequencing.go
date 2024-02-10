package main

import (
	"cmp"
	"fmt"
	"log"
	"slices"
)

type Job struct {
	Id       int
	Deadline int
	Weight   int
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

func find_max_profit(jobs []Job, n int) (int, int) {
	slices.SortFunc(jobs, func(a Job, b Job) int {
		if a.Weight == b.Weight {
			return cmp.Compare(b.Deadline, a.Deadline)
		}
		if a.Weight < b.Weight {
			return 1
		}
		return -1
	})
	fmt.Println(jobs)
	fmt.Println("-----------------------")
	profitWeight := 0
	cnt := 0
	slots := make([]int, n)
	for i := 0; i < n; i++ {
		for j := min(n, jobs[i].Deadline) - 1; j >= 0; j-- {
			if slots[j] == 0 {
				slots[j] = 1
				cnt++
				fmt.Println(jobs[i])
				profitWeight += jobs[i].Weight
				break
			}
		}
	}

	return cnt, profitWeight
}

func main() {
	n := readInt()
	jobs := make([]Job, n)
	for i := 0; i < n; i++ {
		jobs[i] = Job{
			Id:       readInt(),
			Deadline: readInt(),
			Weight:   readInt(),
		}
	}
	fmt.Println(jobs)
	fmt.Println(find_max_profit(jobs, n))
}
