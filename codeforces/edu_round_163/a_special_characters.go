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

func get_output(n int) {
	if n&1 == 1 {
		// odd
		fmt.Println("NO")
	} else {
		// even
		str := strings.Builder{}
		for i := 0; i < n/2; i++ {
			if i&1 == 1 {
				str.WriteString("AA")
			} else {
				str.WriteString("BB")
			}
		}
		fmt.Println("YES")
		fmt.Println(str.String())
		return
	}
}

func main() {
	t := readInt()
	for t > 0 {
		n := readInt()
		get_output(n)
		t--
	}
}
