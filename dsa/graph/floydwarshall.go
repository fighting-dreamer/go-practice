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

func floydwarshall(g *Graph[int]) [][]int {
	inf := 99999
	distance := make([][]int, g.V)
	for i := 0; i < g.V; i++ {
		distance[i] = make([]int, g.V)
	}

	for i := 0; i < g.V; i++ {
		for j := 0; j < g.V; j++ {
			if i == j {
				distance[i][j] = 0
				continue
			}
			distance[i][j] = inf
		}
	}
	edges := []Edge{}
	for i := 0; i < g.V; i++ {
		for j := 0; j < len(g.Adj()[i]); j++ {
			edge := Edge{
				U:  i,
				V:  g.Adj()[i][j].Id,
				Wt: g.Adj()[i][j].Wt,
			}
			edges = append(edges, edge)
			distance[edge.U][edge.V] = edge.Wt
		}
	}
	fmt.Println(distance)
	for k := 0; k < g.V; k++ {
		for i := 0; i < g.V; i++ {
			for j := 0; j < g.V; j++ {
				if distance[i][k] != inf && distance[k][j] != inf &&
					(distance[i][j] == inf || distance[i][j] > distance[i][k]+distance[k][j]) {
					distance[i][j] = distance[i][k] + distance[k][j]
				}
			}
		}
	}

	return distance
}

func main() {
	v, e := readInt(), readInt()
	g := NewGraph[int](v, e)
	for i := 0; i < e; i++ {
		a, b, wt := readInt(), readInt(), readInt()
		g.AddDirectedEdge(a, b, wt)
	}

	fmt.Println(floydwarshall(g))
}
