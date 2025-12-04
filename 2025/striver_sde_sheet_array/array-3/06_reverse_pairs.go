package main

import (
	"fmt"
	"log"
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
	first  int
	second int
}

var pairs map[int][]Pair = map[int][]Pair{}

func counting_reverse_pairs(arr []int, start int, mid int, end int) int {
	i := start
	j := mid + 1

	res := 0
	for i <= mid {
		flag := false
		for j <= end && arr[i] > 2*arr[j] {
			flag = true
			if _, ok := pairs[i]; !ok {
				pairs[i] = []Pair{}
			}

			pairs[i] = append(pairs[i], Pair{
				first:  arr[i],
				second: arr[j],
			})

			j++
		}
		if flag {
			// j must have moved to next element index or (end or array + 1), so we subtract by 1
			res += j - mid - 1
			for _, v := range pairs[i-1] {
				v.first = arr[i]
				pairs[i] = append(pairs[i], v)
			}
		}
		i++
	}

	return res
}

func merge(arr []int, start int, mid int, end int) {
	brr := make([]int, end-start+1)
	curr := 0
	i := start
	j := mid + 1

	for i <= mid && j <= end {
		if arr[i] <= arr[j] {
			brr[curr] = arr[i]
			i++
			curr++
		} else {
			brr[curr] = arr[j]
			curr++
			j++

		}
	}

	for i <= mid {
		brr[curr] = arr[i]
		curr++
		i++
	}

	for j <= end {
		brr[curr] = arr[j]
		j++
		curr++
	}

	for i := start; i <= end; i++ {
		arr[i] = brr[i-start]
	}
}

func count_reverse_pairs_merge(arr []int, start int, mid int, end int) int {
	// fmt.Println("count_reverse_pairs_merge : ", start, mid, end)
	res := counting_reverse_pairs(arr, start, mid, end)
	// fmt.Println("counting_reverse_pairs internal : ", start, mid, end, res)
	merge(arr, start, mid, end)
	// fmt.Println("merge complete internal : ", start, mid, end, res)

	return res
}

func count_reverse_pairs_helper(arr []int, start int, end int) int {
	if start >= end {
		return 0
	}
	res := 0
	mid := (start + end) / 2
	res += count_reverse_pairs_helper(arr, start, mid)
	res += count_reverse_pairs_helper(arr, mid+1, end)
	res += count_reverse_pairs_merge(arr, start, mid, end)
	return res
}

func count_reverse_pairs(arr []int, n int) int {
	return count_reverse_pairs_helper(arr, 0, n-1)
}

func main() {
	n := readInt()
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = readInt()
	}

	fmt.Printf("count_reverse_pairs(arr, n): %v\n", count_reverse_pairs(arr, n))
	for _, v := range pairs {
		for i := 0; i < len(v); i++ {
			fmt.Printf("%d %d\n", v[i].first, v[i].second)
		}
		fmt.Println()
	}
}
