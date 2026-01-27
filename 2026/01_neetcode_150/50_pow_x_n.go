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

func pow_x_n(a int, n int) int {
	// 10 => 1010
	// 1010 => 101 => 10 => 1
	// a => a * a => (a * a)^2 * a => ((a * a)^2 * a)^
	// (a^4 * a)^2 <=> a^10

	res := 1
	for n > 0 {
		if n&1 == 1 {
			res = res * a
		}
		n = n >> 1
		a = a * a
	}

	return res
}

func main() {
	x := readInt()
	n := readInt()

	pow_x_n(x, n)
}
