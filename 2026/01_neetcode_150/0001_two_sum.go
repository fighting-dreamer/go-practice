package main

import (
	"fmt"
	"log"
	"slices"
)

type Pair struct {
	a int
	b int
}

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

func two_sum(arr []int, target int) []Pair {
	slices.Sort(arr)
	n := len(arr)
	l, r := 0, n-1
	res := []Pair{}

	for l < r {
		curr := arr[l] + arr[r]
		if curr == target {
			// handle duplicates and blocks
			if arr[l] == arr[r] {
				for ll := l; ll < r; ll++ {
					for rr := ll + 1; rr < r; rr++ {
						res = append(res, Pair{
							a: ll,
							b: rr,
						})
					}
				}
				break // exchaustive list of pairs.
			} else {
				tempL := l
				tempR := r

				for tempL < r && arr[tempL] == arr[tempL+1] {
					tempL++
				}
				for tempR > l && arr[tempR] == arr[tempR-1] {
					tempR--
				}

				for i := l; i <= tempL; i++ {
					for j := tempR; j <= r; j++ {
						res = append(res, Pair{
							a: i,
							b: j,
						})
					}
				}

				l = tempL + 1
				r = tempR - 1
			}
			continue
		}

		if curr < target {
			l++
		} else {
			r--
		}
	}

	fmt.Println(res)
	return res
}

func main() {
	n := readInt()
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = readInt()
	}

	target := readInt()
	two_sum(arr, target)
}
