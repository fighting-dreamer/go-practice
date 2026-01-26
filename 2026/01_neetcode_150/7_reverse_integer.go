package main

import (
	"fmt"
	"log"
	"slices"
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

func trim(n int) int {
	for n%10 == 0 {
		n = n / 10
	}

	return n
}

func getDigits(n int) []int {
	res := []int{}
	for n >= 9 {
		res = append(res, n%10)
		n = n / 10
	}
	res = append(res, n)
	return res
}

func digitsListToNum(digits []int) int {
	res := 0
	for i := len(digits) - 1; i >= 0; i-- {
		res = res*10 + digits[i]
	}

	return res
}

// 321 => 123, -312 => -123, 120 => 21
func main() {
	n := readInt()
	digits := getDigits(trim(n))
	slices.Reverse(digits) // digits reversed
	res := digitsListToNum(digits)
	fmt.Println(res)
}
