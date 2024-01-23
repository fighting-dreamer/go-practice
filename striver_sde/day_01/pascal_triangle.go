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

func print_nth_pascal_triangle_line(n int) []int {
	if n == 0 {
		return []int{}
	}
	if n == 1 {
		return []int{1}
	}
	if n == 2 {
		return []int{1, 1}
	}
	pascal, tmp := []int{}, []int{1, 1}
	for i := 3; i <= n; i++ {
		pascal = make([]int, i)
		pascal[0] = 1
		for j := 1; j < i-1; j++ {
			pascal[j] = tmp[j-1] + tmp[j]
		}
		pascal[i-1] = 1
		tmp = pascal
	}

	return pascal
}

func main() {
	n := readInt()
	fmt.Println(print_nth_pascal_triangle_line(n))
}
