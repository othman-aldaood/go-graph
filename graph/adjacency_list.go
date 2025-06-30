package graph

import "math"

// Adjazenzlisten-Graph für gerichtete Kanten
type AdjacencyListGraph struct {
	nodes map[string]struct{}           // Set von Knoten
	edges map[string]map[string]float64 // Knoten -> Nachbarknoten -> Kantengewicht
}

// Erzeugt einen neuen leeren gerichteten Graphen
func NewAdjacencyListGraph() *AdjacencyListGraph {
	return &AdjacencyListGraph{
		nodes: make(map[string]struct{}),
		edges: make(map[string]map[string]float64),
	}
}

// Fügt einen Knoten ein (macht nichts, falls der Knoten schon existiert)
func (g *AdjacencyListGraph) AddVertex(nodeId string) {
	if _, exists := g.nodes[nodeId]; exists {
		return
	}
	g.nodes[nodeId] = struct{}{}
	g.edges[nodeId] = make(map[string]float64)
}

// Gibt die Anzahl der Knoten zurück
func (g *AdjacencyListGraph) NumVertices() int {
	return len(g.nodes)
}

// Gibt die Anzahl der Kanten zurück
func (g *AdjacencyListGraph) NumEdges() int {
	count := 0
	for _, targets := range g.edges {
		count += len(targets)
	}
	return count
}

// Fügt eine gerichtete Kante mit Gewicht ein
func (g *AdjacencyListGraph) AddDirectedEdge(nodeId1, nodeId2 string, length float64) {
	g.AddVertex(nodeId1)
	g.AddVertex(nodeId2)
	g.edges[nodeId1][nodeId2] = length
}

// Liefert alle Nachfolger eines Knotens (Successors)
func (g *AdjacencyListGraph) Successors(nodeId string) []string {
	res := []string{}
	for neighbor := range g.edges[nodeId] {
		res = append(res, neighbor)
	}
	return res
}

// Liefert alle Vorgänger eines Knotens (Predecessors)
func (g *AdjacencyListGraph) Predecessors(nodeId string) []string {
	res := []string{}
	for from, neighbors := range g.edges {
		if _, ok := neighbors[nodeId]; ok {
			res = append(res, from)
		}
	}
	return res
}

// BFS: Breadth-First-Search, liefert für jeden erreichbaren Knoten die Distanz in "Layern"
func (g *AdjacencyListGraph) BFS(start string) map[string]int {
	dist := make(map[string]int)
	queue := []string{start}
	dist[start] = 0

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		for neighbor := range g.edges[current] {
			if _, visited := dist[neighbor]; !visited {
				dist[neighbor] = dist[current] + 1
				queue = append(queue, neighbor)
			}
		}
	}
	return dist
}

// DFS: Depth-First-Search, rekursiv
func (g *AdjacencyListGraph) DFS(start string) map[string]bool {
	visited := make(map[string]bool)
	var dfs func(string)
	dfs = func(node string) {
		visited[node] = true
		for neighbor := range g.edges[node] {
			if !visited[neighbor] {
				dfs(neighbor)
			}
		}
	}
	dfs(start)
	return visited
}

// Platzhalter für Dijkstra (noch nicht implementiert)
func (g *AdjacencyListGraph) Dijkstra(start string) map[string]float64 {
	dist := make(map[string]float64)
	visited := make(map[string]bool)
	for node := range g.nodes {
		dist[node] = math.Inf(1)
	}
	dist[start] = 0
	for {
		minDist := math.Inf(1)
		var u string
		found := false
		for node := range g.nodes {
			if !visited[node] && dist[node] < minDist {
				minDist = dist[node]
				u = node
				found = true
			}
		}
		if !found {
			break
		}
		visited[u] = true
		for v, weight := range g.edges[u] {
			if dist[u]+weight < dist[v] {
				dist[v] = dist[u] + weight
			}
		}
	}
	return dist
}
