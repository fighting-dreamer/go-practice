package main

import (
	"fmt"
	"log"

	"github.com/liyue201/gostl/ds/priorityqueue"
)

type Node[T any] struct {
	Id int
	Wt T //
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

func find_balances(g *Graph[int]) map[int]int {
	balanceMap := map[int]int{}
	for i := 0; i < g.V; i++ {
		balanceMap[i] = 0
	}

	for i := 0; i < g.V; i++ {
		for j := 0; j < len(g.Adj()[i]); j++ {
			u := i
			v := g.Adj()[i][j].Id
			cash := g.Adj()[i][j].Wt
			// u let it to v => v have to return this to u
			balanceMap[u] -= cash
			balanceMap[v] += cash
		}
	}

	return balanceMap
}

func find_min_transaction(balanceMap map[int]int) int {
	// this is not completely right way
	// in this approach : we will keep two kinds of people :
	// ---- one who have -ve balance => they gave money as credit,
	// ---- one with +ve balance they took money as debt.
	// we try to mind the max amount that can be settled using these two and what amount is left will be put back in hte balanceMap

	// Implementation : we will use the max heap.
	// for creditors, we will use -ver value or absolute value to ensure right ordering.
	compareFunc := func(a, b Node[int]) int {
		if a.Wt == b.Wt {
			return 0
		}
		if a.Wt < b.Wt {
			return 1
		}
		return -1
	}
	creditorHeap, debtHeap := priorityqueue.New[Node[int]](compareFunc, priorityqueue.WithGoroutineSafe()), priorityqueue.New[Node[int]](compareFunc, priorityqueue.WithGoroutineSafe())

	for k, v := range balanceMap {
		if v > 0 {
			// debtors
			debtHeap.Push(Node[int]{k, v})
		} else {
			creditorHeap.Push(Node[int]{k, -v})
		}
	}

	res := 0

	for !creditorHeap.Empty() && !debtHeap.Empty() {
		topCreditor := creditorHeap.Pop()
		topDebtor := debtHeap.Pop()

		if topCreditor.Wt > topDebtor.Wt {
			// creditor have let more money and will still be required to paid back the remaining.
			topCreditor.Wt -= topDebtor.Wt
			creditorHeap.Push(topCreditor)
		} else if topCreditor.Wt < topDebtor.Wt {
			// debtor have got more money and will still be required to paid back the remaining.
			topDebtor.Wt -= topCreditor.Wt
			debtHeap.Push(topDebtor)
		}
		res++
	}
	
	return res
}

func main() {
	v, e := readInt(), readInt() // number of friends mad transactions count.
	g := NewGraph[int](v, e)
	for i := 0; i < e; i++ {
		g.AddDirectedEdge(readInt(), readInt(), readInt()) // u, v, cash that u have lent to v ie v have to pay back this amount to u.
	}

	friend2balancesMap := find_balances(g)
	fmt.Println(find_min_transaction(friend2balancesMap))
}
