package graph

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"testing"
)

// TestDirected1 prüft die Grundfunktionen des gerichteten Graphen.
func TestDirected1(t *testing.T) {
	g := NewAdjacencyListGraph()
	vertices := []string{"A", "B", "C", "D", "E", "F"}
	for _, v := range vertices {
		g.AddVertex(v)
	}
	for _, v := range vertices[1:5] {
		g.AddDirectedEdge(vertices[0], v, 1.)
	}
	g.AddDirectedEdge(vertices[4], vertices[0], 1.)
	g.AddDirectedEdge(vertices[5], vertices[0], 1.)
	g.AddDirectedEdge(vertices[3], vertices[2], 1.)

	if g.NumVertices() != len(vertices) {
		t.Fatalf("NumVertices failed!")
	}
	if g.NumEdges() != 7 {
		t.Fatalf("NumEdges failed!")
	}
	// Prüfe Successors von A
	suc := g.Successors("A")
	for _, v := range vertices[1:5] {
		found := false
		for _, s := range suc {
			if s == v {
				found = true
			}
		}
		if !found {
			t.Fatalf("%v nicht in Successors von A", v)
		}
	}
	// Prüfe Predecessors von A (E und F)
	pre := g.Predecessors("A")
	foundE, foundF := false, false
	for _, p := range pre {
		if p == "E" {
			foundE = true
		}
		if p == "F" {
			foundF = true
		}
	}
	if !(foundE && foundF) {
		t.Fatalf("E oder F nicht in Predecessors von A")
	}
}

// TestWebGoogleEinlesen testet das Einlesen der Datei web-Google.txt
func TestWebGoogleEinlesen(t *testing.T) {
	g := NewAdjacencyListGraph()
	file, _ := os.Open("web-Google.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := scanner.Text()
		fields := strings.Fields(s)
		if len(fields) == 2 && !strings.HasPrefix(fields[0], "#") {
			g.AddVertex(fields[0])
			g.AddVertex(fields[1])
			g.AddDirectedEdge(fields[0], fields[1], 1.0)
		}
	}
	t.Logf("Knoten: %d, Kanten: %d", g.NumVertices(), g.NumEdges())
}

// initGraph9 liest einen gewichteten Graphen im Aufgabenformat ein.
func initGraph9(filename string, graph *AdjacencyListGraph) {
	file, err := os.Open(filename)
	if err != nil {
		panic("Kann Datei nicht öffnen: " + filename)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := scanner.Text()
		fields := strings.Fields(s)
		id1 := fields[0]
		graph.AddVertex(id1)
		for _, x := range fields[1:] {
			f := strings.Split(x, ",")
			length, _ := strconv.ParseFloat(f[1], 64)
			graph.AddVertex(f[0])
			graph.AddDirectedEdge(id1, f[0], length)
		}
	}
}

// TestDijkstraMitProblem9 prüft Dijkstra mit problem9.8test.txt (bekannte Musterlösung)
func TestDijkstraMitProblem9(t *testing.T) {
	g := NewAdjacencyListGraph()
	initGraph9("problem9.8test.txt", g)
	dist := g.Dijkstra("1")
	expected := map[string]float64{
		"1": 0, "2": 1, "3": 2, "4": 3, "5": 4, "6": 4, "7": 3, "8": 2,
	}
	for node, exp := range expected {
		if dist[node] != exp {
			t.Fatalf("Falsche Distanz für Knoten %s: erwartet %.0f, gefunden %.0f", node, exp, dist[node])
		}
	}
	t.Log("Dijkstra-Test für problem9.8test.txt bestanden!")
}

// TestBFS prüft die BFS-Implementierung
func TestBFS(t *testing.T) {
	g := NewAdjacencyListGraph()
	g.AddDirectedEdge("A", "B", 1)
	g.AddDirectedEdge("A", "C", 1)
	g.AddDirectedEdge("B", "D", 1)
	g.AddDirectedEdge("C", "D", 1)
	g.AddDirectedEdge("D", "E", 1)
	want := map[string]int{
		"A": 0,
		"B": 1,
		"C": 1,
		"D": 2,
		"E": 3,
	}
	got := g.BFS("A")
	for node, exp := range want {
		val, ok := got[node]
		if !ok {
			t.Errorf("Knoten %s fehlt im Ergebnis!", node)
		}
		if val != exp {
			t.Errorf("Knoten %s: erwartet %d, gefunden %d", node, exp, val)
		}
	}
}

// TestDijkstraChallengeProblem9 prüft Dijkstra mit problem9.8.txt und gibt die Distanzen zu den 10 Zielknoten aus.
func TestDijkstraChallengeProblem9(t *testing.T) {
	g := NewAdjacencyListGraph()
	initGraph9("problem9.8.txt", g)
	dist := g.Dijkstra("1")
	nodes := []string{"7", "37", "59", "82", "99", "115", "133", "165", "188", "197"}
	t.Logf("Dijkstra-Distanzen von 1 zu den 10 Zielknoten (problem9.8.txt):")
	for _, node := range nodes {
		t.Logf("%s: %.0f", node, dist[node])
	}
}
