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

func print_array(arr []int, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", arr[i])
	}
	fmt.Println()
}

func merge(arr []int, start int, mid int, end int) int {
	inversionCnt := 0

	brr := make([]int, end-start+1)
	curr := 0
	i := start
	j := mid + 1
	for i <= mid && j <= end {
		if arr[i] <= arr[j] {
			// no inversion
			brr[curr] = arr[i]
			curr++
			i++
		} else {
			// inversino scenario
			brr[curr] = arr[j]
			curr++
			j++

			inversionCnt += (mid - i + 1) // as zero based indices
		}
	}

	for i <= mid {
		// not inversion scenario

		brr[curr] = arr[i]
		curr++
		i++
	}
	for j <= end {
		// inversion scenario
		brr[curr] = arr[j]
		j++
		curr++

		inversionCnt += (mid - i + 1) // this comes to be 0 as i becomes mid + 1
	}

	for i := start; i <= end; i++ {
		arr[i] = brr[i-start]
	}
	return inversionCnt
}

func count_inversions_using_merge_sort_helper(arr []int, start int, end int) int {
	if start >= end {
		return 0
	}
	mid := (start + end) / 2
	res := 0
	res += count_inversions_using_merge_sort_helper(arr, start, mid)
	res += count_inversions_using_merge_sort_helper(arr, mid+1, end)
	res += merge(arr, start, mid, end)
	return res
}

func count_inversions_using_merge_sort(arr []int, n int) int {
	return count_inversions_using_merge_sort_helper(arr, 0, n-1)
}

func main() {
	n := readInt()
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = readInt()
	}

	fmt.Printf("count_inversions_using_merge_sort(arr, n): %v\n", count_inversions_using_merge_sort(arr, n))
}
