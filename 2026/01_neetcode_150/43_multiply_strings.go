package main

import (
	"fmt"
	"log"
	"strconv"
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

func multiply_strings(s1 string, s2 string) string {
	res := 0
	power_a := 1
	for i := len(s2) - 1; i >= 0; i-- {
		curr := 0
		s2Val := int(s2[i] - '0')
		power_b := 1
		for j := len(s1) - 1; j >= 0; j-- {
			s1Val := int(s1[j] - '0')
			curr = curr + s2Val*s1Val*power_b
			power_b = power_b * 10
		}
		res = res + curr*power_a
	}

	return strconv.Itoa(res)
}

func main() {
	s1 := readString()
	s2 := readString()

	fmt.Printf("multiply_strings(s1, s2): %v\n", multiply_strings(s1, s2))
}
