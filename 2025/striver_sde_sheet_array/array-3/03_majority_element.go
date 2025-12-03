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

func majority_element_using_hash_map(arr []int, n int) int {
	mp := map[int]int{}
	for i := 0; i < n; i++ {
		mp[arr[i]]++
	}
	for k, v := range mp {
		if v > n/2 {
			return k
		}
	}
	return -1
}

func using_moore_algorithm_for_majority_element(arr []int, n int) int {
	currAssumption := arr[0]
	count := 1
	// starting from index 1 (zero based indicies),
	// you can start from index-0, but just keep count as zero then
	for i := 1; i < n; i++ {
		if count == 0 {
			currAssumption = arr[i]
		}

		if arr[i] == currAssumption {
			count++
		} else {
			count--
		}
	}

	return currAssumption
}

func majority_element(arr []int, n int) int {
	// approach 1: using a hash map to keep counts
	// approach 2 : moore algorithm
	// - if an element definitey comes more than n/2 times, we can cancel the count of that element v count of others elements
	// we keep track of an element and assume it to be majority and
	// cancel its count if we encounter any other element till-
	// -till the count become zero to assume other element as majority

	res1 := majority_element_using_hash_map(arr, n)
	res2 := using_moore_algorithm_for_majority_element(arr, n)
	fmt.Printf("res1: %v\n", res1)
	fmt.Printf("res2: %v\n", res2)
	return res1
}

func main() {
	n := readInt()
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = readInt()
	}

	majority_element(arr, n)
}
