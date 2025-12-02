package main

import (
	"fmt"
	"log"
	"math"
)

func readInt() int {
	var n int
	_,err := fmt.Scanf("%d", &n)
	if err != nil {
		log.Fatal("couldn;t read number")
	}
	return n
}

func isPrime(n int) bool {
	if n % 2 == 0 {
		return false
	}

	for i := 3; i <= int(math.Sqrt(float64(n))); i+= 2 {
		if n % i == 0 {
			return false
		}
	}
	return true
}

func main() {
	n := readInt()
	fmt.Println("Is Prime ", n, " : ", isPrime(n))
}