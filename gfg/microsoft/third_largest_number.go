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

func process(a, b, c, d int) (int, int, int) {
	if d > a {
		c = b
		b = a
		a = d
	} else if d > b {
		c = b
		b = d
	} else if d > c {
		c = d
	}

	return a, b, c
}

func third_largest_element(nums []int, n int) int {
	first, third := max(nums[0], nums[1], nums[2]), min(nums[0], nums[1], nums[2])
	second := nums[0] + nums[1] + nums[2] - first - third

	for i := 3; i < n; i++ {
		first, second, third = process(first, second, third, nums[i])
	}

	return third
}

func main() {
	n := readInt()
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = readInt()
	}

	fmt.Println(third_largest_element(nums, n))
}
