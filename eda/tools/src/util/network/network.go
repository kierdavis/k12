package network

// Standard library imports
import (
    "fmt"
)

type Node uint

type Weight interface {
	Less(other Weight) bool
	Add(other Weight) Weight
}

type IntWeight int

func (w IntWeight) Less(other Weight) bool {
	return w < other.(IntWeight)
}

func (w IntWeight) Add(other Weight) Weight {
	return w + other.(IntWeight)
}

type FloatWeight float64

func (w FloatWeight) Less(other Weight) bool {
	return w < other.(FloatWeight)
}

func (w FloatWeight) Add(other Weight) Weight {
	return w + other.(FloatWeight)
}

type Edge struct {
	A Node
	B Node
	Weight Weight
}

type Undirected struct {
	Edges []Edge
}

func NewUndirected() (g *Undirected) {
	return &Undirected{}
}

func (g *Undirected) Reset() {
	g.Edges = nil
}

func (g *Undirected) Dump() {
    for _, edge := range g.Edges {
        fmt.Printf("[%d] -> [%d] (weight %v)\n", edge.A, edge.B, edge.Weight)
    }
}

func (g *Undirected) Weight() Weight {
	if len(g.Edges) == 0 {
		return nil
	}
	
	w := g.Edges[0].Weight
	for _, edge := range g.Edges[1:] {
		w = w.Add(edge.Weight)
	}
	return w
}

func (g *Undirected) AddEdge(a Node, b Node, weight Weight) {
	g.Edges = append(g.Edges, Edge{a, b, weight})
}

func (g *Undirected) HasEdge(a Node, b Node) bool {
	for _, edge := range g.Edges {
		if (edge.A == a && edge.B == b) || (edge.A == b && edge.B == a) {
			return true
		}
	}
	return false
}

func (g *Undirected) Neighbours(node Node) (nodes []Node) {
	for _, edge := range g.Edges {
		if edge.A == node {
			nodes = append(nodes, edge.B)
		} else if edge.B == node {
			nodes = append(nodes, edge.A)
		}
	}
	return nodes
}

func (g *Undirected) Nodes() (nodes []Node) {
	var set nodeSet
	
	for _, edge := range g.Edges {
		set.insert(edge.A)
		set.insert(edge.B)
	}
	
	return []Node(set)
}

func (g *Undirected) HasNode(node Node) bool {
	for _, edge := range g.Edges {
		if edge.A == node || edge.B == node {
			return true
		}
	}
	return false
}

// Find a minimal spanning tree using Prim's algorithm.
// https://en.wikipedia.org/wiki/Prim%27s_algorithm
func (g *Undirected) MinimalSpanningTree() (mst *Undirected) {
	nodes := g.Nodes()
	numNodes := len(g.Nodes())
	start := nodes[0]
	
	mst = NewUndirected()
	
	// Begin with tree containing only the start node.
	var nodesInMST nodeSet
	nodesInMST.insert(start)
	
	for len(nodesInMST) < numNodes {
		// Find the lowest weight edge connecting a node in the tree to a node
		// not in the tree.
		var bestEdge Edge
		hasBest := false
		for _, edge := range g.Edges {
			hasA := nodesInMST.has(edge.A)
			hasB := nodesInMST.has(edge.B)
			
			if (hasA && !hasB) || (hasB && !hasA) {
				// edge is a candidate for addition
				if !hasBest || edge.Weight.Less(bestEdge.Weight) {
					bestEdge = edge
					hasBest = true
				}
			}
		}
		
		if hasBest == false {
			panic("something went wrong (Prim's algorithm ran out of edges before all nodes were added to the MST)")
		}
		
		// Add edge to tree and add node to node set.
		mst.Edges = append(mst.Edges, bestEdge)
		if nodesInMST.has(bestEdge.A) {
			nodesInMST.insert(bestEdge.B)
		} else {
			nodesInMST.insert(bestEdge.A)
		}
	}
	
	return mst
}

// a sorted, unique list of nodes
type nodeSet []Node

func (ss *nodeSet) has(node Node) bool {
	for _, n := range *ss {
		if n == node {
			return true
		}
		if n > node {
			return false
		}
	}
	return false
}

func (ss *nodeSet) insert(node Node) {
	s := *ss
	for i, n := range s {
		if n == node {
			return // already exists in set
		}
		if n > node {
			// insert before this position
			// first, shift all elements from i onwards forwards one place
			s = append(s, 0)
			for j := len(s)-2; j >= i; j-- {
				s[j+1] = s[j]
			}
			// then set position i to the new element
			s[i] = node
			*ss = s
			return
		}
	}
	// append to end
	s = append(s, node)
	*ss = s
	return
}
