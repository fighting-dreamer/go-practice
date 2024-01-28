package main

import (
	"fmt"
	"log"

	"github.com/liyue201/gostl/ds/pair"
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

// only checks to right or down.
func wordExists(charMatrix []string, word string, currIndex, limit, row, col, rowCount, colCount int) bool {
	if row >= rowCount || col >= colCount {
		return false
	}
	if currIndex+1 == limit {
		return true
	}
	// charMatrix[row][col] == word[currIndex]
	var flag bool
	if col+1 < colCount && charMatrix[row][col+1] == word[currIndex+1] {
		flag = wordExists(charMatrix, word, currIndex+1, limit, row, col+1, rowCount, colCount)
	}
	if flag == true {
		return flag
	}
	if row+1 < rowCount && charMatrix[row+1][col] == word[currIndex+1] {
		flag = wordExists(charMatrix, word, currIndex+1, limit, row+1, col, rowCount, colCount)
	}
	return flag
}

// only checks right or down
func WordExists(charMatrix []string, word string) bool {
	rowCount := len(charMatrix)
	colCount := len(charMatrix[0])
	l := len(word)

	for i := 0; i < rowCount; i++ {
		for j := 0; j < colCount; j++ {
			if charMatrix[i][j] == word[0] && wordExists(charMatrix, word, 0, l, i, j, rowCount, colCount) {
				return true
			}
		}
	}
	return false
}

var dir = []pair.Pair{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

func FrontInt(p pair.Pair) int {
	v, _ := p.Front.(int)
	return v
}

func BackInt(p pair.Pair) int {
	v, _ := p.Back.(int)
	return v
}

func wordExistsAllDir(charMatrix []string, word string, currIndex, limit, row, col, rowCount, colCount int, vis [][]int) bool {
	// fmt.Println(string(byte(charMatrix[row][col])), " ", row, " ", col, " ", currIndex)
	if currIndex == limit-1 {
		return true
	}

	// we visit only if not visited yet.
	// if vis[row][col] == 1 {
	// 	return false
	// }

	vis[row][col] = 1
	flag := false
	for i := 0; i < len(dir); i++ {
		currDir := dir[i]
		rowNext := row + FrontInt(currDir)
		colNext := col + BackInt(currDir)

		if rowNext < 0 || rowNext >= rowCount ||
			colNext < 0 || colNext >= colCount {
			continue
		}

		if vis[rowNext][colNext] == 1 {
			// already visited
			continue
		}

		if charMatrix[rowNext][colNext] == word[currIndex+1] {
			flag = wordExistsAllDir(charMatrix, word, currIndex+1, limit, rowNext, colNext, rowCount, colCount, vis)
			if flag == true {
				return flag
			}
		}
	}
	vis[row][col] = 0
	return flag
}

func WordExistsAllDir(charMatrix []string, word string) bool {
	rowCount := len(charMatrix)
	colCount := len(charMatrix[0])
	limit := len(word)
	vis := make([][]int, rowCount)
	for i := 0; i < rowCount; i++ {
		vis[i] = make([]int, colCount)
	}
	for i := 0; i < rowCount; i++ {
		for j := 0; j < colCount; j++ {
			if charMatrix[i][j] == word[0] {
				flag := wordExistsAllDir(charMatrix, word, 0, limit, i, j, rowCount, colCount, vis)
				if flag == true {
					return true
				}
			}
		}
	}
	return false
}

func main() {
	n, m := readInt(), readInt()
	charMatrix := make([]string, n)

	for i := 0; i < n; i++ {
		charMatrix[i] = readString()
	}

	for i := 0; i < m; i++ {
		word := readString()
		fmt.Println(word, " : ", WordExistsAllDir(charMatrix, word))
	}
}
