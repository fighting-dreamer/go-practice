package main

import (
	"fmt"
	"log"

	"github.com/liyue201/gostl/ds/queue"
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

func bfs_01(g *Graph[int], src int) []int {
	q := queue.New[int]()
	vis := make([]bool, g.V)
	distance := make([]int, g.V)

	q.Push(src)
	vis[src] = true
	distance[src] = 0

	for !q.Empty() {
		u := q.Front()

		q.Pop()

		for i := 0; i < len(g.Adj()[u]); i++ {
			v := g.Adj()[u][i]
			if !vis[v.Id] {
				q.Push(v.Id)
				vis[v.Id] = true
				distance[v.Id] = distance[u] + 1
			}
		}
	}

	return distance
}

func main() {
	v, e := readInt(), readInt()
	g := NewGraph[int](v, e)
	for i := 0; i < e; i++ {
		u, v, wt := readInt(), readInt(), 1
		g.AddDirectedEdge(u, v, wt)
	}

	fmt.Println(bfs_01(g, 0))
}
