package main

import (
	"fmt"
	"log"

	"github.com/liyue201/gostl/ds/stack"
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

func record_finishing_time_in_dfs(g *Graph[int], st *stack.Stack[int], curr int, vis []bool) {
	vis[curr] = true
	for i := 0; i < len(g.Adj()[curr]); i++ {
		v := g.Adj()[curr][i].Id
		if vis[v] {
			continue
		}
		record_finishing_time_in_dfs(g, st, v, vis)
	}
	st.Push(curr)
}

func dfs(g *Graph[int], curr int, vis []bool, scc []int) []int {
	vis[curr] = true
	scc = append(scc, curr)
	for i := 0; i < len(g.Adj()[curr]); i++ {
		v := g.Adj()[curr][i].Id
		if !vis[v] {
			scc = dfs(g, v, vis, scc)
		}
	}
	return scc
}

func strongly_connected_components_kosaraju_algo(g *Graph[int]) [][]int {
	st := stack.New[int]()
	vis := make([]bool, g.V)
	for i := 0; i < g.V; i++ {
		if vis[i] {
			continue
		}
		record_finishing_time_in_dfs(g, st, i, vis)
	}

	// reverse the edges of the graph
	revG := NewGraph[int](g.V, g.E)
	for i := 0; i < g.V; i++ {
		for j := 0; j < len(g.Adj()[i]); j++ {
			u, v, wt := i, g.Adj()[i][j].Id, g.Adj()[i][j].Wt
			revG.AddDirectedEdge(v, u, wt)
		}
	}
	vis = make([]bool, revG.V)
	res := [][]int{}
	for !st.Empty() {
		top := st.Top()
		st.Pop()
		scc := []int{}
		if !vis[top] {
			scc = dfs(revG, top, vis, scc)
		}
		if len(scc) > 0 {
			res = append(res, scc)
		}
	}

	return res
}

// input  : 5 5 0 2 2 1 1 0 0 3 3 4
// output : [[0 1 2] [3] [4]]

func main() {
	v, e := readInt(), readInt()
	g := NewGraph[int](v, e)
	for i := 0; i < e; i++ {
		u, v, wt := readInt(), readInt(), 1
		g.AddDirectedEdge(u, v, wt)
	}

	fmt.Println(strongly_connected_components_kosaraju_algo(g))

}
