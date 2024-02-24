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

func merge_alternately(str1, str2 string) string {
	l1, l2 := len(str1), len(str2)
	b := make([]byte, l1+l2)

	i, j, k := 0, 0, 0
	for i < l1 && j < l2 {
		if ((i + j) & 1) == 0 {
			b[k] = str1[i]
			i++
		} else {
			b[k] = str2[j]
			j++
		}
		k++
	}

	for i < l1 {
		b[k] = str1[i]
		i++
		k++
	}

	for j < l2 {
		b[k] = str2[j]
		j++
		k++
	}

	return string(b)
}

func main() {
	str1, str2 := readString(), readString()

	fmt.Println(merge_alternately(str1, str2))
}
