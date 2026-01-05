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

func last_remaining_element_jospheh_problem(n, k int) int {
	cnt := n
	it := 0
	for cnt != n {
		it = (it + k - 1) % n
		cnt++
	}
	return 0
}

func main() {
	n := readInt()
	k := readInt()

	fmt.Printf("last_remaining_element_jospheh_problem(n, k): %v\n", last_remaining_element_jospheh_problem(n, k))
}
