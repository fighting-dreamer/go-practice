package main

import (
	"fmt"
	"log"
	"sort"
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

func three_sum_submitted(nums []int) [][]int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	res := [][]int{}
	n := len(nums)
	for i := 0; i < n-2; i++ {
		j := i + 1
		k := n - 1
		target := -nums[i]

		for j < k {
			tmp := nums[j] + nums[k]
			if tmp == target {
				solu := []int{nums[i], nums[j], nums[k]}
				sort.Slice(solu, func(a, b int) bool {
					return solu[a] < solu[b]
				})
				isAlreadyPresent := 0
				for a := 0; a < len(res); a++ {
					if res[a][0] == solu[0] && res[a][1] == solu[1] && res[a][2] == solu[2] {
						isAlreadyPresent = 1
						break
					}
				}
				if isAlreadyPresent == 0 {
					res = append(res, solu)
				}
				// fmt.Println(i, j, k)
				j++
			} else {
				if tmp < target {
					j++
				} else {
					k--
				}
			}
		}
	}

	return res
}

func three_sum(nums []int, n int) [][3]int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	res := [][3]int{}

	for i := 0; i < n-2; i++ {
		j := i + 1
		k := n - 1
		target := -nums[i]

		for j < k {
			tmp := nums[j] + nums[k]
			if tmp == target {
				res = append(res, [3]int{i, j, k})
				fmt.Println(i, j, k)
				j++
			} else {
				if tmp < target {
					j++
				} else {
					k--
				}
			}
		}
	}

	return res
}

func main() {
	n := readInt()
	nums := []int{}
	for i := 0; i < n; i++ {
		nums = append(nums, readInt())
	}

	// fmt.Println(nums)
	// mynums := make([]int, n)
	// c := copy(mynums, nums)

	// fmt.Println("c : ", c)

	// fmt.Println("mynums : ", mynums)
	// copy(mynums, nums)
	res := three_sum(nums, n)
	// fmt.Println(nums)
	for l := 0; l < len(res); l++ {
		i, j, k := res[l][0], res[l][1], res[l][2]
		fmt.Println(i, " : ", nums[i])
		fmt.Println(j, " : ", nums[j])
		fmt.Println(k, " : ", nums[k])
	}
}
