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

func zip(set1 []string, set2 []string) []string {
	fmt.Println(set1, set2)
	res := []string{}
	if len(set1) == 0 {
		return set2
	}
	if len(set2) == 0 {
		return set1
	}
	for i := 0; i < len(set1); i++ {
		for j := 0; j < len(set2); j++ {
			res = append(res, set1[i]+set2[j])
		}
	}
	return res
}

func letter_combinations(s string) []string {
	mapping := map[byte]string{
		1: "",
		2: "abc",
		3: "def",
		4: "ghi",
		5: "jkl",
		6: "mno",
		7: "pqrs",
		8: "tuv",
		9: "wxyz",
	}
	res := []string{}
	for i := 0; i < len(s); i++ {
		if s[i] == '1' || s[i] == '0' {
			continue
		}
		res = zip(res, strings.Split(mapping[s[i]-'0'], ""))
	}

	return res
}

func main() {
	s := readString()

	fmt.Printf("letter_combinations(s): %v\n", letter_combinations(s))
}
