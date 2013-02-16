package go_graph

import "sync"

type document interface{}

type dataNode struct {
	document  // The data stored at this node
	setValue sync.Once // The value can only be set once
}

type GraphNode struct {
	dataNode
	connectsTo []dataNode // The nodes this points to
	connectedFrom []dataNode // The nodes that point to this
}


// returns the value that the GraphNode has
func (g *GraphNode) Value() (document, error) {
	val := g.document
	return val, nil
}

// Set the value this GraphNode stores. Can be called many times but onyl sets a value the first time it has been called.
func (g *GraphNode) SetValue(input document) error {
	g.setValue.Do(func() {

	})
	return nil
}

// Given a GraphNode this graph node will add the given GraphNode to its own connectsTo slice and will also add itself to the given GraphNodes connectedFrom slice.
func (g *GraphNode) Connect(node GraphNode) error {
	g.connectsTo = append(g.connectsTo, node.dataNode)
	node.connectedFrom = append(node.connectedFrom, g.dataNode)
	return nil 
}