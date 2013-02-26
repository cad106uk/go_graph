// This package provides the basic structure for building a graph
package go_graph

type GraphNode struct {
	value               dataNode
	baseConnections     []GraphEdge // The GraphEdges that use this node as a starting point
	terminalConnections []GraphEdge // The GraphEdges the use this node as an end point
}

// returns the value that the GraphNode has
func (g *GraphNode) Value() (document, error) {
	val := g.value.document
	return val, nil
}

// Set the value this GraphNode stores. Can be called many times but onyl sets a value the first time it has been called.
func (g *GraphNode) SetValue(input document) error {
	g.value.setValue.Do(func() {
		g.value.document = input
	})
	return nil
}

func (g *GraphNode) MakeNode(nt nodeType, input document) error {
	dn, err := CreateDataNode(nt, input)
	if err != nil {
		return err
	}
	g.value = dn
	return nil
}

// // Given a GraphNode this graph node will add the given GraphNode to its own connectsTo slice and will also add
// // itself to the given GraphNodes connectedFrom slice.
// func (g *GraphNode) Connect(node GraphNode) error {
// 	g.connectsTo = append(g.connectsTo, node.dataNode)
// 	node.connectedFrom = append(node.connectedFrom, g.dataNode)
// 	return nil
// }
