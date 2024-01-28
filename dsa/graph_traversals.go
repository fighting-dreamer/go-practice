package main

import "fmt"

type Node[T any] struct {
	Id int
	Wt T
}

type Graph[T any] struct {
	V   int
	E   int
	adj [][]*Node[T]
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
	g.adj[u] = append(g.adj[u], &Node[T]{
		Id: v,
		Wt: wt,
	})
}

func (g *Graph[T]) printGraphAdjList() {
	for i := 0; i < g.V; i++ {
		for j := 0; j < len(g.adj[i]); j++ {
			n := g.adj[i][j]
			fmt.Println(i, n.Id, n.Wt)
		}
	}
}

func (g *Graph[T]) BFSTraversal() {
	var q []int
	vis := make([]bool, g.V)

	q = append(q, 0)
	vis[0] = true
	for len(q) > 0 {
		u := q[0]
		q = q[1:]

		fmt.Println(u)

		for i := 0; i < len(g.adj[u]); i++ {
			v := g.adj[u][i].Id
			if vis[v] == false {
				q = append(q, v)
				vis[v] = true
			}
		}
	}
}

func dfsHelper[T any](vis []int, u int, g *Graph[T]) {
	vis[u] = 1

	fmt.Println(u)

	for i := 0; i < len(g.adj[u]); i++ {
		v := g.adj[u][i].Id
		if vis[v] == 0 {
			dfsHelper(vis, v, g)
		}
	}
}

func (g *Graph[T]) DFSTraversal() {
	vis := make([]int, g.V)

	for i := 0; i < g.V; i++ {
		if vis[i] == 1 {
			continue
		}
		dfsHelper(vis, i, g)
	}
}

func main() {
	g := NewGraph[int](5, 6)
	g.AddEdge(0, 1, 1)
	g.AddEdge(0, 4, 1)
	g.AddEdge(1, 3, 1)
	g.AddEdge(1, 2, 1)
	g.AddEdge(2, 3, 1)
	g.AddEdge(3, 4, 1)

	g.printGraphAdjList()
	fmt.Println()
	g.BFSTraversal()
	fmt.Println()
	g.DFSTraversal()
}
