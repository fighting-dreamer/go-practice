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

func gcd_strings(str1, str2 string) string {
	// fmt.Println(str1, str2)
	if str1 == str2 {
		return str1
	}
	i, j := 0, 0
	l1, l2 := len(str1), len(str2)
	for i < l1 {
		if j == l2 {
			j = 0
		}
		if str2[j] == str1[i] {
			i++
			j++
		} else {
			break
		}
	}
	if i != l1 {
		// we got an index i that is not matching with jth relevant index.
		return "" // example (ABCDBC, ABC)
	}
	fmt.Println(i, j, l1, l2)
	if i == l1 && j == l2 {
		return str2 // example (ABCABC, ABC)
	}
	// else we have some gcd to be checked more
	return gcd_strings(str2, str1[l1-j:])
}

func main() {
	str1 := readString()
	str2 := readString()

	fmt.Println(gcd_strings(str1, str2))
}
