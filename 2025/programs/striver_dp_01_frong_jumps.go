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

// absolute : returns the postive value
func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

func frog_jump(arr []int, n int) int {
	dp := make([]int, n)
	par := make([]int, n)
	dp[0] = 0
	dp[1] = abs(arr[1] - arr[0])
	for i := 2; i < n; i++ {
		a := dp[i-1] + abs(arr[i]-arr[i-1])
		b := dp[i-2] + abs(arr[i]-arr[i-2])
		if a < b {
			par[i] = i - 1
			dp[i] = a
		} else {
			par[i] = i - 2
			dp[i] = b
		}
	}
	i := n - 1
	for i != 0 {
		fmt.Println(i, par[i], abs(arr[par[i]]-arr[i]))
		i = par[i]
	}
	return dp[n-1]
}

func main() {
	n := readInt()

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = readInt()
	}

	fmt.Println("frog_jump(arr, n) : ", frog_jump(arr, n))
}
