package main

import (
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

func kidsWithMaxCandies(candies []int, extraCandies int) []bool {
	maxElement := slices.Max(candies)
	n := len(candies)
	res := make([]bool, n)
	for i := 0; i < n; i++ {
		if candies[i]+extraCandies >= maxElement {
			res[i] = true
		}
	}

	return res
}

func main() {
	n := readInt()
	candies := make([]int, n)
	for i := 0; i < n; i++ {
		candies[i] = readInt()
	}
	extraCandies := readInt()
	fmt.Println(kidsWithMaxCandies(candies, extraCandies))
}
