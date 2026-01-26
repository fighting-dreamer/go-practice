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

func getDigits(n int) []int {
	res := []int{}
	for n > 0 {
		res = append(res, n%10)
		n = n / 10
	}
	return res
}

func squareSum(arr []int) int {
	res := 0
	for i := 0; i < len(arr); i++ {
		res += arr[i] * arr[i]
	}

	return res
}

func is_happy_number(n int) bool {
	for {
		if n%10 == n { // n < 10
			return n == 1
		}
		digitsArr := getDigits(n)
		n = squareSum(digitsArr)
	}
}

func main() {
	n := readInt()
	is_happy_number(n)
}
