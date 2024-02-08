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

func sub_set_sum(set []int, n, s int) bool {
	dp := make([][]bool, 1+n)
	for i := 0; i < 1+n; i++ {
		dp[i] = make([]bool, 1+s)
	}

	for i := 0; i < n+1; i++ {
		dp[i][0] = true
	}

	for i := 1; i < 1+s; i++ {
		dp[0][i] = false
	}

	for i := 1; i < 1+n; i++ {
		for j := 1; j < 1+s; j++ {
			if set[i-1] > j {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j] || dp[i-1][j-set[i-1]]
			}
		}
	}

	return dp[n][s]
}

func equal_sum_partition(s []int, n int) bool {
	sum := func(arr []int) int {
		res := 0
		for i := 0; i < len(arr); i++ {
			res += arr[i]
		}

		return res
	}(s)
	if sum%2 != 0 {
		return false
	}

	return sub_set_sum(s, n, sum/2)
}

func main() {
	n := readInt()

	s := make([]int, n)
	for i := 0; i < n; i++ {
		s[i] = readInt()
	}

	fmt.Println(equal_sum_partition(s, n))
}
