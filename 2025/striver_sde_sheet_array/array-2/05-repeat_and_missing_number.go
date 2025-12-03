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

func linear_sum(arr []int, n int) int {
	res := 0
	for i := 0; i < n; i++ {
		res += arr[i]
	}

	return res
}

func square_sum(arr []int, n int) int {
	res := 0
	for i := 0; i < n; i++ {
		res += arr[i] * arr[i]
	}

	return res
}

func using_linear_and_square_sums(arr []int, n int) (int, int) {
	arrSum := linear_sum(arr, n)
	sn := (n * (n + 1) / 2)
	a := arrSum - sn

	arrSqSum := square_sum(arr, n)
	ssn := (n * (n + 1) * (2*n + 1)) / 6
	b := arrSqSum - ssn

	X := (b + a*a) / (2 * a)
	Y := (b - a*a) / (2 * a)

	return X, Y
}

func rightmost_bit(n int) int {
	return n & (-n)
}

func using_xor_approach(arr []int, n int) (int, int) {
	xor := 1
	for i := 0; i < n; i++ {
		xor = (xor ^ arr[i]) ^ (i + 1)
	}

	mask := rightmost_bit(xor)

	zero_xor := 0
	one_xor := 0
	for i := 1; i <= n; i++ {
		if (arr[i-1] & mask) == 0 {
			// b'th bit is zero
			zero_xor = zero_xor ^ arr[i-1]
		} else {
			// b'th bit is set to 1
			one_xor = one_xor ^ arr[i-1]
		}

		if (mask & i) == 0 {
			// b'th bit is zero
			zero_xor = zero_xor ^ i
		} else {
			// b'th bit is set to 1
			one_xor = one_xor ^ i
		}
	}
	// zero_xor, one_xor : one is repeating other missing
	for i := 0; i < n; i++ {
		if arr[i] == zero_xor {
			return zero_xor, one_xor
		} else if arr[i] == one_xor {
			return one_xor, zero_xor
		}
	}
	return -1, -1
}

func repeat_and_missing_number(arr []int, n int) {
	// you can create two equations :
	// say X is repeating one and Y is the missing one
	// Equaiton one : (sum of elements) - sum of (1-N) = X - Y
	// Equation two : (sum of square of elements) - (sum of square of element(1-N)) = X^2 - Y^2
	// solve the equations to get X and Y
	// ----------------------------
	// using XOR approach
	// we get the xor of all elements, and (1-N) elements => it will not be zero
	// the bit at which it differ, is the one different in X and Y
	// for each element, you put it to zero or one category, you actually xor it and do the xor with (1-N) too
	// we know if say i is repeating, it will go to same category of zero or one in all repetitions
	// and hte missing set will have only one entry for the missing number
	// xor in either category => missing or repeating element
	// we find out which one is repeating or missing by checking original array in linear time.

	a, b := using_linear_and_square_sums(arr, n)
	fmt.Printf("repeating: %v, missing: %v\n", a, b)

	c, d := using_xor_approach(arr, n)
	fmt.Printf("repeating: %v, missing: %v\n", c, d)
}

func main() {
	n := readInt()
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = readInt()
	}

	repeat_and_missing_number(arr, n)
}
