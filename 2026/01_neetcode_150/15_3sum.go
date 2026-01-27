package main

import (
	"fmt"
	"log"
	"slices"
	"strconv"
	"strings"
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

// []int -> string
func sliceToString(nums []int) string {
	res := make([]string, len(nums))
	for i, v := range nums {
		res[i] = strconv.Itoa(v)
	}
	return strings.Join(res, ",")
}

// string -> []int
func stringToSlice(s string) []int {
	if s == "" {
		return nil
	}
	fields := strings.Split(s, ",")
	res := make([]int, len(fields))
	for i, v := range fields {
		res[i], _ = strconv.Atoi(v)
	}
	return res
}

func three_values_sum_to_zero(arr []int, n int) [][]int {
	res := [][]int{}
	slices.Sort(arr)
	for i := 0; i < n-2; i++ {
		target := -arr[i]
		l := i + 1
		r := n - 1
		for l < r {
			if arr[l]+arr[r] == target {
				// res = append(res, []int{i, l, r})
				if arr[l] == arr[r] {
					for ii := l; ii < r; ii++ {
						for jj := ii + 1; jj <= r; jj++ {
							res = append(res, []int{i, ii, jj})
						}
					}
					break // exhaustive for this target and starting l and r
				} else {
					tempL := l
					tempR := r
					for tempL < r && arr[tempL] == arr[tempL+1] {
						tempL++
					}
					for tempR > l && arr[tempR] == arr[tempR-1] {
						tempR--
					}

					for ii := l; ii <= tempL; ii++ {
						for jj := tempR; jj <= r; jj++ {
							res = append(res, []int{i, ii, jj})
						}
					}
					l = tempL + 1
					r = tempR - 1
				}
				continue
			}
			if arr[l]+arr[r] < target {
				l++
			} else {
				r--
			}
		}
	}

	mp := map[string]bool{}
	for i := 0; i < len(res); i++ {
		mp[sliceToString([]int{arr[res[i][0]], arr[res[i][1]], arr[res[i][2]]})] = true
	}
	res = [][]int{}
	for k, _ := range mp {
		res = append(res, stringToSlice(k))
	}

	return res
}

func main() {
	n := readInt()
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = readInt()
	}
	for i := 0; i < n; i++ {
		fmt.Println(i, " => ", arr[i])
	}
	fmt.Printf("three_values_sum_to_zero(arr, n): %v\n", three_values_sum_to_zero(arr, n))
}
