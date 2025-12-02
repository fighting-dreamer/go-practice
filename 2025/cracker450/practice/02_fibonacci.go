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

func fib_memo(dp map[int]int, n int) int {
	if n == 0 {
		dp[n] = 0
	} else if n <= 2 {
		dp[n] = 1
	} else {
		dp[n] = fib_memo(dp, n-1) + fib_memo(dp, n-2)
	}
	return dp[n]
}

func fibonacci(n int) int {
	dp := map[int]int{}
	return fib_memo(dp, n)
}

func main() {
	n := readInt()
	fmt.Println(fibonacci(n))
}
