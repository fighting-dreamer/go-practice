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

func assign_cookies(students []int, n int, cookies []int, m int) int {
	i := 0
	j := 0
	res := 0

	for ; j < m; j++ {
		for ; i < n; i++ {
			if students[i] <= cookies[j] {
				i++
				res++
				break
			}
		}
	}

	return res
}

func main() {
	n := readInt()
	m := readInt()
	students := make([]int, n)
	cookies := make([]int, m)
	for i := 0; i < n; i++ {
		students[i] = readInt()
	}
	for i := 0; i < m; i++ {
		cookies[i] = readInt()
	}

	res := assign_cookies(students, n, cookies, m)
	fmt.Printf("res: %v\n", res)
}
