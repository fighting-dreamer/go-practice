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

func find_max_struct_height(nums [][]int, n int) int {
	height := func(s []int) int {
		return s[0]*12 + s[1]
	}
	res := 0
	for i := 1; i < n; i++ {
		if height(nums[res]) < height(nums[i]) {
			res = i
		}
	}

	return height(nums[res])
}

func main() {
	n := readInt()
	nums := make([][]int, n)
	for i := 0; i < n; i++ {
		nums[i] = []int{readInt(), readInt()}
	}

	fmt.Println(find_max_struct_height(nums, n))
}
