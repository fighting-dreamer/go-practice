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

func isValidString(s string) bool {
	mp := map[string]bool{
		"1":  true,
		"2":  true,
		"3":  true,
		"4":  true,
		"5":  true,
		"6":  true,
		"7":  true,
		"8":  true,
		"9":  true,
		"10": true,
		"11": true,
		"12": true,
		"13": true,
		"14": true,
		"15": true,
		"16": true,
		"17": true,
		"18": true,
		"19": true,
		"20": true,
		"21": true,
		"22": true,
		"23": true,
		"24": true,
		"25": true,
		"26": true,
	}
	_, exists := mp[s]
	return exists
}

func decodeStringWays(curr int, s string) int {
	if curr == len(s) {
		return 1
	}

	if s[curr] == '0' {
		return 0
	}

	one_split := 0
	two_split := 0
	if isValidString(s[curr : curr+1]) {
		one_split = decodeStringWays(curr+1, s)
	}

	if curr+2 <= len(s) && isValidString(s[curr:curr+2]) {
		two_split = decodeStringWays(curr+2, s)
	}
	return one_split + two_split
}

func DecodeStringWays(s string) int {
	return decodeStringWays(0, s)
}

func main() {
	s := readString()

	fmt.Println(DecodeStringWays(s))
}
