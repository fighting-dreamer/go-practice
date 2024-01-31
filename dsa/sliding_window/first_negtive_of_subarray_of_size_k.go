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

func isNegtive(a int) bool {
	return a < 0
}

func first_negtive_of_all_subarray_of_size_k(nums []int, n, k int) []int {
	if k > n {
		return []int{}
	}
	negs := []int{}
	res := []int{}
	for i, j := 0, 0; i < n && j < n; {
		if nums[j] < 0 {
			negs = append(negs, nums[j])
		}

		if j-i+1 < k {
			j++
		} else if j-i+1 == k {
			if len(negs) > 0 {
				res = append(res, negs[0])
			} else {
				res = append(res, 0)
			}

			if len(negs) > 0 && nums[i] == negs[0] {
				negs = negs[1:]
			}

			i++
			j++
		}
	}

	return res
}

// input : 8 3 12 -1 -7 8 -15 30 16 28
// output :
func main() {
	n, k := readInt(), readInt()
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = readInt()
	}

	fmt.Println(first_negtive_of_all_subarray_of_size_k(nums, n, k))
}
