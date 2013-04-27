package go_graph

/*
 This package is not here to find the sortest path, this is here to analyses the structure of the data in the graph.

 Each valid path will be searched in its own goroutine returning results down a channel. The output from the searches will be largely unstructured. The options for return values being limited to an unorder list of graph nodes, or node - edge - node structure being returned in any order. Building a complete set output data is not done here. (Each option below can be run in either mode)

 * An array of an array of strings each string is an edgeType and the order of strings array says what edges should be followed.
 * An array of an array of regex to find which edges can be followed from this node.
 * A function to programically decide what to do next
*/

type pathWalker interface {
	NextStep(node *GraphNode) (edges []GraphEdge, err NodeError) // Each node has many edges
	TakeStep(edge *GraphEdge) (node GraphNode, err NodeError)    // Each edge points to only 1 node.
}

type PathOutput struct {
	FromNode    *GraphNode
	ConnectEdge *GraphEdge
	ToNode      *GraphNode
}

type ArrayWalkStringNodesOutput struct {
	edges      [][]string
	OutputChan chan []GraphNode
}

func (aws *ArrayWalkStringNodesOutput) NextStep(node *GraphNode) (edges []GraphEdge, err NodeError) {
	// Write me
}

func (aws *ArrayWalkStringNodesOutput) TakeStep(edge *GraphEdge) (node GraphNode, err NodeError) {
	// Write me
}
