package main

import (
	"fmt"
	"log"
	"strconv"
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

func encode_string(strs []string, n int) string {
	sb := strings.Builder{}
	for i := 0; i < n; i++ {
		sb.Write([]byte(strconv.Itoa(len(strs[i]))))
		sb.WriteRune('|')
		sb.WriteString(strs[i])
	}
	return sb.String()
}

func decode_string(str string) []string {
	dstrs := []string{}

	for i := 0; i < len(str); {
		start := i
		end := i
		for str[end] != '|' {
			end++
		}
		len, err := strconv.Atoi(str[start:end])
		if err != nil {
			fmt.Println(str, " ", start, " ", end)
			log.Fatal("could not read length")
			return []string{}
		}
		dstrs = append(dstrs, str[end+1:len+end+1])
		i = end + len + 1
	}

	return dstrs
}

// input : ["qwerty", "121452", "abcdef"]
// encoded string : "something....", here it is "length of the string word, '|' as a seperator, string itself... for all words"
// decode from encoded string : ["qwerty", "121452", "abcdef"]

func main() {
	n := readInt()
	strs := []string{}
	for i := 0; i < n; i++ {
		strs = append(strs, readString())
	}
	estr := encode_string(strs, n)
	fmt.Println(estr)

	dstrs := decode_string(estr)

	for i := 0; i < n; i++ {
		if dstrs[i] != strs[i] {
			fmt.Println("Fault in decoding", strs[i], dstrs[i])
		}
	}

	fmt.Println(dstrs)
}
