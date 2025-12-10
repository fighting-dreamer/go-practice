package main

import (
	"cmp"
	"fmt"
	"log"
	"slices"
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

type Pair[T, U any] struct {
	First  T
	Second U
}

// schedule as `pair` with start and end
func N_MeetingRoom(arr []Pair[int, int], n int) int {
	slices.SortFunc(arr, func(a Pair[int, int], b Pair[int, int]) int {
		return cmp.Compare(a.Second, b.Second)
	})
	fmt.Println(arr)
	res := 1 // including the first pair(schedule)
	curr := arr[0]
	resultingMeeting := []Pair[int, int]{curr}
	for i := 1; i < n; i++ {
		if curr.Second < arr[i].First {
			res++
			resultingMeeting = append(resultingMeeting, arr[i])
			curr = arr[i]
		}
	}
	fmt.Println("Res:", resultingMeeting)
	return res
}

func main() {
	n := readInt()
	arr := make([]Pair[int, int], n) // schedule as pair with start and end
	for i := 0; i < n; i++ {
		arr[i] = Pair[int, int]{
			First:  readInt(),
			Second: readInt(),
		}
	}

	res := N_MeetingRoom(arr, n)
	fmt.Printf("res: %v\n", res)
}
