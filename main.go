package main

import (
	"fmt"
	"imp/graph"
)

// Main-Demo für die wichtigsten Methoden des Adjazenzlisten-Graphen
func main() {
	// Neuen gerichteten Graphen erzeugen
	g := graph.NewAdjacencyListGraph()

	// Knoten und Kanten per Methoden anlegen
	g.AddVertex("A")
	g.AddVertex("B")
	g.AddVertex("C")
	g.AddDirectedEdge("A", "B", 5)
	g.AddDirectedEdge("A", "C", 2)
	g.AddDirectedEdge("B", "C", 1)

	// Anzahl Knoten und Kanten ausgeben
	fmt.Println("Anzahl Knoten:", g.NumVertices())
	fmt.Println("Anzahl Kanten:", g.NumEdges())

	// Nachfolger von A ausgeben
	fmt.Println("Nachfolger von A:", g.Successors("A"))

	// Beispiel: BFS ab "A"
	dist := g.BFS("A")
	fmt.Println("BFS-Distanzen von A:", dist)

	// Beispiel: DFS ab "A"
	visited := g.DFS("A")
	fmt.Println("Von A aus besuchte Knoten (DFS):", visited)

	// Beispiel: Dijkstra ab "A"
	shortest := g.Dijkstra("A")
	fmt.Println("Kürzeste Wege von A (Dijkstra):", shortest)
}
