package graph

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"testing"
)

// Test für kleinen Beispiel-Graphen, wie in der Aufgabenstellung
func TestDirected1(t *testing.T) {
	// Erzeuge einen neuen Graphen
	g := NewAdjacencyListGraph()

	// Lege die Knoten an
	vertices := []string{"A", "B", "C", "D", "E", "F"}
	for _, v := range vertices {
		g.AddVertex(v)
	}

	// Füge Kanten von A zu B, C, D, E mit Gewicht 1 hinzu
	for _, v := range vertices[1:5] {
		g.AddDirectedEdge(vertices[0], v, 1.)
	}
	// Füge Kante von E zu A mit Gewicht 1 hinzu
	g.AddDirectedEdge(vertices[4], vertices[0], 1.)
	// Füge Kante von F zu A mit Gewicht 1 hinzu
	g.AddDirectedEdge(vertices[5], vertices[0], 1.)
	// Füge Kante von D zu C mit Gewicht 1 hinzu
	g.AddDirectedEdge(vertices[3], vertices[2], 1.)

	// Prüfe die Anzahl der Knoten
	if g.NumVertices() != len(vertices) {
		t.Fatalf("NumVertices failed!") // Test schlägt fehl, wenn Anzahl nicht stimmt
	}

	// Prüfe die Anzahl der Kanten (sollte 7 sein)
	if g.NumEdges() != 7 {
		t.Fatalf("NumEdges failed!")
	}

	// Prüfe die Successors von A (B, C, D, E)
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

	// Prüfe die Predecessors von A (E und F)
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

// Test für das Einlesen der Datei web-Google.txt
func TestWebGoogleEinlesen(t *testing.T) {
	// Erzeuge neuen leeren Graphen
	g := NewAdjacencyListGraph()

	// Öffne die Datei (muss im gleichen Ordner liegen)
	file, _ := os.Open("web-Google.txt")
	defer file.Close()

	// Zeile für Zeile einlesen
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := scanner.Text()
		fields := strings.Fields(s)
		// Nur Zeilen mit zwei Werten und ohne '#' am Anfang nehmen
		if len(fields) == 2 && !strings.HasPrefix(fields[0], "#") {
			g.AddVertex(fields[0])
			g.AddVertex(fields[1])
			g.AddDirectedEdge(fields[0], fields[1], 1.0)
		}
	}
	// Gibt Anzahl Knoten und Kanten zur Kontrolle im Test-Log aus
	t.Logf("Knoten: %d, Kanten: %d", g.NumVertices(), g.NumEdges())

}

// Hilfsfunktion zum Einlesen von problem9.8test.txt
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

// Test für Dijkstra mit problem9.8test.txt
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

func TestBFS(t *testing.T) {
	g := NewAdjacencyListGraph()
	g.AddDirectedEdge("A", "B", 1)
	g.AddDirectedEdge("A", "C", 1)
	g.AddDirectedEdge("B", "D", 1)
	g.AddDirectedEdge("C", "D", 1)
	g.AddDirectedEdge("D", "E", 1)

	// Erwartete Distanzen von "A"
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

// Test für Dijkstra mit der großen Datei problem9.8.txt und den 10 Zielknoten aus der Aufgabe
func TestDijkstraChallengeProblem9(t *testing.T) {
	// Erzeuge einen neuen leeren gerichteten Graphen (Adjazenzlisten-Graph)
	g := NewAdjacencyListGraph()

	// Initialisiere den Graphen aus der Datei problem9.8.txt
	// (die Datei muss im gleichen Ordner wie dieses Test-Programm liegen)
	initGraph9("problem9.8.txt", g)

	// Berechne die kürzesten Wege von Knoten "1" zu allen anderen Knoten mit Dijkstra
	dist := g.Dijkstra("1")

	// Liste der gefragten 10 Zielknoten aus der Aufgabenstellung
	nodes := []string{"7", "37", "59", "82", "99", "115", "133", "165", "188", "197"}

	// Gib für jeden Zielknoten die berechnete Distanz im Test-Log aus
	t.Logf("Dijkstra-Distanzen von 1 zu den 10 Zielknoten (problem9.8.txt):")
	for _, node := range nodes {
		t.Logf("%s: %.0f", node, dist[node])
	}
}
