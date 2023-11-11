package graphs

import (
	"container/heap"
	"fmt"
	"math"
	"sort"
)

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

// DFSTopologicalSort performs a topological sort on a graph using DFS.
func (g *Graph) DFSTopologicalSort() ([]*Vertex, error) {
	// Track visited vertices.
	visited := make(map[*Vertex]struct{})
	// Track sorted vertices.
	var sorted []*Vertex

	for _, v := range g.vertices {
		if _, ok := visited[v]; !ok {
			if err := g.topologicalSort(v, visited, &sorted, nil); err != nil {
				return nil, err
			}
		}
	}

	return sorted, nil
}

// topologicalSort is a helper function for TopologicalSort.
func (g *Graph) topologicalSort(v *Vertex, visited map[*Vertex]struct{}, sorted *[]*Vertex, parent *Vertex) error {
	// Mark v as visited.
	visited[v] = struct{}{}

	// Visit v's neighbors.
	for _, a := range v.adjacent {
		if _, ok := visited[a]; !ok {
			if err := g.topologicalSort(a, visited, sorted, v); err != nil {
				return err
			}
		} else if a != parent {
			return fmt.Errorf("graph is not a DAG")
		}
	}

	// Prepend v to sorted.
	*sorted = append([]*Vertex{v}, *sorted...)

	return nil
}

// KahnTopologicalSort performs a topological sort on a graph using Kahn's algorithm.
func (g *Graph) KahnTopologicalSort() ([]*Vertex, error) {
	// Track in-degrees of vertices.
	inDegrees := make(map[*Vertex]int)
	for _, v := range g.vertices {
		inDegrees[v] = 0
		for _, a := range v.adjacent {
			inDegrees[a]++
		}
	}

	// Initialize queue with vertices with in-degree 0.
	var queue []*Vertex
	for v, inDegree := range inDegrees {
		if inDegree == 0 {
			queue = append(queue, v)
		}
	}

	// Track sorted vertices.
	var sorted []*Vertex

	for len(queue) > 0 {
		// Pop vertex from queue and append to sorted.
		curr, queue := queue[0], queue[1:]
		sorted = append(sorted, curr)

		// Decrement in-degrees of curr's neighbors.
		// Essentially removing edges from curr to its neighbors.
		for _, a := range curr.adjacent {
			inDegrees[a]--
			if inDegrees[a] == 0 {
				queue = append(queue, a)
			}
		}
	}

	// If there are any vertices left in the graph, then there is a cycle.
	for _, v := range g.vertices {
		if inDegrees[v] != 0 {
			return nil, fmt.Errorf("graph is not a DAG")
		}
	}

	return sorted, nil
}

type UnionFind struct {
	parent []int
	size   []int
}

func NewUnionFind(size int) *UnionFind {
	// Each vertex is its own parent and the representative of its own set.
	parent := make([]int, size)
	for i := range parent {
		parent[i] = i
	}

	// Each set has size 1 initially since each vertex is its own set.
	sizeArr := make([]int, size)
	for i := range sizeArr {
		sizeArr[i] = 1
	}

	return &UnionFind{parent: parent, size: sizeArr}
}

func (uf *UnionFind) Find(rep int) int {
	// If rep is not the representative of its set, then recursively find the representative of its set.
	if uf.parent[rep] != rep {
		// Path compression.
		uf.parent[rep] = uf.Find(uf.parent[rep])
	}
	return uf.parent[rep]
}

func (uf *UnionFind) Union(rep1, rep2 int) {
	// Find the representatives of the sets that rep1 and rep2 belong to.
	rep1 = uf.Find(rep1)
	rep2 = uf.Find(rep2)

	// If rep1 and rep2 are already in the same set, then do nothing.
	if rep1 == rep2 {
		return
	}

	// Merge the smaller set into the larger set.
	if uf.size[rep1] < uf.size[rep2] {
		uf.parent[rep1] = rep2
		uf.size[rep2] += uf.size[rep1]
	} else {
		uf.parent[rep2] = rep1
		uf.size[rep1] += uf.size[rep2]
	}
}

func (uf *UnionFind) Connected(rep1, rep2 int) bool {
	return uf.Find(rep1) == uf.Find(rep2)
}

