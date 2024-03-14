package main

import (
	"fmt"
	"log"

	"github.com/liyue201/gostl/ds/priorityqueue"
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
	U  int
	V  int
	Wt int
}

func find(par []int, u int) int {
	if par[u] == u {
		return u
	}
	par[u] = find(par, par[u])
	return par[u]
}

func union(par []int, u int, v int) {
	up := find(par, u)
	vp := find(par, v)
	if up != vp {
		par[up] = vp
	}
}

go_min_heap
priorityqueue.New[T](comparator.Reverse(func(a,b T)int {
	return //ascending
}))


func kruskal(g *Graph[int]) (int, []Edge) {
	par := make([]int, g.V)
	for i := 0; i < g.V; i++ {
		par[i] = i
	}
	minHeap := priorityqueue.New[Edge](func(a, b Edge) int {
		if a.Wt == b.Wt {
			return 0
		}
		if a.Wt < b.Wt {
			return -1
		}
		return 1
	})
	for i := 0; i < g.V; i++ {
		for j := 0; j < len(g.Adj()[i]); j++ {
			node := g.Adj()[i][j]
			minHeap.Push(Edge{
				i, node.Id, node.Wt,
			})
		}
	}
	resWt := 0
	resEdges := []Edge{}
	for !minHeap.Empty() {
		edge := minHeap.Pop()
		uf := find(par, edge.U)
		vf := find(par, edge.V)
		if uf != vf {
			union(par, uf, vf)
			resWt += edge.Wt
			resEdges = append(resEdges, edge)
		}
	}

	return resWt, resEdges
}

/*
V = 9, E = 14
u v wt
0 1 4
0 7 8
1 7 11
1 2 8
7 8 7
7 6 1
2 8 2
8 6 6
2 3 7
2 5 4
6 5 2
3 5 14
3 4 9
5 4 10

result :
37 [{6 7 1} {2 8 2} {5 6 2} {2 5 4} {0 1 4} {2 3 7} {2 1 8} {3 4 9}]

// there is other solution coz 0-7 edge and 1-2 edge are of same weight of 8 and connect graph to make it a MST tree

37 [{6 7 1} {2 8 2} {5 6 2} {2 5 4} {0 1 4} {2 3 7} {0 7 8} {3 4 9}]
*/
func main() {
	v, e := readInt(), readInt()
	g := NewGraph[int](v, e)
	for i := 0; i < e; i++ {
		a, b, wt := readInt(), readInt(), readInt()
		g.AddEdge(a, b, wt)
	}

	fmt.Println(kruskal(g))
}
