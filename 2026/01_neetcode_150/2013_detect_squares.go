package main

import (
	"fmt"
	"log"
)

type Coord struct {
	X int
	Y int
}

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

func detect_squares(n int) {
	coordSet := map[Coord]bool{}
	xSet := map[int][]Coord{}
	ySet := map[int][]Coord{}

	for i := 0; i < n; i++ {
		t, x, y := readString(), readInt(), readInt()
		coord := Coord{
			X: x,
			Y: y,
		}
		if t == "add" {
			xSet[x] = append(xSet[x], coord)
			ySet[y] = append(ySet[y], coord)
			coordSet[coord] = true
		} else if t == "count" {
			count := 0
			xCoords := xSet[x]
			yCoords := ySet[y]
			for i := 0; i < len(xCoords); i++ {
				for j := 0; j < len(yCoords); j++ {
					reqdCoord := Coord{
						Y: xCoords[i].Y,
						X: yCoords[j].X,
					}
					if coordSet[reqdCoord] {
						count++
					}
				}
			}

			fmt.Println(count)
		}
	}
}

func main() {
	n := readInt()

	detect_squares(n)
}
