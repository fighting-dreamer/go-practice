package main

import (
	"fmt"
	"log"
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

func print_pattern(n, m int) {
	sb := strings.Builder{}

	for i := 0; i < m; i++ {
		sb.WriteString("*")
	}
	str := sb.String()
	for i := 0; i < n; i++ {
		fmt.Println(str)
	}
}

func main() {
	n, m := readInt(), readInt()
	print_pattern(n, m)
}
