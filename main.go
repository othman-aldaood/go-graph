package main

import "fmt"

type AdjacencyListGraph struct {
	nodes map[string]struct{}
	edges map[string]map[string]float64
}

// Gibt alle Nachfolger eines Knotens zurück
func (g *AdjacencyListGraph) Successors(nodeId string) []string {
	res := []string{}
	for neighbor := range g.edges[nodeId] {
		res = append(res, neighbor)
	}
	return res
}

// Konstruktor für einen neuen leeren Graphen
func NewAdjacencyListGraph() *AdjacencyListGraph {
	return &AdjacencyListGraph{
		nodes: make(map[string]struct{}),
		edges: make(map[string]map[string]float64),
	}
}

func main() {
	g := NewAdjacencyListGraph()

	g.nodes["A"] = struct{}{}
	g.nodes["B"] = struct{}{}
	g.nodes["C"] = struct{}{}

	if g.edges["A"] == nil {
		g.edges["A"] = make(map[string]float64)
	}
	g.edges["A"]["B"] = 5
	g.edges["A"]["C"] = 2

	if g.edges["B"] == nil {
		g.edges["B"] = make(map[string]float64)
	}
	g.edges["B"]["C"] = 1

	// Anzahl Knoten und Kanten ausgeben
	fmt.Println("Anzahl Knoten:", len(g.nodes))
	count := 0
	for _, targets := range g.edges {
		count += len(targets)
	}
	fmt.Println("Anzahl Kanten:", count)

	fmt.Println("Nachfolger von A:", g.Successors("A"))

}
