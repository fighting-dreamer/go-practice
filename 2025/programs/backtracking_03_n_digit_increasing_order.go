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

func print_all_n_digit_numbers_with_incresing_digits_helper(n int, it int, last int, res []int) {
	if it == n {
		fmt.Println(res)
		return
	}
	for i := last + 1; i <= 9; i++ {
		res[it] = i
		print_all_n_digit_numbers_with_incresing_digits_helper(n, it+1, i, res)
	}
}

func print_all_n_digit_numbers_with_incresing_digits(n int) {
	res := make([]int, n)
	print_all_n_digit_numbers_with_incresing_digits_helper(n, 0, -1, res)
}

func main() {
	n := readInt()

	print_all_n_digit_numbers_with_incresing_digits(n)
}
