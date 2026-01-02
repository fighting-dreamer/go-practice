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

func isPalindrome(sb []byte) func(i, j int) bool {
	flag := true
	n := len(sb)
	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
	}
	dp[0][0] = true
	for i := 0; i < n-1; i++ {
		dp[i+1][i+1] = true
		dp[i][i+1] = sb[i] == sb[i+1]
	}
	for l := 3; l <= n; l++ {
		for i := 0; i <= n-l; i++ {
			j := i + l - 1
			dp[i][j] = dp[i+1][j-1] && sb[i] == sb[j]
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Printf("%9v", dp[i][j])
		}
		fmt.Println()
	}

	return func(i, j int) bool {
		fmt.Println(flag, i, j, dp[i][j])
		return dp[i][j]
	}
}

func copySlice(arr [][]byte) [][]byte {
	n := len(arr)
	temp := make([][]byte, n)
	for i := 0; i < n; i++ {
		temp[i] = make([]byte, len(arr[i]))
		copy(temp[i], arr[i])
	}
	return temp
}

func palindrome_partition_helper(sb []byte, it int, n int, isPalindromeFunc func(i, j int) bool, curr [][]byte, res *[][][]byte) {
	if it == n {
		*res = append(*res, copySlice(curr))
		return
	}
	for i := it; i < n; i++ {
		if isPalindromeFunc(it, i) {
			curr = append(curr, sb[it:i+1]) // it inclusive, i+1 exclusive
			palindrome_partition_helper(sb, i+1, n, isPalindromeFunc, curr, res)
			curr = curr[:len(curr)-1]
		}
	}
}

func palindrome_partitions(sb []byte, n int) [][][]byte {
	res := [][][]byte{}
	curr := [][]byte{}

	it := 0
	isPalindromeFunc := isPalindrome(sb)

	palindrome_partition_helper(sb, it, n, isPalindromeFunc, curr, &res)
	return res
}

func main() {
	s := readString()
	sb := []byte(s)
	n := len(s)

	res := palindrome_partitions(sb, n)
	for i := 0; i < len(res); i++ {
		for j := 0; j < len(res[i]); j++ {
			fmt.Printf("%s|", string(res[i][j]))
		}
		fmt.Println()
	}
}
