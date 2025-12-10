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

type Pair struct {
	First  int
	Second int
}

// first is the deadline, second is the profit
func job_sequencing(arr []Pair, n int) int {
	mp := map[int]int{}
	slices.SortFunc(arr, func(a Pair, b Pair) int {
		return cmp.Compare(b.Second, a.Second)
	})
	fmt.Println(arr)
	maxProfit := 0
	jobs := []Pair{}
	for i := 0; i < n; i++ {
		for j := arr[i].First; j > 0; j-- {
			_, ok := mp[j]
			if !ok {
				mp[j] = 1
				maxProfit += arr[i].Second
				jobs = append(jobs, arr[i])
				break
			}
		}
	}
	fmt.Println(jobs)
	return maxProfit
}

func main() {
	n := readInt()
	arr := make([]Pair, n)
	for i := 0; i < n; i++ {
		arr[i] = Pair{
			First:  readInt(),
			Second: readInt(),
		}
	}

	res := job_sequencing(arr, n)
	fmt.Printf("res: %v\n", res)
}
