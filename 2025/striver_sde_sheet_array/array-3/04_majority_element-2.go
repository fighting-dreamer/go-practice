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

func using_moore_algorithm_for_majority_element_ndiv3(arr []int, n int) []int {
	currAssumption1 := arr[0]
	currAssumption2 := arr[1]
	count0 := 1
	count1 := 1
	// starting from index 1 (zero based indicies),
	// you can start from index-0, but just keep count as zero then
	for i := 2; i < n; i++ {
		if count1 == 0 && currAssumption2 != arr[i] {
			currAssumption1 = arr[i]
		}

		if count1 == 0 && currAssumption1 != arr[i] {
			currAssumption2 = arr[i]
		}

		if arr[i] == currAssumption1 {
			count1++
		} else if arr[i] == currAssumption2 {
			count1++
		} else {
			count0--
			count1--
		}
	}

	count0 = 0
	count1 = 0
	res := []int{}
	for i := 0; i < n; i++ {
		if arr[i] == currAssumption1 {
			count0++
		}
		if arr[i] == currAssumption2 {
			count1++
		}
	}
	if count0 > n/3 {
		res = append(res, currAssumption1)
	}
	if count1 > n/3 {
		res = append(res, currAssumption2)
	}
	return res
}

func majority_element_2(arr []int, n int) []int {
	// using moore algo
	/*
		Initialize four variables: cnt1 and cnt2 for tracking the counts of elements, and el1 and el2 for storing the potential majority elements.
		Traverse through the given array:
		If cnt1 is 0 and the current element is not equal to el2, set el1 to the current element and increment cnt1 by 1.
		If cnt2 is 0 and the current element is not equal to el1, set el2 to the current element and increment cnt2 by 1.
		If the current element is equal to el1, increment cnt1 by 1.
		If the current element is equal to el2, increment cnt2 by 1.
		In all other cases, decrease cnt1 and cnt2 by 1.
		After processing all elements, el1 and el2 should be the candidate elements for majority. To confirm:
		Use another loop to manually check the counts of el1 and el2 in the array.
		If either el1 or el2's count is greater than floor(N/3), it is considered a valid majority element.
	*/

	res := using_moore_algorithm_for_majority_element_ndiv3(arr, n)
	fmt.Printf("res: %v\n", res)
	return res
}

func main() {
	n := readInt()
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = readInt()
	}

	majority_element_2(arr, n)
}
