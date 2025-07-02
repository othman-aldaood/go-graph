// heap.go
package graph

// HeapNode speichert einen Knoten und dessen Distanz (Priority)
type HeapNode struct {
	node     string
	distance float64
}

// MinHeap implementiert einen Min-Heap für HeapNode
type MinHeap struct {
	array  []HeapNode     // Slice der HeapNodes
	lookup map[string]int // Knotenname -> Index in array (für O(1)-Zugriff)
}

// NewMinHeap erzeugt einen neuen leeren MinHeap
func NewMinHeap() *MinHeap {
	return &MinHeap{
		array:  []HeapNode{},
		lookup: make(map[string]int),
	}
}

// Len gibt die Länge des Heaps zurück
func (h *MinHeap) Len() int {
	return len(h.array)
}

// Push fügt ein neues Element in den Heap ein (oder updated es)
func (h *MinHeap) Push(node string, dist float64) {
	if i, ok := h.lookup[node]; ok {
		// Update, falls schon im Heap
		h.array[i].distance = dist
		h.up(i)
		h.down(i)
		return
	}
	hn := HeapNode{node: node, distance: dist}
	h.array = append(h.array, hn)
	idx := len(h.array) - 1
	h.lookup[node] = idx
	h.up(idx)
}

// Pop entfernt und gibt das Element mit kleinster Distanz zurück
func (h *MinHeap) Pop() (string, float64) {
	if h.Len() == 0 {
		return "", 0
	}
	minNode := h.array[0]
	lastIdx := h.Len() - 1
	h.array[0] = h.array[lastIdx]
	h.lookup[h.array[0].node] = 0
	h.array = h.array[:lastIdx]
	delete(h.lookup, minNode.node)
	h.down(0)
	return minNode.node, minNode.distance
}

// up verschiebt ein Element nach oben, bis Heap-Bedingung stimmt
func (h *MinHeap) up(i int) {
	for {
		parent := (i - 1) / 2
		if i == 0 || h.array[parent].distance <= h.array[i].distance {
			break
		}
		h.swap(i, parent)
		i = parent
	}
}

// down verschiebt ein Element nach unten, bis Heap-Bedingung stimmt
func (h *MinHeap) down(i int) {
	n := h.Len()
	for {
		left := 2*i + 1
		right := 2*i + 2
		smallest := i
		if left < n && h.array[left].distance < h.array[smallest].distance {
			smallest = left
		}
		if right < n && h.array[right].distance < h.array[smallest].distance {
			smallest = right
		}
		if smallest == i {
			break
		}
		h.swap(i, smallest)
		i = smallest
	}
}

// swap tauscht zwei Elemente im Heap (inkl. lookup-Update)
func (h *MinHeap) swap(i, j int) {
	h.array[i], h.array[j] = h.array[j], h.array[i]
	h.lookup[h.array[i].node] = i
	h.lookup[h.array[j].node] = j
}

// Update ändert die Distanz eines Knotens, falls er im Heap ist
func (h *MinHeap) Update(node string, dist float64) {
	if idx, ok := h.lookup[node]; ok {
		h.array[idx].distance = dist
		h.up(idx)
		h.down(idx)
	}
}

// Exist prüft, ob ein Knoten noch im Heap liegt
func (h *MinHeap) Exist(node string) bool {
	_, ok := h.lookup[node]
	return ok
}
