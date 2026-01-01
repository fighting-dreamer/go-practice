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

func ninja_training(arr [][]int, n int) int {
	dp := make([][]int, 1+n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, 4)
	}
	// assuming on day : -1, we did jth task (ie 0, or 1, 2 or nothing), what is the max on 0th day for diff tasks
	dp[0][0] = max(arr[0][1], arr[0][2])
	dp[0][1] = max(arr[0][0], arr[0][2])
	dp[0][2] = max(arr[0][0], arr[0][1])
	dp[0][0] = max(arr[0][0], max(arr[0][1], arr[0][2]))

	for i := 1; i < n; i++ {
		for last := 0; last < 4; last++ {
			dp[i][last] = 0
			for task := 0; task < 3; task++ {
				if task != last {
					dp[i][last] = max(dp[i][last], arr[i][task]+dp[i-1][task])
				}
			}
		}
	}

	return dp[n-1][3]
}

func main() {
	n := readInt()
	arr := make([][]int, n)

	for i := 0; i < n; i++ {
		arr[i] = []int{readInt(), readInt(), readInt()}
	}

	fmt.Println("ninja_training(arr, n) : ", ninja_training(arr, n))
}
