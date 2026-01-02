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

func swap(arr []byte, i, j int) {
	if i != j {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func permuatation_helper(s []byte, it int, n int) {
	if it == n {
		fmt.Println(string(s))
		return
	}
	for i := it; i < n; i++ {
		swap(s, i, it)
		permuatation_helper(s, it+1, n)
		swap(s, i, it)
	}
}

func permutations_of_string(s string) {
	ss := []byte(s)
	fmt.Println("OUTPUT : ")
	permuatation_helper(ss, 0, len(s))
}

func main() {
	s := readString()

	permutations_of_string(s)
}
