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

func maximum_array(arr []int, n int) []int {
	l := 0
	tempL := 0
	r := 0
	sumTillNow := 0
	maxRes := 0
	for i := 0; i < n; i++ {
		sumTillNow = sumTillNow + arr[i]
		if sumTillNow > maxRes {
			maxRes = sumTillNow
			l = tempL
			r = i
		}

		if sumTillNow < 0 {
			sumTillNow = 0
			tempL = i + 1
		}
	}

	return arr[l : r+1]
}

func main() {
	n := readInt()
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = readInt()
	}

	maximum_array(arr, n)
}
