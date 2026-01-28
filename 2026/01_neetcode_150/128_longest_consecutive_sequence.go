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

func longest_consecutive_sequence(arr []int, n int) int {
	numSet := map[int]bool{}
	for i := 0; i < n; i++ {
		numSet[arr[i]] = true
	}
	maxStreak := 0
	for i := 0; i < n; i++ {
		if numSet[arr[i]-1] == false {
			// arr[i] is the start of the possible sequence coz arr[i] - 1 is not in the set
			curr := arr[i]
			streak := 1
			for numSet[curr+1] {
				streak++
				curr = curr + 1
			}
			maxStreak = max(streak, maxStreak)
		}
	}

	return maxStreak
}

func main() {
	n := readInt()
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = readInt()
	}

	fmt.Printf("longest_consecutive_sequence(arr, n): %v\n", longest_consecutive_sequence(arr, n))
}
