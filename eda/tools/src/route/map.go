package main

// Standard library imports
import (
	"log"
	"math"
)

// Local imports
import (
	"format/footprint"
	"format/layout"
	"format/netlist"
	"util/network"
)

// Convert the given netlist into a set of wires, which are attached to the layout.
// It is assumed that the netlist is valid (no net has less than two nodes etc.)
func Map(l *layout.Layout, fps map[string]*footprint.Footprint, nets netlist.Netlist) {
	var wires []*layout.Wire
	
	g := network.NewUndirected()
	
	// Route each net individually.
	for net, nodes := range nets {
		// Obtain node coordinates.
		xs := make([]int, len(nodes))
		ys := make([]int, len(nodes))
		for i, node := range nodes {
			comp := l.Components[node.Component]
			if comp == nil {
				log.Fatalf("undefined component '%s'", node.Component)
			}
			x, y := pinPos(comp, fps, node.Pin)
			xs[i] = x
			ys[i] = y
		}
		
		// Add an edge between between each pair of nodes.
		g.Reset()
		for i := 0; i < len(nodes); i++ {
			x1, y1 := xs[i], ys[i]
			
			for j := i+1; j < len(nodes); j++ {
				x2, y2 := xs[j], ys[j]
				dx, dy := x2-x1, y2-y1
				lengthSquared := dx*dx + dy*dy
				
				g.AddEdge(network.Node(i), network.Node(j), network.IntWeight(lengthSquared))
			}
		}
		
		// Generate a minimal spanning tree.
		mst := g.MinimalSpanningTree()
		
		// Add wires to the layout where edges were placed in the MST.
		for _, edge := range mst.Edges {
			node1 := nodes[edge.A]
			node2 := nodes[edge.B]
			squaredLength := edge.Weight.(network.IntWeight)
			length := math.Sqrt(float64(squaredLength))
			wires = append(wires, &layout.Wire{
				Component1: node1.Component,
				Pin1: node1.Pin,
				Component2: node2.Component,
				Pin2: node2.Pin,
				Net: net,
				Length: length,
			})
		}
	}
	
	l.Wires = wires
}


/*
// Convert the given netlist into a set of wires, which are attached to the
// layout as a new layer.
// It is assumed that the netlist is valid (no net has less than two nodes etc.)
func Map(l *layout.Layout, fps map[string]*footprint.Footprint, nets netlist.Netlist) {
	// create a node->pinpos mapping and a node->net mapping
	positions := make(map[netlist.Node][2]int)
	revNets := make(map[netlist.Node]string)
	totalX, totalY, numPins := 0, 0, 0
	for net, nodes := range nets {
		for _, node := range nodes {
			c := l.Components[node.Component]
			if c == nil {
				panic("invalid component")
			}
			x, y := pinPos(c, fps, node.Pin)
			positions[node] = [2]int{x, y}
			revNets[node] = net
			totalX += x
			totalY += y
			numPins++
		}
	}
	
	meanX := int(float64(totalX) / float64(numPins))
	meanY := int(float64(totalY) / float64(numPins))
	
	var wires []*layout.Wire
	
	for len(positions) > 0 {
		// find the node furthest from the centre
		var bestNode netlist.Node
		bestDistSquared := 0
		for node, pos := range positions {
			dx, dy := pos[0]-meanX, pos[1]-meanY
			distSquared := dx*dx + dy*dy
			if distSquared > bestDistSquared {
				bestDistSquared = distSquared
				bestNode = node
			}
		}
		
		node1 := bestNode
		pos1 := positions[node1]
		net := revNets[node1]
		
		// this node is already considered placed, so remove it from positions & revNets
		delete(positions, node1)
		delete(revNets, node1)
		
		// begin routing wires from this node
		nodes := nets[net]
		nodesToRoute := len(nodes) - 1
		for nodesToRoute > 0 {
			bestDistSquared = 0
			for _, node2 := range nodes {
				_, ok := positions[node2]
				if ok {
					// node2 is a candidate for routing
					pos2 := positions[node2]
					dx, dy := pos2[0]-pos1[0], pos2[1]-pos1[1]
					distSquared := dx*dx + dy*dy
					if bestDistSquared == 0 || distSquared < bestDistSquared {
						bestDistSquared = distSquared
						bestNode = node2
					}
				}
			}
			if bestDistSquared == 0.0 {
				// no nodes available to route to
				panic("something went wrong (cannot rely on nodesToRoute for termination of routing loop)")
			}
			node2 := bestNode
			pos2 := positions[node2]
			// route from node1 to node2
			wires = append(wires, &layout.Wire{
				Component1: node1.Component,
				Pin1: node1.Pin,
				Component2: node2.Component,
				Pin2: node2.Pin,
				Net: net,
			})
			// update current node
			node1 = node2
			pos1 = pos2
			// remove new current node from positions & revNets
			delete(positions, node1)
			delete(revNets, node1)
			nodesToRoute--
		}
	}
	
	l.WireLayers = append(l.WireLayers, &layout.WireLayer{
		Wires: wires,
	})
}
*/
