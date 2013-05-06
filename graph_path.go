package go_graph

import "sync"

/*
 This package is not here to find the sortest path, this is here to analyses the structure of the data in the graph.

 Each valid path will be searched in its own goroutine returning results down a channel. The output from the searches will be largely unstructured. The options for return values being limited to an unorder list of graph nodes, or node - edge - node structure being returned in any order. Building a complete set output data is not done here. (Each option below can be run in either mode)

 * An array of an array of strings each string is an edgeType and the order of strings array says what edges should be followed.
 * An array of an array of regex to find which edges can be followed from this node.
 * A function to programically decide what to do next
*/

type NodeStep struct {
	Node  GraphNode
	Edges [][]string
}

type EdgeStep struct {
	Edge  GraphEdge
	Edges [][]string
}

type PathWalker interface {
	NextStep(node *NodeStep)
	TakeStep(edge *EdgeStep)
	ProcessEdges()
	ProcessNodes()
}

func walkPath(pw PathWalker) {
	pw.ProcessEdges()
	pw.ProcessNodes()
}

// This is a control structure to keep track of the number of goroutines
// running, while the graph is being traversed.
type concurrentCount struct {
	edgeCount int64
	nodeCount int64
	edgeStep  chan EdgeStep
	nodeStep  chan NodeStep
	output    chan GraphNode
	once      sync.Once
}
