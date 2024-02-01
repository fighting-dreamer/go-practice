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

func largest_subarray(nums []int, n, t int) []int {
	fmt.Println("largest subarray : ", nums, n, t)
	resSize := 0
	start, end := 0, 0
	tempSum := 0
	for i, j := 0, 0; i < n && j < n; {
		tempSum += nums[j]
		if tempSum < t {
			j++
		} else if tempSum == t {
			fmt.Println(i, j, tempSum, nums[i:j+1])
			if resSize < j-i+1 {
				start = i
				end = j
				resSize = j - i + 1
			}

			// tempSum -= nums[i]

			// i++
			j++
		} else {
			for tempSum > t {
				tempSum -= nums[i]
				i++
			}

			if tempSum == t {
				fmt.Println("-> ", i, j, tempSum, nums[i:j+1])

				if resSize < j-i+1 {
					start = i
					end = j
					resSize = j - i + 1
				}
			}

			j++
		}
	}

	return nums[start : end+1]
}

// not work for negtive and postive array combined.
// works with postive array only.
func main() {
	n, t := readInt(), readInt()
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = readInt()
	}

	fmt.Println(largest_subarray(nums, n, t))
}
