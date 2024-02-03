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

type Edge struct {
	U  int
	V  int
	Wt int
}

func bellman_ford(g *Graph[int], source int) []int {
	distance := make([]int, g.V)
	for i := 0; i < g.V; i++ {
		distance[i] = -1
	}
	distance[source] = 0
	edges := []Edge{}
	for i := 0; i < g.V; i++ {
		for j := 0; j < len(g.Adj()[i]); j++ {
			edges = append(edges, Edge{
				U:  i,
				V:  g.Adj()[i][j].Id,
				Wt: g.Adj()[i][j].Wt,
			})
		}
	}

	for i := 1; i < g.V; i++ {
		for j := 0; j < len(edges); j++ {
			edge := edges[j]
			u := edge.U
			v := edge.V
			wt := edge.Wt
			if distance[u] == -1 {
				continue
			}
			if distance[v] == -1 || distance[v] > distance[u]+wt {
				distance[v] = distance[u] + wt
			}
		}
	}

	for i := 0; i < len(edges); i++ {
		edge := edges[i]
		u := edge.U
		v := edge.V
		wt := edge.Wt
		if distance[u] == -1 {
			continue
		}
		if distance[v] == -1 || distance[v] > distance[u]+wt {
			return []int{-1} // graph contain negtive cycle
		}
	}

	return distance
}

/*
// With negtive cycle :
V = 6, E =7
0 1 5
1 3 2
3 5 2
5 4 -3
4 2 1
1 2 1
4 3 -1

// With negtive edge but no negtive cycle
V = 6, E =7
0 1 5
1 3 2
3 5 2
5 4 -3
4 2 1
1 2 1
4 3 4

// with no negtive edges
V = 6, E =7
0 1 5
1 3 2
3 5 2
5 4 3
4 2 1
1 2 1
4 3 4

*/

func main() {
	v, e := readInt(), readInt()
	g := NewGraph[int](v, e)
	for i := 0; i < e; i++ {
		a, b, wt := readInt(), readInt(), readInt()
		g.AddDirectedEdge(a, b, wt)
	}

	fmt.Println(bellman_ford(g, 0))
}
