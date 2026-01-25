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

func fib_n(n int) int {
	if n >= 0 && n <= 1 {
		return n
	}
	a := 0
	b := 1
	for i := 2; i <= n; i++ {
		c := a + b
		a = b
		b = c
	}
	return b
}

func main() {
	n := readInt()
	fmt.Printf("Climbing stairs for %d : %d", n, fib_n(n+1))
}
