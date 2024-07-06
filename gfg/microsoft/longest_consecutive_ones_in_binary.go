package main

import (
	"fmt"
	"log"
	"slices"
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

func binary_representation(n int) []byte {
	t := n
	binaryRep := strings.Builder{}
	for t > 0 {
		binaryRep.WriteString(fmt.Sprintf("%d", t&1))
		t >>= 1
	}
	rep := []byte(binaryRep.String())
	slices.Reverse([]byte(rep))

	return rep
}

func find_longest_consecutive_ones(n int) int {
	binaryRep := binary_representation(n)
	cnt, res, flag := 0, 0, 0
	for i := 0; i < len(binaryRep); i++ {
		if binaryRep[i] == '1' {
			if i > 0 && binaryRep[i-1] == '1' {
				cnt++
			}
			flag = 1
		} else {
			cnt = 0
		}
		res = max(res, cnt+flag)
	}

	return res
}

func main() {
	n := readInt()

	fmt.Println(find_longest_consecutive_ones(n))
}
