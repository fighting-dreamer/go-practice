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

func decode_ways_helper(s string, l int, curr int, path []byte, res *[]string) {
	if curr >= l {
		*res = append(*res, string(path))
		return
	}

	singleDigit := s[curr] - '0'
	if singleDigit != 0 {
		decode_ways_helper(s, l, curr+1, append(path, s[curr]-'1'+'A'), res)
	}
	if curr+2 <= l {
		doubleDigits, _ := strconv.Atoi(s[curr : curr+2])
		if doubleDigits >= 10 && doubleDigits <= 26 {
			decode_ways_helper(s, l, curr+2, append(path, byte(doubleDigits+'A'-1)), res)
		}
	}

	return
}

func decode_ways(s string) []string {
	l := len(s)
	res := []string{}
	decode_ways_helper(s, l, 0, []byte{}, &res)
	return res
}

func main() {
	s := readString()
	fmt.Printf("decode_ways(s): %v\n", decode_ways(s))
}
