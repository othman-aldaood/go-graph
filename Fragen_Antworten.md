
# Fragen und Antworten zur Abgabe – mit Dateiangaben

---

## Frage 1:
**Wo hast du die Datenstruktur für deinen gerichteten Graphen implementiert?**

**Antwort:**  
Die Datenstruktur für den gerichteten Graphen habe ich in der Datei  
`directed_list.go`  
als Struct `AdjacencyListGraph` definiert.  
Hier speichere ich alle Knoten in einer Map und die Kanten als Map von Map.

```go
// Datei: directed_list.go
type AdjacencyListGraph struct {
    nodes map[string]struct{}
    edges map[string]map[string]float64
}
```

---

## Frage 2:
**Wie fügst du einen Knoten und eine Kante hinzu?**

**Antwort:**  
Das mache ich auch in `directed_list.go`
mit den Methoden `AddVertex` und `AddDirectedEdge`.  
Hier ein Ausschnitt:

```go
// Datei: directed_list.go

// Fügt einen neuen Knoten hinzu
func (g *AdjacencyListGraph) AddVertex(nodeId string) {
    g.nodes[nodeId] = struct{}{}
}

// Fügt eine gerichtete Kante mit Gewicht hinzu
func (g *AdjacencyListGraph) AddDirectedEdge(nodeId1, nodeId2 string, length float64) {
    g.AddVertex(nodeId1)
    g.AddVertex(nodeId2)
    g.edges[nodeId1][nodeId2] = length
}
```

---

## Frage 3:
**Wo steht deine Implementierung für Dijkstra?**

**Antwort:**  
Die Implementierung von Dijkstra steht in der Datei `directed_list.go`
in der Methode `Dijkstra`.  
Damit berechne ich die kürzesten Wege von einem Startknoten zu allen anderen Knoten.

```go
// Datei: directed_list.go

func (g *AdjacencyListGraph) Dijkstra(start string) map[string]float64 {
    // ... Dijkstra-Algorithmus wie vorher beschrieben ...
}
```

---

## Frage 4:
**Wie hast du deine Implementierung getestet?**

**Antwort:**  
Die Tests habe ich in der Datei `directed_list_test.go` geschrieben.  
Dort habe ich verschiedene Testfunktionen, zum Beispiel:

```go
// Datei: directed_list_test.go

func TestDirected1(t *testing.T) {
    // Test für einen kleinen Beispielgraphen
    // ...
}

func TestWebGoogleEinlesen(t *testing.T) {
    // Test für das Einlesen einer großen Datei wie web-Google.txt
    // ...
}

func TestDijkstraMitProblem9(t *testing.T) {
    // Test für Dijkstra mit der Datei problem9.8test.txt
    // ...
}
```

---

## Frage 5:
**Wie liest du eine Graph-Datei ein?**

**Antwort:**  
Ich benutze dafür die Hilfsfunktion `initGraph9` in `directed_list_test.go`.  
Damit kann ich zum Beispiel die Datei `problem9.8test.txt` laden und direkt auf den Graphen anwenden.

```go
// Datei: directed_list_test.go

func initGraph9(filename string, graph *AdjacencyListGraph) {
    // ... Zeilenweise einlesen und Knoten/Kanten anlegen ...
}
```

---

## Frage 6:
**Wie rufst du Dijkstra in einem Test auf?**

**Antwort:**  
Zum Beispiel so in `directed_list_test.go`:

```go
func TestDijkstraMitProblem9(t *testing.T) {
    g := NewAdjacencyListGraph()
    initGraph9("problem9.8test.txt", g)
    dist := g.Dijkstra("1")
    // ... Vergleiche das Ergebnis mit dem Erwartungswert ...
}
```

---

## Frage 7:
**Wie kann ich dein Projekt testen?**

**Antwort:**  
Du kannst einfach im Terminal den Befehl  
`go test -v`  
im Projektordner ausführen.  
Dann laufen alle Tests aus der Datei `directed_list_test.go`.

---
