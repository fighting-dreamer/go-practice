package main

func edit_distance(s1 string, s2 string) int {
	l1, l2 := len(s1), len(s2)
	dp := make([][]int, l1+1)
	for i := 0; i < 1+l1; i++ {
		dp[i] = make([]int, 1+l2)
	}

	for i := 0; i < 1+l1; i++ {
		for j := 0; j < 1+l2; j++ {
			if i == 0 {
				// base case : transforming empty-string s1 to s2 till i'ths length
				dp[i][j] = j
				continue
			}
			if j == 0 {
				// base case : transforming empty-string s2 to s1 till i'ths length
				dp[i][j] = i
				continue
			}

			if s1[i-1] == s2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				// min of insert, delete, replace
				dp[i][j] = 1 + min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1])
			}
		}
	}

	return dp[l1][l2]
}

func main() {
	s1 := readString()
	s2 := readString()

	edit_distance(s1, s2)
}
