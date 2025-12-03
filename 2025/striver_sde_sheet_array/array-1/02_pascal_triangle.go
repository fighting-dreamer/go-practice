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

func print_pascal_triangle(n int) {
	arr := []int{1}
	fmt.Println(arr)
	for i := 1; i <= n; i++ {
		res := []int{1}
		for j := 1; j < i; j++ {
			temp := arr[j-1] + arr[j]
			res = append(res, temp)
		}
		res = append(res, 1)
		fmt.Println(res)
		arr = res
	}
}

func main() {
	n := readInt()
	print_pascal_triangle(n)
}
