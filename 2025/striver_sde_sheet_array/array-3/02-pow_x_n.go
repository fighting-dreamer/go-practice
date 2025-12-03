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

func pow_(x, n int) float64 {
	isNegFlag := false
	if n < 0 {
		n = -n
		isNegFlag = true
	}
	res := 1
	p := x
	for n > 0 {
		if n&1 == 1 {
			// n is odd
			res = res * p // multiple to the power of x
			n--           // reduce the multiplication for that power by 1
		} else {
			n = n >> 1 // divide by 2
			p = p * p  // square existing power of x
		}
	}
	if isNegFlag {
		return 1.0 / float64(res)
	}
	return float64(res)
}

func main() {
	x := readInt()
	n := readInt()
	fmt.Printf("pow_(x,n): %v\n", pow_(x, n))
}
