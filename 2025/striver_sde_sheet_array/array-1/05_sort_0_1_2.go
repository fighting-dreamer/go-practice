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

func print_array(arr []int, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", arr[i])
	}
	fmt.Println()
}

func sort_a_b(arr []int, start int, end int, a int) (int, int) {
	l := start
	r := end
	for l < r {
		for l < r && arr[l] == a {
			l++
		}
		for l < r && arr[r] != a {
			r--
		}
		if l == r {
			break
		}
		arr[l], arr[r] = arr[r], arr[l]
		l++
		r--
	}

	return l, r
}

func sort_0_1_2(arr []int, n int) {
	// sort all zero vs all-non-zero
	start, _ := sort_a_b(arr, 0, n-1, 0)
	// start points to end of zeros always
	// sort between all 1s and non-1s
	sort_a_b(arr, start, n-1, 1)
}

func main() {
	n := readInt()
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = readInt()
	}

	print_array(arr, n)
	sort_0_1_2(arr, n)
	print_array(arr, n)
}
