package main

import (
	"fmt"
	"log"
	"math"
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

func kandane_algorithm(arr []int, n int) int {
	sum_till_now := arr[0]
	max_till_now := math.MinInt
	global_start := 0
	global_end := 0
	local_start := 0

	for i := 0; i < n; i++ {
		if sum_till_now == 0 {
			local_start = i
		}

		sum_till_now += arr[i]

		if sum_till_now > max_till_now {
			max_till_now = sum_till_now
			global_start = local_start
			global_end = i
		}

		if sum_till_now < 0 {
			sum_till_now = 0
		}
	}
	fmt.Println("range : ", global_start, global_end)
	fmt.Println(arr[global_start : global_end+1])
	return max_till_now
}

func main() {
	n := readInt()
	arr := make([]int, n)

	for i := 0; i < n; i++ {
		arr[i] = readInt()
	}

	print_array(arr, n)
	kandane_algorithm(arr, n)
	print_array(arr, n)
}