// Edge represents an edge in a graph.
type Edge struct {
	src, dst, weight int
}

// WeightedGraph represents an adjacency list weighted graph.
type WeightedGraph struct {
	numVertices int
	edges       []Edge
}

// AddEdge adds an edge to a weighted graph.
func (g *WeightedGraph) AddEdge(src, dst, weight int) {
	g.edges = append(g.edges, Edge{src, dst, weight})
	g.edges = append(g.edges, Edge{dst, src, weight})
}

// NewWeightedGraph constructs a new weighted graph.
func NewWeightedGraph(numVertices int) *WeightedGraph {
	return &WeightedGraph{numVertices: numVertices}
}

type MinHeap []Edge

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].weight < h[j].weight }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x any) {
	*h = append(*h, x.(Edge))
}

func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// MST returns the minimum spanning tree of a weighted graph.
type MST []Edge

// PrimMST returns the minimum spanning tree of a weighted graph using Prim's algorithm.
func (g *WeightedGraph) PrimMST() MST {
	if g.numVertices == 0 {
		return MST{}
	}

	// Initialize all keys as infinite and mstSet as false for all vertices
	keys := make([]int, g.numVertices)
	mstSet := make([]bool, g.numVertices)
	for i := range keys {
		keys[i] = math.MaxInt32
	}
	keys[0] = 0 // Start from the first vertex

	pq := new(MinHeap)
	heap.Init(pq)
	heap.Push(pq, Edge{weight: 0, src: -1, dst: 0})

	mst := MST{}

	for pq.Len() > 0 {
		// Pick the smallest edge in the min heap
		u := heap.Pop(pq).(Edge).dst
		if mstSet[u] {
			continue
		}
		mstSet[u] = true // Include vertex in MST
		if u != 0 {      // If not the first vertex, add to the MST result
			mst = append(mst, Edge{src: u, dst: u, weight: keys[u]})
		}

		// Update the key values and parent index of the adjacent vertices
		for _, e := range g.edges {
			if !mstSet[e.dst] && e.weight < keys[e.dst] {
				keys[e.dst] = e.weight
				heap.Push(pq, Edge{weight: keys[e.dst], src: u, dst: e.dst})
			}
		}
	}

	return mst
}

// KruskalMST returns the minimum spanning tree of a weighted graph using Kruskal's algorithm.
func (g *WeightedGraph) KruskalMST() MST {
	if g.numVertices == 0 {
		return MST{}
	}

	// Sort edges by weight.
	edges := g.edges
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].weight < edges[j].weight
	})

	uf := NewUnionFind(g.numVertices)
	mst := MST{}

	for _, e := range edges {
		if !uf.Connected(e.src, e.dst) {
			uf.Union(e.src, e.dst)
			mst = append(mst, e)
		}
	}

	return mst
}

// Dijkstra returns the shortest path from a source vertex to all other vertices in a weighted graph.
func (g *WeightedGraph) Dijkstra(src int) []int {
	if g.numVertices == 0 {
		return []int{}
	}

	// Initialize all distances as infinite and sptSet as false for all vertices
	dist := make([]int, g.numVertices)
	sptSet := make([]bool, g.numVertices)
	for i := range dist {
		dist[i] = math.MaxInt32
	}
	dist[src] = 0 // Distance from source to source is 0

	for i := 0; i < g.numVertices-1; i++ {
		// Pick the smallest distance vertex from the set of vertices not yet processed.
		u := minDistance(dist, sptSet)
		sptSet[u] = true // Include vertex in shortest path tree

		// Update the distance value of the adjacent vertices of the picked vertex.
		for _, e := range g.edges {
			if !sptSet[e.dst] && dist[u] != math.MaxInt32 && dist[u]+e.weight < dist[e.dst] {
				dist[e.dst] = dist[u] + e.weight
			}
		}
	}

	return dist
}

// minDistance returns the vertex with the minimum distance from the set of vertices not yet processed.
func minDistance(dist []int, sptSet []bool) int {
	minV := math.MaxInt32
	minIndex := -1

	for i, d := range dist {
		if !sptSet[i] && d <= minV {
			minV = d
			minIndex = i
		}
	}

	return minIndex
}
