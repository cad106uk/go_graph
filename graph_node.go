// This package provides the basic structure for building a graph
// the graph node is ment encapsulate a data node and all the
// graph edges the come to and from this node

package go_graph

type GraphNode struct {
	value               dataNode
	baseConnections     []GraphEdge // The GraphEdges that use this node as a starting point
	terminalConnections []GraphEdge // The GraphEdges the use this node as an end point
}

// returns the value that the GraphNode has
func (g *GraphNode) Value() (*data, error) {
	val := g.value.data
	return &val, nil
}

// Set the value this GraphNode stores. Can be called many times but onyl sets a value the first time it has been called.
func (g *GraphNode) SetValue(input data) error {
	g.value.setValue.Do(func() {
		g.value.data = input
	})
	return nil
}

func (g *GraphNode) MakeNode(nt *nodeType, input []byte) error {
	dn, err := CreateDataNode(nt, input)
	if err != nil {
		return err
	}
	g.value = dn
	return nil
}

// // A.ConnectTo(B) Means A->From-GraphEdge-To->B
// // If the types allow.
// func (g *GraphNode) ConnectTo(node GraphNode) error {
// 	g.connectsTo = append(g.connectsTo, node.dataNode)
// 	node.connectedFrom = append(node.connectedFrom, g.dataNode)
// 	return nil
// }
