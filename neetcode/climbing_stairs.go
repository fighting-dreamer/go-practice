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

func climbing_stairs(n int) int {
	if n == 0 || n == 1 {
		return 1
	}

	// n = 0 => 1
	// n = 1 => 1
	// n = 2 => f(n - 1) + f(n - 2)
	a, b := 1, 1
	for i := 2; i <= n; i++ {
		temp := b
		b = a + b
		a = temp
	}
	return b
}

func main() {
	n := readInt()

	fmt.Println(climbing_stairs(n))
}
