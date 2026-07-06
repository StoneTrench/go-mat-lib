package lib

import "iter"

type Ident uint32

type Graph[T any] struct {
	IndexCounter Ident
	Nodes        map[Ident]T
	Incoming     map[Ident][]Ident
	Outgoing     map[Ident][]Ident
}

func sliceRemoveAt[T any](s []T, i int) []T {
	last := len(s) - 1
	if last <= 0 {
		return s
	}
	s[i] = s[last]
	s = s[:last]
	return s
}

func (g Graph[T]) Init() Graph[T] {
	g.IndexCounter = 0
	g.Nodes = make(map[Ident]T)
	g.Incoming = make(map[Ident][]Ident)
	g.Outgoing = make(map[Ident][]Ident)

	return g
}
func (g *Graph[T]) getNextIdent() (index Ident) {
	index = g.IndexCounter
	g.IndexCounter = Ident(index + 1)
	return index
}
func (g *Graph[T]) AddNode(body T) Ident {
	i := g.getNextIdent()
	g.Nodes[i] = body
	return i
}
func (g *Graph[T]) AddEdge(a, b Ident) {
	g.Outgoing[a] = append(g.Outgoing[a], b)
	g.Incoming[b] = append(g.Incoming[b], a)
}

func (g *Graph[T]) RemoveEdges(a Ident) {
	outgoing := g.Outgoing[a]
	for _, b_index := range outgoing {
		b := g.Incoming[b_index]
		for i := len(b) - 1; i >= 0; i-- {
			if b[i] == a {
				b = sliceRemoveAt(b, i)
			}
		}
		g.Incoming[b_index] = b
	}
	delete(g.Outgoing, a)
}
func (g *Graph[T]) RemoveNode(a Ident) {
	g.RemoveEdges(a)
	delete(g.Nodes, a)
}
func (g *Graph[T]) RemoveEdge(a, b Ident) {
	var lst []Ident
	lst = g.Outgoing[a]
	for i, e := range lst {
		if e == b {
			lst = sliceRemoveAt(lst, i)
			break
		}
	}
	g.Outgoing[a] = lst

	lst = g.Incoming[b]
	for i, e := range lst {
		if e == a {
			lst = sliceRemoveAt(lst, i)
			break
		}
	}
	g.Incoming[b] = lst
}

func (g *Graph[T]) Iterate() iter.Seq[Ident] {
	return func(yield func(Ident) bool) {
		for i := range g.Nodes {
			if !yield(i) {
				return
			}
		}
	}
}
