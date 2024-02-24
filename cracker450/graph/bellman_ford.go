package main

import (
	"fmt"
	"log"
	"math"
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

type Edge struct {
	u  int
	v  int
	wt int
}

func bellman_ford(g *Graph[int], src int) []int {
	distance := make([]int, g.V)
	for i := 0; i < g.V; i++ {
		distance[i] = math.MinInt
	}
	distance[src] = 0

	edges := make([]Edge, g.E)
	k := 0
	for i := 0; i < g.V; i++ {
		for j := 0; j < len(g.Adj()[i]); j++ {
			v, wt := g.Adj()[i][j].Id, g.Adj()[i][j].Wt
			edges[k] = Edge{
				u:  i,
				v:  v,
				wt: wt,
			}
			k++
		}
	}

	// relax edges g.V - 1 times
	for i := 1; i < g.V; i++ {
		for j := 0; j < g.E; j++ {
			if distance[edges[j].u] == math.MinInt {
				// we have not reached 'u'
				continue
			}
			if distance[edges[j].v] == math.MinInt || distance[edges[j].v] < distance[edges[j].u]+edges[j].wt {
				distance[edges[j].v] = distance[edges[j].u] + edges[j].wt
			}
		}
	}

	// check if there is a negtive cycle
	for j := 0; j < g.E; j++ {
		if distance[edges[j].v] < distance[edges[j].u]+edges[j].wt {
			// fmt.Println(distance)
			// fmt.Println(edges[j])
			// distance[edges[j].v] = distance[edges[j].u] + edges[j].wt
			return []int{-1} // negtive cycle
		}
	}
	return distance
}

func main() {
	v, e := readInt(), readInt()
	g := NewGraph[int](v, e)

	for i := 0; i < e; i++ {
		u, v, wt := readInt(), readInt(), readInt()
		g.AddDirectedEdge(u, v, wt)
	}

	src := readInt()
	fmt.Println(bellman_ford(g, src))

}
