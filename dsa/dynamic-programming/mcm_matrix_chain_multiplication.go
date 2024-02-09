package main

import (
	"fmt"
	"log"
	"math"

	"github.com/liyue201/gostl/ds/pair"
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

func computeCost(a, b, c int) int {
	return a * b * c
}

func mcm_helper(ms []pair.Pair, start, end int) int {
	if start == end {
		return 0
	}
	if start == end-1 {
		return computeCost(ms[start].Front.(int), ms[start].Back.(int), ms[end].Back.(int))
	}

	res := math.MaxInt
	for k := start; k < end; k++ {
		temp1 := mcm_helper(ms, start, k)
		temp2 := mcm_helper(ms, k+1, end)
		// fmt.Println(start, end, k, temp1, temp2)
		res = min(res, temp1+temp2+computeCost(ms[start].Front.(int), ms[k].Back.(int), ms[end].Back.(int)))
	}

	return res
}

func mcm(nums []int, n int) int {
	// minimum cost by grouping certain matrix multiplications under paranthesis
	ms := make([]pair.Pair, n-1)
	for i := 0; i < n-1; i++ {
		ms[i] = *pair.MakePair(nums[i], nums[i+1])
	}

	fmt.Println(ms)

	return mcm_helper(ms, 0, len(ms)-1)
}

func mcm_memoization_helper(ms []pair.Pair, start, end int, dp [][]int) int {
	if start == end {
		return 0
	}

	if dp[start][end] != -1 {
		return dp[start][end]
	}

	if start == end-1 {
		dp[start][end] = computeCost(ms[start].Front.(int), ms[start].Back.(int), ms[end].Back.(int))
		return dp[start][end]
	}

	res := math.MaxInt
	for k := start; k < end; k++ {
		temp1 := mcm_memoization_helper(ms, start, k, dp)
		temp2 := mcm_memoization_helper(ms, k+1, end, dp)
		res = min(res, temp1+temp2+computeCost(ms[start].Front.(int), ms[k].Back.(int), ms[end].Back.(int)))
	}
	dp[start][end] = res

	return dp[start][end]
}

func mcm_memoization(nums []int, n int) int {
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			dp[i][j] = -1
		}
	}

	ms := make([]pair.Pair, n-1)
	for i := 0; i < n-1; i++ {
		ms[i] = *pair.MakePair(nums[i], nums[i+1])
	}

	mcm_memoization_helper(ms, 0, len(ms)-1, dp)
	// for i := 0; i < n; i++ {
	// 	for j := 0; j < n; j++ {
	// 		fmt.Print(dp[i][j], " ")
	// 	}
	// 	fmt.Println()
	// }

	return dp[0][len(ms)-1]
}

func main() {
	n := readInt()
	nums := make([]int, n)

	for i := 0; i < n; i++ {
		nums[i] = readInt()
	}

	fmt.Println(mcm(nums, n))
	fmt.Println(mcm_memoization(nums, n))
}
