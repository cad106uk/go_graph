package go_graph

import (
	"github.com/cad106uk/go_graph/node_edges"
	"sync"
)

type arrayStringCount struct {
	concurrentCount
}

// If we have no more running goroutines close down this search
func (cc *arrayStringCount) closeChannels() {
	closeChannels := func() {
		close(cc.edgeStep)
		close(cc.nodeStep)
	}
	if cc.edgeCount == 0 && cc.nodeCount == 0 {
		cc.once.Do(closeChannels)
	}
}

func (cc *arrayStringCount) NextStep(node *NodeStep) {
	defer cc.closeChannels()
	if len(node.Edges) == 0 {
		cc.nodeCount--
		return
	}
	currentEdges := node.Edges[0]

	for _, val := range node.Node.GetConnectTo() {
		edgeTypeName := val.EdgeType.GetName()
		for i := 0; i < len(currentEdges); i++ {
			if edgeTypeName == currentEdges[i] {
				cc.edgeCount++
				cc.edgeStep <- EdgeStep{val, node.Edges[1:]}
			}
		}
	}
	cc.nodeCount--
}

func (cc *arrayStringCount) TakeStep(edge *EdgeStep) {
	node := edge.Edge.ConnectTo
	cc.nodeStep <- NodeStep{*node, edge.Edges}
	cc.nodeCount++
	cc.edgeCount--
}

func (cc *arrayStringCount) ProcessEdges() {
	go func() {
	breakLabel:
		for {
			select {
			case edge, ok := <-cc.edgeStep:
				if ok {
					go cc.TakeStep(&edge)
				} else {
					break breakLabel
				}
			}
		}
	}()
}

func (cc *arrayStringCount) ProcessNodes() {
	go func() {
	breakLabel:
		for {
			select {
			case node, ok := <-cc.nodeStep:
				if ok {
					cc.output <- node.Node
					go cc.NextStep(&node)
				} else {
					close(cc.output)
					break breakLabel
				}
			}
		}
	}()
}

// A helper function to start walking a graph. Output is the interfaces concern. Always start on a GraphNode
func StartArrayStringWalkingPath(edges [][]string, output chan node_edges.GraphNode, start *node_edges.GraphNode) {
	edgeStep := make(chan EdgeStep, 10)
	nodeStep := make(chan NodeStep, 10)
	var once sync.Once

	// Setup the graph walker correctly before running it
	cc := arrayStringCount{concurrentCount{0, 1, edgeStep, nodeStep, output, once}}
	cc.NextStep(&NodeStep{*start, edges})

	pw := PathWalker(&cc)
	walkPath(pw)
}
