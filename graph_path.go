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

type concurrentCount struct {
	edgeCount int64
	nodeCount int64
	edgeStep  chan EdgeStep
	nodeStep  chan NodeStep
	once      sync.Once
}

// A helper function to start walking a graph. Output is the interfaces concern. Always start on a GraphNode
func StartWalkingPath(edges [][]string, output chan GraphNode, start *GraphNode) {
	go func() {
		edgeStep := make(chan EdgeStep, 10)
		nodeStep := make(chan NodeStep, 10)
		var once sync.Once
		cc := concurrentCount{0, 1, edgeStep, nodeStep, once}
		cc.NextStep(&NodeStep{*start, edges}, edgeStep)

	breakLabel:
		for {
			select {
			case edge, ok := <-edgeStep:
				if ok {
					go cc.TakeStep(&edge, nodeStep)
				}
			case node, ok := <-nodeStep:
				if ok {
					output <- node.Node
					go cc.NextStep(&node, edgeStep)
				} else {
					close(output)
					break breakLabel
				}
			}
		}
	}()
}

// If we have no more running goroutines close down this search
func (cc *concurrentCount) closeChannels() {
	closeChannels := func() {
		close(cc.edgeStep)
		close(cc.nodeStep)
	}
	if cc.edgeCount == 0 && cc.nodeCount == 0 {
		cc.once.Do(closeChannels)
	}
}

func (cc *concurrentCount) NextStep(node *NodeStep, step chan EdgeStep) {
	defer cc.closeChannels()
	if len(node.Edges) == 0 {
		cc.nodeCount--
		return
	}
	currentEdges := node.Edges[0]

	for _, val := range node.Node.connectTo {
		edgeTypeName := val.EdgeType.edgeTypeName
		for i := 0; i < len(currentEdges); i++ {
			if edgeTypeName == currentEdges[i] {
				cc.edgeCount++
				step <- EdgeStep{val, node.Edges[1:]}
			}
		}
	}
	cc.nodeCount--
}

func (cc *concurrentCount) TakeStep(edge *EdgeStep, nodeStep chan NodeStep) {
	cc.nodeCount++
	node := edge.Edge.ConnectTo
	nodeStep <- NodeStep{*node, edge.Edges}
	cc.edgeCount--
}
