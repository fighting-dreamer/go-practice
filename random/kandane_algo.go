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

func kandane_algo(nums []int, n int) []int {
	maxTillNow := 0
	resLeft := 0
	temp := 0
	resRight := 0
	sumTillNow := 0
	for i := 0; i < n; i++ {
		sumTillNow = sumTillNow + nums[i]
		if sumTillNow > maxTillNow {
			maxTillNow = sumTillNow
			resLeft = temp
			resRight = i
		}
		if sumTillNow <= 0 {
			sumTillNow = 0
			temp = i + 1
		}
	}

	fmt.Println(resLeft, resRight, maxTillNow)

	return nums[resLeft : resRight+1]
}

func main() {
	n := readInt()
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = readInt()
	}

	fmt.Println(kandane_algo(nums, n))
}
