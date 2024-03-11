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

func can_place_flowers(nums []int, n int) bool {
	l := len(nums)
	for i := 0; i < l; i++ {
		if nums[i] == 0 && (i == 0 || nums[i-1] == 0) && (i == l-1 || nums[i+1] == 0) {
			nums[i] = 1
			n--
		}
		if n <= 0 {
			return true
		}
	}
	return false
}

func main() {
	l, n := readInt(), readInt()
	nums := make([]int, l)
	for i := 0; i < l; i++ {
		nums[i] = readInt()
	}

	fmt.Println("can place flowers : ", can_place_flowers(nums, n))
}
