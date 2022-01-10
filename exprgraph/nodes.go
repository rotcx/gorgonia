package exprgraph

import (
	"gonum.org/v1/gonum/graph"
)

// IterNodes is an iterator over a slice of *Node
type IterNodes struct {
	ns []*Node
	i  int
}

// Len returns the remaining number of nodes to be iterated over.
func (ns *IterNodes) Len() int { return len(ns.ns) }

// Next returns whether the next call of Node will return a valid node.
func (ns *IterNodes) Next() bool { ns.i++; return ns.i < len(ns.ns) }

// Node returns the current node of the iterator. Next must have been
// called prior to a call to Node.
func (ns *IterNodes) Node() graph.Node {
	if ns.i < 0 || ns.i >= len(ns.ns) {
		return nil
	}
	return ns.ns[ns.i]
}

// Reset returns the iterator to its initial state.
func (n *IterNodes) Reset() { n.i = -1 }

// NodeSlice returns all the remaining nodes in the iterator and advances
// the iterator. The order of nodes within the returned slice is not specified.
func (ns *IterNodes) NodeSlice() []*Node {
	if ns.i < 0 {
		ns.i = 0
	}
	retVal := ns.ns[ns.i:len(ns.ns)]
	ns.i = len(ns.ns)
	return retVal
}

// NodeIDSlice returns all the remaining nodes in the iterator and advances
// the iterator. The order of nodes within the returned slice is not specified.
func (ns *IterNodes) NodeIDSlice() []NodeID {
	if ns.i < 0 {
		ns.i = 0
	}
	retVal := make([]NodeID, 0, len(ns.ns)-ns.i+1)
	for _, n := range ns.ns {
		retVal = append(retVal, NodeID(n.id))
	}
	ns.i = len(ns.ns)
	return retVal
}

// NodesFromOrdered returns a Nodes initialized with the
// provided nodes, a map of node IDs to graph.Nodes, and the set
// of edges, a map of to-node IDs to graph.WeightedEdge, that can be
// traversed to reach the nodes that the NodesByEdge will iterate
// over.
func NodesFromOrdered(ns []*Node) *IterNodes { return &IterNodes{ns: ns, i: -1} }

// NodesFromIDs creates a *Nodes (an iterator) from a list of int64 IDs.
func NodesFromIDs(g *Graph, ns []int64) *IterNodes {
	nodes := make([]*Node, 0, len(ns))
	for _, id := range ns {
		nodes = append(nodes, g.nodes[id])
	}
	return NodesFromOrdered(nodes)
}

// NodesFromNodeIDs creates a *Nodes (an iterator) from a list of NodeIDs.
func NodesFromNodeIDs(g *Graph, ns []NodeID) *IterNodes { return NodesFromIDs(g, nodeIDs2IDs(ns)) }

// NodeIDs is a set of NodeIDs.
// It implements sort.Sort as the basis of the set.
// (see github.com/xtgo/set)
type NodeIDs []NodeID

func (ns NodeIDs) Contains(a NodeID) bool {
	for _, n := range ns {
		if n == a {
			return true
		}
	}
	return false
}

/* NodeIDs implements sort.Sort */

func (ns NodeIDs) Len() int           { return len(ns) }
func (ns NodeIDs) Less(i, j int) bool { return ns[i] < ns[j] }
func (ns NodeIDs) Swap(i, j int)      { ns[i], ns[j] = ns[j], ns[i] }
