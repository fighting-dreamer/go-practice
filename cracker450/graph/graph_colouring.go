package main

import (
	"fmt"
	"log"
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

func isSafe(currNode int, g *Graph[int], c int, colours []int) bool {
	for i := 0; i < len(g.Adj()[currNode]); i++ {
		v := g.Adj()[currNode][i].Id
		if colours[v] == c {
			return false
		}
	}
	return true
}

func graph_colouring_helper(currNode int, g *Graph[int], color []int, m int) bool {
	if currNode == g.V {
		return true
	}

	for i := 1; i <= m; i++ { // default zero is like we haven;t added a colour yet.
		if isSafe(currNode, g, i, color) {
			color[currNode] = i

			return graph_colouring_helper(currNode+1, g, color, m)
		}
	}

	return false
}

func graph_colouring(g *Graph[int], m int) bool {
	colourOfNode := make([]int, g.V)
	if graph_colouring_helper(0, g, colourOfNode, m) == true {
		return true
	}
	return false
}

func main() {
	v, e := readInt(), readInt()
	g := NewGraph[int](v, e)
	for i := 0; i < g.E; i++ {
		u, v, wt := readInt(), readInt(), 1
		g.AddEdge(u, v, wt)
	}

	m := readInt()
	fmt.Println(graph_colouring(g, m))
}
