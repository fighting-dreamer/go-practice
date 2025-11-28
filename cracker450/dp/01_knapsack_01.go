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

/*
W : weight capacity of knapsack
n : count of items
wt : weight of each item
val : profit value for each item
curr : current index
currWt : currentWeight
*/
func knapsack_01_recusrive(W int, n int, wt []int, val []int, curr int, currWt int) int {
	if curr == n {
		return 0
	}

	left := -1
	if currWt+wt[curr] <= W {
		left = val[curr] + knapsack_01_recusrive(W, n, wt, val, curr+1, currWt+wt[curr])
	}
	right := knapsack_01_recusrive(W, n, wt, val, curr+1, currWt)

	return max(left, right)
}

func knapsack_01_recurisve_memo(dp [][]int, W int, n int, wt []int, val []int, curr int, currWt int) int {
	if curr == n {
		dp[curr][currWt] = 0
	}
	if dp[curr][currWt] != -1 {
		return dp[curr][currWt]
	}

	left := -1
	if currWt+wt[curr] <= W {
		left = val[curr] + knapsack_01_recurisve_memo(dp, W, n, wt, val, curr+1, currWt+wt[curr])
	}
	right := knapsack_01_recurisve_memo(dp, W, n, wt, val, curr+1, currWt)
	dp[curr][currWt] = max(left, right)
	return dp[curr][currWt]
}

func knapsack_01_table(W int, n int, wt []int, val []int) int {
	dp := make([][]int, 1+n)
	for i := 0; i < 1+n; i++ {
		dp[i] = make([]int, 1+W)
		for j := 0; j < 1+W; j++ {
			dp[i][j] = -1
		}
	}

	for i := 0; i < 1+n; i++ {
		dp[i][0] = 0
	}
	for i := 0; i < 1+W; i++ {
		dp[0][i] = 0
	}

	for i := 1; i < 1+n; i++ {
		for j := 1; j < 1+W; j++ {
			dp[i][j] = dp[i-1][j]
			if j >= wt[i-1] {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-wt[i-1]]+val[i-1])
			}
		}
	}

	return dp[n][W]
}

func knapsack_01_table_wt_first(W int, n int, wt []int, val []int) int {
	dp := make([][]int, 1+n)
	for i := 0; i < 1+n; i++ {
		dp[i] = make([]int, 1+W)
		for j := 0; j < 1+W; j++ {
			dp[i][j] = 0
		}
	}

	for i := 0; i < 1+n; i++ {
		dp[i][0] = 0
	}

	for i := 0; i < 1+W; i++ {
		dp[0][i] = 0
	}

	for currWt := 1; currWt < 1+W; currWt++ {
		for index := 1; index < 1+n; index++ {
			dp[index][currWt] = dp[index-1][currWt]
			if currWt >= wt[index-1] {
				dp[index][currWt] = max(
					dp[index-1][currWt],
					val[index-1]+dp[index-1][currWt-wt[index-1]],
				)
			}
		}
	}

	return dp[n][W]
}

func knapsack_helper(W int, wt []int, val []int) int {
	recusrive_res := knapsack_01_recusrive(W, len(wt), wt, val, 0, 0)
	n := len(wt)
	dp := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, W+1)
	}
	for i := 0; i < n+1; i++ {
		for j := 0; j < W+1; j++ {
			dp[i][j] = -1
		}
	}
	recursive_memo_res := knapsack_01_recurisve_memo(dp, W, n, wt, val, 0, 0)
	res_tablulation := knapsack_01_table(W, n, wt, val)
	fmt.Println("knapsack_01_recusrive: ", recusrive_res)
	fmt.Println("knapsack_01_recurisve_memo: ", recursive_memo_res)
	fmt.Printf("knapsack_01_table: %v\n", res_tablulation)
	fmt.Printf("knapsack_01_table_wt_first: %v\n", knapsack_01_table_wt_first(W, n, wt, val))
	return recusrive_res
}

func main() {
	n := readInt()
	W := readInt()
	wt := make([]int, n)
	for i := 0; i < n; i++ {
		wt[i] = readInt()
	}
	val := make([]int, n)
	for i := 0; i < n; i++ {
		val[i] = readInt()
	}
	fmt.Println(knapsack_helper(W, wt, val))
}
