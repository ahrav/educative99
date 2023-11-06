package graphs

import "fmt"

// Vertex represents a graph vertex.
type Vertex struct {
	key      int
	adjacent []*Vertex
}

// Graph represents an adjacency list graph.
type Graph struct {
	vertices []*Vertex
}

// InitVertex initializes a vertex.
func InitVertex(k int) *Vertex {
	return &Vertex{key: k}
}

// AddVertex adds a vertex to a graph.
func (g *Graph) AddVertex(v *Vertex) {
	g.vertices = append(g.vertices, v)
}

// AddEdgeWithVertices adds an edge to a graph between two vertices.
func (g *Graph) AddEdgeWithVertices(u, v *Vertex) {
	u.adjacent = append(u.adjacent, v)
	v.adjacent = append(v.adjacent, u)
}

// AddEdge adds an edge to a graph between two keys.
func (g *Graph) AddEdge(u, v int) {
	uVertex := g.getVertex(u)
	vVertex := g.getVertex(v)
	g.AddEdgeWithVertices(uVertex, vVertex)
}

// getVertex returns a vertex with a given key.
func (g *Graph) getVertex(k int) *Vertex {
	for _, v := range g.vertices {
		if v.key == k {
			return v
		}
	}
	return nil
}

// Print prints a graph.
func (g *Graph) Print() {
	for _, v := range g.vertices {
		for _, a := range v.adjacent {
			fmt.Printf("%v -> %v\n", v.key, a.key)
		}
	}
}

// BFS performs a breadth-first search on a graph.
func (g *Graph) BFS(start int) {
	startVertex := g.getVertex(start)
	if startVertex == nil {
		return // No vertex to start from.
	}

	visited := make(map[*Vertex]struct{})
	queue := []*Vertex{startVertex}
	visited[startVertex] = struct{}{}

	for len(queue) > 0 {
		curr, queue := queue[0], queue[1:]

		fmt.Printf("Visited %v\n", curr)

		for _, adj := range curr.adjacent {
			if _, ok := visited[adj]; !ok {
				visited[adj] = struct{}{}
				queue = append(queue, adj)
			}
		}
	}
}
