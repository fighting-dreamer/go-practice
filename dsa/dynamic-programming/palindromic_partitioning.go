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

func palindromic_partitioning(str string) int {
	l := len(str)
	isPalindrome := make([][]bool, l)
	countPartitions := make([][]int, l)
	for i := 0; i < l; i++ {
		isPalindrome[i] = make([]bool, l)
		countPartitions[i] = make([]int, l)
	}

	for i := 0; i < l; i++ {
		isPalindrome[i][i] = true
		countPartitions[i][i] = 0
	}

	for li := 2; li <= l; li++ {
		for start := 0; start <= l-li; start++ {
			end := start + li - 1
			if li == 2 {
				isPalindrome[start][end] = str[start] == str[end]
			} else {
				isPalindrome[start][end] = str[start] == str[end] && isPalindrome[start+1][end-1]
			}

			if isPalindrome[start][end] == true {
				countPartitions[start][end] = 0
			} else {
				res := math.MaxInt
				for k := start; k < end; k++ {
					temp := countPartitions[start][k] + countPartitions[k+1][end] + 1
					res = min(res, temp)
				}
				countPartitions[start][end] = res
			}
		}
	}

	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {
			fmt.Print(isPalindrome[i][j], " ")
		}
		fmt.Println()
	}
	fmt.Println("--------------------")
	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {
			fmt.Print(countPartitions[i][j])
		}
		fmt.Println()
	}

	return countPartitions[0][l-1]
}

func main() {
	str := readString()
	fmt.Println(palindromic_partitioning(str))
}
