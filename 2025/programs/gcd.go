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

func gcd(a, b int) int {
	if a == 0 || b == 0 {
		return 0
	}
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func main() {
	a := readInt()
	b := readInt()

	fmt.Printf("GCD of %d, %d : %d", a, b, gcd(a, b))
}
