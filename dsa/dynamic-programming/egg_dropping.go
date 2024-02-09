package main

import (
	"fmt"
	"log"
	"math"
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

func egg_dropping_recursive(e, f int) int {
	if f == 0 || f == 1 || e == 1 {
		return f
	}
	res := math.MaxInt
	for k := 1; k <= f; k++ { // we have k starting form floor 1 upto f'th floor
		breaking_case := egg_dropping_recursive(e-1, k-1)
		not_breaking_case := egg_dropping_recursive(e, f-k)
		temp := 1 + max(breaking_case, not_breaking_case)
		res = min(res, temp)
	}
	return res
}

func egg_dropping_memoization_helper(e, f int, dp [][]int) int {
	if f == 0 || f == 1 || e == 1 {
		dp[e][f] = f
		return dp[e][f]
	}

	if dp[e][f] != 0 {
		return dp[e][f]
	}
	res := math.MaxInt
	for k := 1; k <= f; k++ {
		breaking_case := egg_dropping_memoization_helper(e-1, k-1, dp)
		not_breaking_case := egg_dropping_memoization_helper(e, f-k, dp)
		temp := 1 + max(breaking_case, not_breaking_case)
		res = min(res, temp)
	}
	dp[e][f] = res
	return dp[e][f]
}

func egg_dropping_memoization(e, f int) int {
	dp := make([][]int, 1+e)
	for i := 0; i < 1+e; i++ {
		dp[i] = make([]int, f+1)
	}

	egg_dropping_memoization_helper(e, f, dp)
	return dp[e][f]
}

func egg_dropping(e, f int) int {
	// we have to find the minimum number of attempts required to find the floor where egg does not break.
	// we have e number of eggs and f floors.

	// if egg count is 0 => not a valid input
	// if egg count is 1 => we have to check all floors f => f
	// if f == 0 or 1 => f

	// if we check at certain floor, say k
	// either the egg will break or not break
	// if it not breaks we are supposed to find the floor in remaining f-k floors upside to k with e eggs remaining.
	// if it breaks, we are supposed to find the floor in k - 1 floors downside to k with e-1 eggs remaining.

	res_recrusive := egg_dropping_recursive(e, f)
	res_memoization := egg_dropping_memoization(e, f)
	fmt.Println("Egg dropping recursive : ", res_recrusive)
	fmt.Println("egg dropping memoization : ", res_memoization)
	return res_recrusive
}

func main() {
	e, f := readInt(), readInt()

	fmt.Println(egg_dropping(e, f))
}
