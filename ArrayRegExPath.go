package go_graph

import (
	"go_graph/node_edges"
	"regexp"
	"sync"
)

type regexStringCount struct {
	concurrentCount
}

func (rsc *regexStringCount) closeChannels() {
	closeChannels := func() {
		close(rsc.edgeStep)
		close(rsc.nodeStep)
	}
	if rsc.edgeCount == 0 && rsc.nodeCount == 0 {
		rsc.once.Do(closeChannels)
	}
}

func (rsc *regexStringCount) NextStep(node *NodeStep) {
	defer rsc.closeChannels()
	if len(node.Edges) == 0 {
		rsc.nodeCount--
		return
	}
	currentEdges := make([]regexp.Regexp, 0)
	for _, edge := range node.Edges[0] {
		reg, _ := regexp.Compile(edge)
		currentEdges = append(currentEdges, *reg)
	}

	for _, val := range node.Node.GetConnectTo() {
		edgeTypeName := val.EdgeType.GetName()
		for i := 0; i < len(currentEdges); i++ {
			if currentEdges[i].MatchString(edgeTypeName) == true {
				rsc.edgeCount++
				rsc.edgeStep <- EdgeStep{val, node.Edges[1:]}
			}
		}
	}
	rsc.nodeCount--
}

func (rsc *regexStringCount) TakeStep(edge *EdgeStep) {
	node := edge.Edge.ConnectTo
	rsc.nodeStep <- NodeStep{*node, edge.Edges}
	rsc.nodeCount++
	rsc.edgeCount--
}

func (rsc *regexStringCount) ProcessEdges() {
	go func() {
	breakLabel:
		for {
			select {
			case edge, ok := <-rsc.edgeStep:
				if ok {
					go rsc.TakeStep(&edge)
				} else {
					break breakLabel
				}
			}
		}
	}()
}

func (rsc *regexStringCount) ProcessNodes() {
	go func() {
	breakLabel:
		for {
			select {
			case node, ok := <-rsc.nodeStep:
				if ok {
					rsc.output <- node.Node
					go rsc.NextStep(&node)
				} else {
					close(rsc.output)
					break breakLabel
				}
			}
		}
	}()
}

// A helper function to start walking a graph. Output is the interfaces concern. Always start on a GraphNode
func StartArrayRegExWalkingPath(edges [][]string, output chan node_edges.GraphNode, start *node_edges.GraphNode) {
	edgeStep := make(chan EdgeStep, 10)
	nodeStep := make(chan NodeStep, 10)
	var once sync.Once

	// Setup the graph walker correctly before running it
	rsc := regexStringCount{concurrentCount{0, 1, edgeStep, nodeStep, output, once}}
	rsc.NextStep(&NodeStep{*start, edges})

	pw := PathWalker(&rsc)
	walkPath(pw)
}
