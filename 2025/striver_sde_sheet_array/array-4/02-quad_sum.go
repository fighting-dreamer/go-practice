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

func quad_sum(arr []int, n int, target int) int {
	res := 0

	for i := 0; i < n-3; i++ {
		if i > 0 && arr[i] == arr[i-1] {
			continue
		}
		for j := i + 1; j < n-2; j++ {
			if j > i+1 && arr[j] == arr[j-1] {
				continue
			}

			l := j + 1
			r := n - 1
			for l < r {
				sum := arr[i] + arr[j] + arr[l] + arr[r]
				if sum == target {
					res++
					for arr[l] == arr[l+1] && l < r {
						l++
					}
					for arr[r] == arr[r-1] && l < r {
						r--
					}
					l++
					r--
				} else {
					if sum > target {
						r--
					} else {
						l++
					}
				}
			}
		}
	}
	return res
}

func main() {
	n := readInt()
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = readInt()
	}

	target := readInt()

	fmt.Printf("quad_sum(arr, n, target): %v\n", quad_sum(arr, n, target))
}
