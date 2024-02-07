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

func knapsack_01(profit []int, wt []int, n, cap int) int {
	// we want to fill the bag of capacity "cap" with items say i, with each item, we earn a profit of profit[i] and fill the capacity of wt[i]
	// maximize the profits we can earn.

	// profit earned depend on cap, n
	// there are (cap + 1)* (n + 1) * states DAG 1 is added for zero value.
	// the state is the result, choices, some other info
	// we want to just figure out the max profit.

	dp := make([][]int, cap+1)
	for i := 0; i < cap+1; i++ {
		dp[i] = make([]int, n+1)
	}
	dp[0][0] = 0
	for i := 1; i < n+1; i++ {
		dp[0][i] = 0
	}
	for i := 1; i < cap+1; i++ {
		dp[i][0] = 0
	}

	for i := 1; i < cap+1; i++ {
		for j := 1; j < n+1; j++ {
			if i >= wt[j-1] {
				// coz we are choosing the item and choosing it only once, we are going for dp[i - wt[j - 1]][j - 1]
				// if we can choose the same item multiple (integer) number of times
				dp[i][j] = max(profit[j-1]+dp[i-wt[j-1]][j-1], dp[i][j-1])
			} else {
				dp[i][j] = dp[i][j-1]
			}
		}
	}

	for i := 0; i < cap+1; i++ {
		for j := 0; j < n+1; j++ {
			fmt.Print(dp[i][j], " ")
		}
		fmt.Println()
	}

	return dp[cap][n]
}

func knapsack_01_2(profit, wt []int, n, cap int) int {
	dp := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, cap+1)
	}

	dp[0][0] = 0
	for i := 1; i < n+1; i++ {
		dp[i][0] = 0
	}
	for i := 1; i < cap+1; i++ {
		dp[0][i] = 0
	}

	for i := 1; i < n+1; i++ {
		for j := 1; j < cap+1; j++ {
			if wt[i-1] <= j {
				dp[i][j] = max(profit[i-1]+dp[i-1][j-wt[i-1]], dp[i-1][j])
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}

	for i := 0; i < n+1; i++ {
		for j := 0; j < cap+1; j++ {
			fmt.Print(dp[i][j], " ")
		}
		fmt.Println()
	}

	return dp[n][cap]
}

func knapsack_01_recursive_helper(profit, wt []int, n, cap, profitTillNow, weightTillNow int) int {
	if n == 0 || cap == 0 {
		return profitTillNow
	}

	if wt[n-1] > weightTillNow {
		return knapsack_01_recursive_helper(profit, wt, n-1, cap, profitTillNow, weightTillNow)
	}
	return max(
		knapsack_01_recursive_helper(profit, wt, n-1, cap, profit[n-1]+profitTillNow, weightTillNow-wt[n-1]),
		knapsack_01_recursive_helper(profit, wt, n-1, cap, profitTillNow, weightTillNow))
}

func knapsack_01_recursive(profit, wt []int, n, cap int) int {

	return knapsack_01_recursive_helper(profit, wt, n, cap, 0, cap)
}

func main() {
	n, cap := readInt(), readInt()
	profit := make([]int, n)
	wt := make([]int, n)
	for i := 0; i < n; i++ {
		profit[i] = readInt()
	}

	for i := 0; i < n; i++ {
		wt[i] = readInt()
	}

	fmt.Println("bottom up with iteration on capacity weight : ", knapsack_01(profit, wt, n, cap))
	fmt.Println("bottom up with iteration on items : ", knapsack_01_2(profit, wt, n, cap))

	fmt.Println("Recursive : ", knapsack_01_recursive(profit, wt, n, cap))
}
