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

func is_interleaving_recursive_helper(s1, s2, s3 string, itr1, itr2, itr3 int, l1, l2, l3 int) bool {
	if itr3 == l3 {
		return true
	}
	if itr1 == l1 {
		// cehck if remaining from s2 matches s3 remaining
		// if so => return true
		for i := 0; i < l2-itr2; i++ {
			if s2[i+itr2] != s3[itr3+i] {
				return false
			}
		}
		return true
	}
	if itr2 == l2 {
		// cehck if remaining from s1 matches s3 remaining
		// if so => return true
		for i := 0; i < l1-itr1; i++ {
			if s1[i+itr1] != s3[itr3+i] {
				return false
			}
		}
		return true
	}

	if s1[itr1] == s3[itr3] {
		if is_interleaving_recursive_helper(s1, s2, s3, itr1+1, itr2, itr3+1, l1, l2, l3) {
			return true
		}
	}
	if s2[itr2] == s3[itr3] {
		return is_interleaving_recursive_helper(s1, s2, s3, itr1, itr2+1, itr3+1, l1, l2, l3)
	}
	return false
}

func is_interleaving_recursive(s1, s2, s3 string) bool {

	l1 := len(s1)
	l2 := len(s2)
	l3 := len(s3)
	if l1+l2 != l3 {
		return false
	}
	return is_interleaving_recursive_helper(s1, s2, s3, 0, 0, 0, l1, l2, l3)
}

func is_interleaving_iterative_using_dp(s1, s2, s3 string) bool {
	l1, l2 := len(s1), len(s2)
	l3 := len(s3)
	if l1+l2 != l3 {
		return false
	}

	dp := make([][]bool, 1+l1)
	for i := 0; i < 1+l1; i++ {
		dp[i] = make([]bool, 1+l2)
	}

	for i := 0; i < 1+l1; i++ {
		for j := 0; j < 1+l2; j++ {
			if i == 0 && j == 0 {
				dp[i][j] = true
			} else if i == 0 {
				dp[0][j] = dp[0][j-1] && s2[j-1] == s3[j-1]
			} else if j == 0 {
				dp[i][0] = dp[i-1][0] && s1[i-1] == s3[i-1]
			} else {
				dp[i][j] = (dp[i-1][j] && s1[i-1] == s3[i+j-1]) || (dp[i][j-1] && s2[j-1] == s3[i+j-1])
			}

		}
	}

	return dp[l1][l2]
}

func is_interleaving(s1, s2, s3 string) bool {
	res1 := is_interleaving_recursive(s1, s2, s3)
	res2 := is_interleaving_iterative_using_dp(s1, s2, s3)
	fmt.Println(res1, res2)
	return res1
}

func main() {
	s1 := readString()
	s2 := readString()
	s3 := readString()

	fmt.Printf("is_interleaving(s1, s2, s3): %v\n", is_interleaving(s1, s2, s3))
}
