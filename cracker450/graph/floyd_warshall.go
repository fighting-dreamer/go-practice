package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"strings"
)

type Node[T any] struct {
	Id int
	Wt T
}

type Graph[T any] struct {
	V   int
	E   int
	adj [][]*Node[T]
}

func (g *Graph[T]) Adj() [][]*Node[T] {
	return g.adj
}
func NewGraph[T any](v int, e int) *Graph[T] {
	return &Graph[T]{
		V:   v,
		E:   e,
		adj: make([][]*Node[T], v),
	}
}

func (g *Graph[T]) AddEdge(u, v int, wt T) {
	if u < 0 || u >= g.V || v < 0 || v >= g.V {
		return
	}
	g.adj[u] = append(g.adj[u], &Node[T]{
		Id: v,
		Wt: wt,
	})
	g.adj[v] = append(g.adj[v], &Node[T]{
		Id: u,
		Wt: wt,
	})
}
func (g *Graph[T]) AddDirectedEdge(u, v int, wt T) {
	if u < 0 || u >= g.V || v < 0 || v >= g.V {
		return
	}
	g.adj[u] = append(g.adj[u], &Node[T]{Id: v, Wt: wt})
}
func (g *Graph[T]) printGraphAdjList() {
	for i := 0; i < g.V; i++ {
		for j := 0; j < len(g.adj[i]); j++ {
			n := g.adj[i][j]
			fmt.Println(i, n.Id, n.Wt)
		}
	}
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

type IntMatrix [][]int

func (im IntMatrix) String() string {
	d := [][]int(im)
	var sb strings.Builder
	for i := 0; i < len(d); i++ {
		for j := 0; j < len(d[i]); j++ {
			sb.WriteString(fmt.Sprintf("%d ", d[i][j]))
		}
		sb.WriteString(fmt.Sprintln())
	}
	return sb.String()
}

func test_IntMatrix_string() {
	res := make([][]int, 5)
	for i := 0; i < 5; i++ {
		res[i] = ([]int{1, 2, 3, 4, 5})[:rand.Intn(5)]
	}
	fmt.Print(IntMatrix(res))

}

func floyd_warshall(g *Graph[int]) [][]int {
	distance := make([][]int, g.V)
	for i := 0; i < g.V; i++ {
		distance[i] = make([]int, g.V)
		// distance[i][i] = 0 // default zero value of int is 0
	}
	for i := 0; i < g.V; i++ {
		for j := 0; j < g.V; j++ {
			if i == j {
				continue
			}
			distance[i][j] = math.MaxInt
		}
	}
	for i := 0; i < g.V; i++ {
		for j := 0; j < len(g.Adj()[i]); j++ {
			v, wt := g.Adj()[i][j].Id, g.Adj()[i][j].Wt
			distance[i][v] = wt
		}
	}
	fmt.Println("--------------------")
	fmt.Println(IntMatrix(distance))
	fmt.Println("--------------------")

	for i := 0; i < g.V; i++ {
		for j := 0; j < g.V; j++ {
			for k := 0; k < g.V; k++ {
				if distance[i][k] == math.MaxInt || distance[k][j] == math.MaxInt {
					continue
				}
				if distance[i][j] == math.MaxInt || distance[i][k]+distance[k][j] < distance[i][j] {
					distance[i][j] = distance[i][k] + distance[k][j]
				}
			}
		}
	}

	return distance
}

// input : 3 6 0 1 1 0 2 43 1 0 1 1 2 6 2 0 -1 2 1 -1
// output : 0 1 7 \n1 0 6 \n-1 -1 0
func main() {
	// test_IntMatrix_string()

	v, e := readInt(), readInt()
	g := NewGraph[int](v, e)

	for i := 0; i < g.E; i++ {
		u, v, wt := readInt(), readInt(), readInt()
		g.AddDirectedEdge(u, v, wt)
	}

	fmt.Println(IntMatrix(floyd_warshall(g)))

}
