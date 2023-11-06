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

// Neighbors returns a vertex's neighbors.
func (v *Vertex) Neighbors() []*Vertex {
	return v.adjacent
}

// IsBipartite returns true if a graph is bipartite.
func (g *Graph) IsBipartite() bool {
	if len(g.vertices) == 0 {
		return true
	}

	const (
		uncolored = 0
		red       = 1
		blue      = -1
	)

	colors := make(map[*Vertex]int)
	for _, vertex := range g.vertices {
		if colors[vertex] != uncolored {
			continue // Skip vertex if it's already colored.
		}

		// Perform BFS on vertex.
		queue := []*Vertex{vertex}
		colors[vertex] = red

		for len(queue) > 0 {
			curr, queue := queue[0], queue[1:]

			for _, neighbor := range curr.Neighbors() {
				if _, ok := colors[neighbor]; !ok {
					// Neighbor is uncolored, so color it the opposite color of curr.
					queue = append(queue, neighbor)
					colors[neighbor] = colors[curr] * -1
				} else if colors[neighbor] == colors[curr] {
					// Neighbor is colored the same color as curr, so graph is not bipartite.
					return false
				}
			}
		}
	}

	return true
}

// DFS performs a depth-first search on a graph.
func (g *Graph) DFS() {
	// Track visited vertices.
	visited := make(map[*Vertex]struct{})
	for _, v := range g.vertices {
		if _, ok := visited[v]; !ok {
			g.dfs(v, visited)
		}
	}
}

// dfs is a helper function for DFS.
func (g *Graph) dfs(v *Vertex, visited map[*Vertex]struct{}) {
	// Mark v as visited.
	visited[v] = struct{}{}
	fmt.Printf("Visited %v\n", v)

	// Visit v's neighbors.
	for _, a := range v.adjacent {
		if _, ok := visited[a]; !ok {
			g.dfs(a, visited)
		}
	}
}

// HasCycle returns true if a graph has a cycle.
func (g *Graph) HasCycle() bool {
	// Track visited vertices.
	visited := make(map[*Vertex]struct{})
	for _, v := range g.vertices {
		if _, ok := visited[v]; !ok {
			if g.hasCycle(v, visited, nil) {
				return true
			}
		}
	}
	return false
}

// hasCycle is a helper function for HasCycle.
func (g *Graph) hasCycle(v *Vertex, visited map[*Vertex]struct{}, parent *Vertex) bool {
	// Mark v as visited.
	visited[v] = struct{}{}

	// Visit v's neighbors.
	for _, a := range v.adjacent {
		if _, ok := visited[a]; !ok {
			if g.hasCycle(a, visited, v) {
				return true
			}
		} else if a != parent {
			return true
		}
	}
	return false
}
