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

type Triplet struct {
	a int
	b int
	c int
}

func three_sum(arr []int, n int, target int) []Triplet {
	res := []Triplet{}

	for i := 0; i < n; i++ {
		newTarget := target - arr[i]
		l := i + 1
		r := n - 1
		for l < r {
			sum := arr[l] + arr[r]
			if sum == newTarget {
				res = append(res, Triplet{
					a: i,
					b: l,
					c: r,
				})

				// if you donr want same number taken, you can increment say l till consecutive array element are same
				// and do so for right too.
				// if you want all different indices then
				// for each consecutive left side same element use the same r and do same for right side
				var k int
				lfinal := l
				rfinal := r
				for k = l + 1; arr[k] == arr[l] && k < r; k++ {
				}
				lfinal = k
				fmt.Println("L : lfinal", l, lfinal)
				for k = r - 1; arr[k] == arr[r] && k > l; k-- {
				}
				rfinal = k
				fmt.Println("R : rfinal", r, rfinal)

				for ii := l; ii < lfinal; ii++ {
					for jj := r; jj > rfinal; jj-- {
						if (ii == l && jj == r) || ii >= jj {
							continue // this is already added above
						}
						res = append(res, Triplet{
							a: i,
							b: ii,
							c: jj,
						})
					}
				}

				l = lfinal
				r = rfinal
			} else {
				if sum > newTarget {
					r--
				} else {
					l++
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

	res := three_sum(arr, n, target)
	for i := 0; i < len(res); i++ {
		fmt.Println(res[i].a, res[i].b, res[i].c)
	}
}
